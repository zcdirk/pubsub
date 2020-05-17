package server

import (
	"fmt"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

// ConfigurePubsubServer configures PubSub server implementation based on use config
func ConfigurePubsubServer(svr *grpc.Server, cfg *pb.ServerConfig) error {
	switch cfg.ReplicationMode {
	case pb.ServerConfig_SINGLE_MACHINE:
		pb.RegisterPubSubServer(svr, NewSingleMachineServer())
		return nil

	case pb.ServerConfig_MASTER_SLAVE:
		switch msCfg := cfg.MasterSlaveConfig; msCfg.Mode {
		case pb.ServerConfig_MasterSlaveConfig_MASTER:
			pb.RegisterPubSubServer(svr, NewMasterServer())

		case pb.ServerConfig_MasterSlaveConfig_SLAVE:
			connTimeout, err := time.ParseDuration(msCfg.MasterTimeout)
			if err != nil {
				return err
			}

			conn, err := grpc.Dial(
				fmt.Sprintf("%s:%d", msCfg.MasterAddress, msCfg.MasterPort),
				grpc.WithTimeout(connTimeout),
				grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("cannot connect to master: %v", err)
			}

			pb.RegisterPubSubServer(svr, NewSlaveServer(pb.NewPubSubClient(conn)))
		}

		return nil
	}

	return fmt.Errorf("current user config is not supported exclusively")
}
