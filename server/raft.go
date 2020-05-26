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
	id     string // addr:port
	peers  *sync.Map
	role   pb.Role
	term uint64
	leaderId string
	heartbeatTimeout time.Duration
}

func NewRaftServer(cfg *pb.ServerConfig) *RaftServer {
	// Get the id
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	id := fmt.Sprintf("%s:%d", conn.LocalAddr().(*net.UDPAddr).IP.String() , cfg.RaftConfig.Port)

	// Parse peers
	peers := &sync.Map{}
	timeout, err := time.ParseDuration("10m")
	for _, peer := range strings.Split(cfg.RaftConfig.Peers, ";") {
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

	heartbeatTimeout, _ := time.ParseDuration(cfg.RaftConfig.HeartbeatTimeout)

	return &RaftServer{
		id: id,
		peers: peers,
		role: pb.Role_Candidate,
		term: 0,
		heartbeatTimeout: heartbeatTimeout,
	}
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
