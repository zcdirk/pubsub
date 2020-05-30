package server

import (
	"context"
	"fmt"
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

// RaftSever is implementation for PubSub
type RaftServer struct {
	SingleMachineServer
	id                string // addr:port
	peers             *sync.Map
	peerNum           uint64
	role              pb.Role
	term              uint64
	leaderId          string
	heartbeatInterval time.Duration
	ticker            *time.Ticker
	log               []*pb.LogEntry
	voteFor           string
}

type Peer struct {
	client   pb.PubSubClient
	sideCar  pb.RaftSidecarClient
	logIndex int
}

func NewRaftServer(cfg *pb.ServerConfig) *RaftServer {
	// Get the id
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	id := fmt.Sprintf("%s:%d", conn.LocalAddr().(*net.UDPAddr).IP.String(), cfg.Port)
	log.Printf("Server id: %s", id)

	// Parse peers
	peers := &sync.Map{}
	timeout, err := time.ParseDuration("10m")
	peerStrs := strings.Split(cfg.RaftConfig.Peers, ";")
	for _, peer := range peerStrs {
		log.Printf("connect to peer: %s", peer)
		conn, err := grpc.Dial(
			peer,
			grpc.WithTimeout(timeout),
			grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		peers.Store(
			peer,
			&Peer{
				client:   pb.NewPubSubClient(conn),
				sideCar:  pb.NewRaftSidecarClient(conn),
				logIndex: 1,
			})
	}

	// Parse heart beat interval
	heartbeatInterval, _ := time.ParseDuration(cfg.RaftConfig.HeartbeatInterval)

	svr := &RaftServer{
		id:                id,
		peers:             peers,
		peerNum:           uint64(len(peerStrs)),
		role:              pb.Role_Candidate,
		term:              0,
		heartbeatInterval: heartbeatInterval,
		log:               []*pb.LogEntry{{Term: 0}}, // Append an empty log entry for prev log index and term
	}

	// Set ticker
	svr.ticker = time.NewTicker(svr.getTimeOut())
	go svr.heartBeat(context.Background())
	return svr
}

// Publish a message
func (s *RaftServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	// TODO: Implement the method to redirect the message to leader.
	return s.SingleMachineServer.Publish(ctx, req)
}

func (s *RaftServer) AppendEntries(ctx context.Context, req *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	log.Printf("AppendEntries: %+v", req)
	// Reset ticker
	s.resetTicker(s.getTimeOut())
	if s.voteFor != "" {
		s.voteFor = ""
	}
	if s.term > req.Term {
		return &pb.AppendEntriesResponse{
			Term:   s.term,
			Status: pb.AppendEntriesResponse_FAILURE,
			Info:   "Leader's term need to be updated",
		}, nil
	}
	if len(s.log) <= int(req.PrevLogIndex) || s.log[req.PrevLogIndex].Term != req.PrevLogTerm {
		return &pb.AppendEntriesResponse{
			Term:   s.term,
			Status: pb.AppendEntriesResponse_FAILURE,
			Info:   "Prev index and term mismatch",
		}, nil
	}
	s.role = pb.Role_Follower
	s.leaderId = req.LeaderId
	s.term = req.Term
	s.log = append(s.log[:req.PrevLogIndex+1], req.Entries...)
	if len(req.Entries) > 0 {
		for _, e := range req.Entries {
			s.SingleMachineServer.Publish(ctx, &pb.PublishRequest{Msg: e.Msg, Topic: e.Topic})
		}
	}
	return &pb.AppendEntriesResponse{
		Term:   s.term,
		Status: pb.AppendEntriesResponse_SUCCESS,
	}, nil
}

func (s *RaftServer) RequestVote(ctx context.Context, req *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	log.Printf("RequestVote: %+v", req)
	if s.role != pb.Role_Leader &&
		s.term <= req.Term &&
		s.voteFor == "" &&
		len(s.log) <= int(req.LastLogIndex+1) {
		s.voteFor = req.CandidateId
		return &pb.RequestVoteResponse{
			Term:        s.term,
			VoteGranted: true,
		}, nil
	}
	return &pb.RequestVoteResponse{
		Term:        s.term,
		VoteGranted: false,
	}, nil
}

func (s *RaftServer) heartBeat(ctx context.Context) {
	for {
		select {
		case <-s.ticker.C:
			switch s.role {
			case pb.Role_Leader:
				log.Printf("Leader heartbeat")
				// Send heart beat request to all the followers
				s.sendHeartBeatRequest(ctx)
			case pb.Role_Follower, pb.Role_Candidate:
				log.Printf("Timeout heartbeat")
				if s.role == pb.Role_Follower {
					s.term++
					s.role = pb.Role_Candidate
				} else {
					// Sleep a random time to avoid conflict for leader election
					time.Sleep(s.heartbeatInterval  * time.Duration(rand.Intn(5000)) / time.Duration(1000))
				}
				if s.voteFor != "" {
					continue
				}
				// Request vote
				voteChannel, voteNum := make(chan bool), uint64(0)
				s.peers.Range(func(k, v interface{}) bool {
					go func() {
						log.Printf("node:%s, peer:%+v", k, v)
						res, _ := v.(*Peer).sideCar.RequestVote(ctx, &pb.RequestVoteRequest{
							Term:         s.term,
							CandidateId:  s.id,
							LastLogIndex: uint64(len(s.log) - 1),
							LastLogTerm:  s.log[len(s.log)-1].Term,
						})
						log.Printf("res:%+v", res)
						if res != nil && res.VoteGranted {
							voteChannel <- res.VoteGranted
						}
					}()
					return true
				})
				timer := time.NewTimer(s.getTimeOut())
				for s.role == pb.Role_Candidate && !s.isMajority(voteNum) && !expired(timer) {
					select {
					case <-voteChannel:
						voteNum++
					default:
					}
				}
				if s.role == pb.Role_Candidate && s.isMajority(voteNum) &&
					!expired(timer) && s.voteFor == "" {
					// Upgrade to leader
					log.Printf("Upgrade to leader")
					s.role = pb.Role_Leader
					s.sendHeartBeatRequest(ctx)
					s.resetTicker(s.heartbeatInterval)
				}
			}
		}

	}
}

func (s *RaftServer) sendHeartBeatRequest(ctx context.Context) {
	s.peers.Range(func(k, v interface{}) bool {
		req, index := &pb.AppendEntriesRequest{}, v.(*Peer).logIndex
		if index <= len(s.log)-1 {
			req.Term = s.log[index].Term
			req.Entries = s.log[index:]
			req.LeaderId = s.id
			req.PrevLogTerm = s.log[index-1].Term
			req.PrevLogIndex = uint64(index - 1)
		}
		v.(*Peer).sideCar.AppendEntries(ctx, req)
		return true
	})
}

func expired(T *time.Timer) bool {
	select {
	case <-T.C:
		return true
	default:
		return false
	}
}

func (s *RaftServer) isMajority(voteNum uint64) bool {
	return voteNum > (s.peerNum + 1) /2
}

func (s *RaftServer) getTimeOut() time.Duration {
	return s.heartbeatInterval * 3
}

func (s *RaftServer) resetTicker(t time.Duration) {
	s.ticker.Stop()
	s.ticker = time.NewTicker(t)
}
