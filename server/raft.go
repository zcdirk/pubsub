package server

import (
	"context"
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"sync"
)

// RaftSever is implementation for PubSub
type RaftServer struct {
	SingleMachineServer
	id    string  // addr:port
	peers *sync.Map
}

func NewRaftServer() *RaftServer {
	peers := &sync.Map{}
	return &RaftServer{
		peers: peers,
	}
}

// Publish a message
func (s *RaftServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	// TODO: Implement the method to redirect the message to leader.
	return s.SingleMachineServer.Publish(ctx, req)
}
