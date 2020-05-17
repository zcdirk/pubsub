package server

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/golang/protobuf/proto"
)

var (
	sysHeader = "sys"
)

func isInReservedNamespace(t string) bool {
	for _, h := range []string{sysHeader} {
		if strings.HasPrefix(t, h+"+") {
			return true
		}
	}
	return false
}

type masterServer struct {
	singleMachineServer
}

// NewMasterServer creates master server in master-slave mode.
func NewMasterServer() pb.PubSubServer {
	return &masterServer{
		singleMachineServer: singleMachineServer{&sync.Map{}},
	}
}

func (s *masterServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	if isInReservedNamespace(req.Topic.Name) {
		return &pb.PublishResponse{
			Status: &pb.PublishResponse_Failure_{
				Failure: &pb.PublishResponse_Failure{
					Reason: fmt.Sprintf("topic name is in namespace reserved for internal use only: %s", req.Topic.Name),
				},
			},
		}, nil
	}

	if chs, ok := s.m.Load("sys-slave"); ok {
		for _, c := range chs.([]chan *pb.Message) {
			c <- &pb.Message{Content: req.String()}
		}
	}

	return s.singleMachineServer.Publish(ctx, req)
}

type slaveServer struct {
	singleMachineServer
	mst pb.PubSubClient
}

// NewSlaveServer creates slave server in master-slave mode.
func NewSlaveServer(mst pb.PubSubClient) pb.PubSubServer {
	svr := &slaveServer{
		singleMachineServer: singleMachineServer{&sync.Map{}},
		mst:                 mst,
	}

	if err := svr.init(); err != nil {
		log.Fatalf("cannot initiate slave server: %s", err)
	}

	return svr
}

func (s *slaveServer) init() error {
	stream, err := s.mst.Subscribe(context.Background(),
		&pb.SubscribeRequest{
			Topic: []*pb.Topic{&pb.Topic{Name: "sys-slave"}},
		})

	if err != nil {
		return err
	}

	// Launch side-car process to broadcast messages from master.
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}

			req := &pb.PublishRequest{}
			if err := proto.UnmarshalText(res.Msg.Content, req); err != nil {
				log.Fatalf("cannot unmarshal publish request from master: %s", err)
			}

			if chs, ok := s.m.Load(req.Topic.Name); ok {
				for _, c := range chs.([]chan *pb.Message) {
					c <- req.Msg
				}
			}
		}
	}()

	return nil
}

func (s *slaveServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	log.Printf("publish reroute to master %s", req)

	return s.mst.Publish(ctx, req)
}
