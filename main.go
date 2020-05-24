package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/cs244b-2020-spring-pubsub/pubsub/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/encoding/prototext"
)

var (
	path = flag.String("config", "", "path to pubsub server config")
)

func main() {
	flag.Parse()

	cfg := &pb.ServerConfig{}

	log.Printf("config path: %s", *path)
	content, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalf("cannot read server config")
	}

	if prototext.Unmarshal(content, cfg) != nil {
		log.Fatalf("cannot parse server config")
	}

	log.Printf("server config: %s", cfg)

	svr, err := server.CreatePubsubServer(cfg,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}))
	if err != nil {
		log.Fatalf("failed to configure server: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svr.Serve(lis)
}
