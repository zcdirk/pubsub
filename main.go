package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cs244b-2020-spring-pubsub/pubsub/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"log"
	"net"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/protobuf/encoding/prototext"
)

const (
	defaultPort = 7476
)

var (
	path = flag.String("config", "", "path to pubsub server config")
)

type pubsubServer struct {
	pb.UnimplementedPubSubServer
}

// Publish a message
func (s *pubsubServer) Publish(ctx context.Context, request *pb.PublishRequest) (*pb.PublishResponse, error) {
	log.Printf("Publish: %v\n", request)
	return service.PersistAndPublishMessage(request.Topic, request.Msg)
}

// Subscribe a topic
func (s *pubsubServer) Subscribe(request *pb.SubscribeRequest, stream pb.PubSub_SubscribeServer) error {
	log.Printf("SubScribe: %v\n", request)
	err := service.PersistTopics(stream, request.Topic)
	select {}
	return err
}

func main() {
	flag.Parse()

	config := pb.ServerConfig{
		Port: defaultPort,
	}

	if *path != "" {
		content, err := ioutil.ReadFile(*path)
		if err != nil {
			log.Fatalf("cannot read server config")
		}

		if prototext.Unmarshal(content, &config) != nil {
			log.Fatalf("cannot parse server config")
		}
	}

	log.Println(config)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	pb.RegisterPubSubServer(server, &pubsubServer{})
	server.Serve(lis)
}
