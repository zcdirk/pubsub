package server

import (
	"context"
	"log"
	"sync"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
)

type singleMachineServer struct {
	m *sync.Map
}

// NewSingleMachineServer creates single machine PubSub implementation
func NewSingleMachineServer() pb.PubSubServer {
	return &singleMachineServer{&sync.Map{}}
}

// Publish a message
func (s *singleMachineServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	log.Printf("publish %s", req)

	if chs, ok := s.m.Load(req.Topic.Name); ok {
		for _, c := range chs.([]chan *pb.Message) {
			c <- req.Msg
		}
	}

	return &pb.PublishResponse{
		Status: &pb.PublishResponse_Success_{
			Success: &pb.PublishResponse_Success{},
		},
	}, nil
}

// Subscribe a topic
func (s *singleMachineServer) Subscribe(req *pb.SubscribeRequest, stream pb.PubSub_SubscribeServer) error {
	log.Printf("subscribe %s", req)

	c := make(chan *pb.Message)
	for _, t := range req.Topic {
		n := t.Name
		chs, ok := s.m.Load(n)
		if ok {
			s.m.Store(n, append(chs.([]chan *pb.Message), c))
		} else {
			s.m.Store(n, []chan *pb.Message{c})
		}
	}

	for {
		log.Printf("incoming message to subscriber %s", req)
		msg := <-c
		err := stream.Send(&pb.SubscribeResponse{Msg: msg})
		if err != nil {
			log.Printf("failed to publish: %v\n", err)
			return err
		}
	}
}
