package server

import (
	"context"
	"log"
	"sync"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

const (
	slaveTag = "slave"
)

// MasterServer implements master nodes in master-slave mode.
type MasterServer struct {
	SingleMachineServer
	slaves *sync.Map
}

// NewMasterServer creates master server in master-slave mode.
func NewMasterServer() *MasterServer {
	slv := &sync.Map{}
	slv.Store(slaveTag, []chan *pb.SlaveSubscribeResponse{})
	return &MasterServer{
		SingleMachineServer: SingleMachineServer{&sync.Map{}},
		slaves:              slv,
	}
}

// Publish a message
func (s *MasterServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	slv, _ := s.slaves.Load(slaveTag)
	for _, c := range slv.([]chan *pb.SlaveSubscribeResponse) {
		c <- &pb.SlaveSubscribeResponse{
			Topic: req.Topic,
			Msg:   req.Msg,
		}
	}
	return s.SingleMachineServer.Publish(ctx, req)
}

// SubscribeFromMaster allows slaves to subscribe to master
func (s *MasterServer) SubscribeFromMaster(req *pb.SlaveSubscribeRequest, stream pb.MasterSidecar_SubscribeFromMasterServer) error {
	c := make(chan *pb.SlaveSubscribeResponse)

	slv, _ := s.slaves.Load(slaveTag)
	s.slaves.Store(slaveTag, append(slv.([]chan *pb.SlaveSubscribeResponse), c))

	for {
		msg := <-c
		log.Printf("incoming message to slave %s", msg)
		if err := stream.Send(msg); err != nil {
			log.Printf("failed to publish: %v\n", err)
			return err
		}
	}
}

// SlaveServer implements slave nodes in master-slave mode.
type SlaveServer struct {
	MasterServer
	mst   pb.PubSubClient
	mstSc pb.MasterSidecarClient
}

// NewSlaveServer creates slave server in master-slave mode.
func NewSlaveServer(conn *grpc.ClientConn) *SlaveServer {
	svr := &SlaveServer{
		MasterServer: *NewMasterServer(),
		mst:          pb.NewPubSubClient(conn),
		mstSc:        pb.NewMasterSidecarClient(conn),
	}

	if err := svr.init(); err != nil {
		log.Fatalf("cannot initiate slave server: %s", err)
	}

	return svr
}

func (s *SlaveServer) init() error {
	stream, err := s.mstSc.SubscribeFromMaster(context.Background(), &pb.SlaveSubscribeRequest{})

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

			slv, _ := s.slaves.Load(slaveTag)
			for _, c := range slv.([]chan *pb.SlaveSubscribeResponse) {
				c <- res
			}

			if chs, ok := s.m.Load(res.Topic.Name); ok {
				for _, c := range chs.([]chan *pb.Message) {
					c <- res.Msg
				}
			}
		}
	}()

	return nil
}

// Publish a message
func (s *SlaveServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	log.Printf("publish reroute to master %s", req)
	return s.mst.Publish(ctx, req)
}
