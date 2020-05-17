package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/cs244b-2020-spring-pubsub/pubsub/server"
	"google.golang.org/protobuf/encoding/prototext"
)

const (
	defaultPort = 7476
)

var (
	path = flag.String("config", "", "path to pubsub server config")
)

func main() {
	flag.Parse()

	cfg := &pb.ServerConfig{
		Port: defaultPort,
	}

	content, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalf("cannot read server config")
	}

	if prototext.Unmarshal(content, cfg) != nil {
		log.Fatalf("cannot parse server config")
	}

	svr := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))

	server.ConfigurePubsubServer(svr, cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svr.Serve(lis)
}
