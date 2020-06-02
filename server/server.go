package server

import (
	"fmt"
	"log"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

// CreatePubsubServer creates PubSub server implementation based on user config
func CreatePubsubServer(cfg *pb.ServerConfig, opts ...grpc.ServerOption) (*grpc.Server, error) {
	svr := grpc.NewServer(opts...)

	switch cfg.ReplicationMode {
	case pb.ServerConfig_SINGLE_MACHINE:
		pb.RegisterPubSubServer(svr, NewSingleMachineServer())
		return svr, nil

	case pb.ServerConfig_MASTER_SLAVE:
		switch msCfg := cfg.MasterSlaveConfig; msCfg.Mode {
		case pb.ServerConfig_MasterSlaveConfig_MASTER:
			mst := NewMasterServer()
			pb.RegisterPubSubServer(svr, mst)
			pb.RegisterMasterSidecarServer(svr, mst)

		case pb.ServerConfig_MasterSlaveConfig_SLAVE:
			connTimeout, err := time.ParseDuration(msCfg.MasterTimeout)
			if err != nil {
				return nil, err
			}

			mst := fmt.Sprintf("%s:%d", msCfg.MasterAddress, msCfg.MasterPort)
			log.Printf("connect to master server: %s", mst)
			conn, err := grpc.Dial(
				mst,
				grpc.WithTimeout(connTimeout),
				grpc.WithInsecure())
			if err != nil {
				return nil, fmt.Errorf("cannot connect to master: %v", err)
			}

			slv := NewSlaveServer(conn)
			pb.RegisterPubSubServer(svr, slv)
			pb.RegisterMasterSidecarServer(svr, slv)
		}

		return svr, nil

	case pb.ServerConfig_RAFT:
		raft := NewRaftServer(cfg.RaftConfig)
		pb.RegisterPubSubServer(svr, raft)
		pb.RegisterRaftSidecarServer(svr, raft)
		return svr, nil
	}

	return nil, fmt.Errorf("current user config is not supported exclusively")
}
