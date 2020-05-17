package server

import (
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

// ConfigurePubsubServer configures PubSub server implementation based on use config
func ConfigurePubsubServer(svr *grpc.Server, cfg *pb.ServerConfig) {
	switch cfg.ReplicationMode {
	case pb.ServerConfig_SINGLE_MACHINE:
		pb.RegisterPubSubServer(svr, NewSingleMachineServer())
	}
}
