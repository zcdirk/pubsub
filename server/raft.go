package server

import (
	"context"
	"fmt"
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
	"log"
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
	peerNum			  uint64
	role              pb.Role
	term              uint64
	leaderId          string
	heartbeatInterval time.Duration
	ticker            *time.Ticker
	log               []*pb.LogEntry
}

func NewRaftServer(cfg *pb.ServerConfig) *RaftServer {
	// Get the id
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	id := fmt.Sprintf("%s:%d", conn.LocalAddr().(*net.UDPAddr).IP.String(), cfg.RaftConfig.Port)

	// Parse peers
	peers := &sync.Map{}
	timeout, err := time.ParseDuration("10m")
	peerStrs := strings.Split(cfg.RaftConfig.Peers, ";")
	for _, peer := range peerStrs {
		go func() {
			log.Printf("connect to peer: %s", peer)
			conn, err := grpc.Dial(
				peer,
				grpc.WithTimeout(timeout),
				grpc.WithInsecure())
			if err != nil {
				log.Fatal(err)
			}
			peers.Store(peer, pb.NewRaftSidecarClient(conn))
		}()
	}

	// Parse heart beat interval
	heartbeatInterval, _ := time.ParseDuration(cfg.RaftConfig.HeartbeatInterval)

	svr := &RaftServer{
		id:                id,
		peers:             peers,
		peerNum: 		   uint64(len(peerStrs)),
		role:              pb.Role_Candidate,
		term:              0,
		heartbeatInterval: heartbeatInterval,
		log:               []*pb.LogEntry{},
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
	return &pb.AppendEntriesResponse{}, nil
}

func (s *RaftServer) RequestVote(ctx context.Context, req *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	return &pb.RequestVoteResponse{}, nil
}

func (s *RaftServer) heartBeat(ctx context.Context) {
	for {
		select {
		case <-s.ticker.C:
			switch s.role {
			case pb.Role_Leader:
				// Send heart beat request to all the followers
				s.peers.Range(func(k, v interface{}) bool {
					v.(pb.RaftSidecarClient).AppendEntries(ctx, &pb.AppendEntriesRequest{})
					return true
				})
			case pb.Role_Follower, pb.Role_Candidate:
				if s.role == pb.Role_Follower {
					s.term++
				}
				// Change to candidate and request vote
				s.role = pb.Role_Candidate
				voteChannel, voteNum := make(chan bool), uint64(0)
				s.peers.Range(func(k, v interface{}) bool {
					req, _ := v.(pb.RaftSidecarClient).RequestVote(ctx, &pb.RequestVoteRequest{
						Term:         s.term,
						CandidateId:  s.id,
						LastLogIndex: uint64(len(s.log)),
						LastLogTerm:  s.log[len(s.log) - 1].Term,
					})
					if req.VoteGranted {
						voteChannel <- req.VoteGranted
					}
					return true
				})
				timer := time.NewTimer(s.getTimeOut())
				for s.role == pb.Role_Candidate && !s.isMajority(voteNum) && !Expired(timer) {
					<-voteChannel
					voteNum++
				}
				if s.role == pb.Role_Candidate && s.isMajority(voteNum) && !Expired(timer){
					// Upgrade to leader
					s.role = pb.Role_Leader
					s.heartBeat(ctx)
					s.ticker = time.NewTicker(s.heartbeatInterval)
				}
			}
		}

	}
}

func Expired(T *time.Timer) bool {
	select {
	case <-T.C:
		return true
	default:
		return false
	}
}

func (s *RaftServer) isMajority(voteNum uint64) bool {
	return voteNum > s.peerNum / 2
}

func (s *RaftServer) getTimeOut() time.Duration {
	return s.heartbeatInterval * 3
}
