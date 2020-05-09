package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/protobuf/encoding/prototext"
)

const (
	defaultPort = 7476
)

var (
	path = flag.String("config", "", "path to pubsub server config")
	topicToChannelsMap = make(map[string][]chan *pb.Message)
	mapLock = new(sync.Mutex)
)

type pubsubServer struct {
	pb.UnimplementedPubSubServer
	quit bool
}

// Publish a message
func (s *pubsubServer) Publish(ctx context.Context, request *pb.PublishRequest) (*pb.PublishResponse, error) {
	log.Printf("Publish: %v\n", request)

	mapLock.Lock()
	channels, ok := topicToChannelsMap[request.Topic.Name]
	mapLock.Unlock()

	message := request.Msg
	if ok {
		for _, channel := range channels {
			channel <- message
		}
	}
	return &pb.PublishResponse{Status: pb.PublishResponse_OK}, nil
}

// Subscribe a topic
func (s *pubsubServer) Subscribe(request *pb.SubscribeRequest, stream pb.PubSub_SubscribeServer) error {
	log.Printf("Subscribe: %v\n", request)
	c := make(chan *pb.Message)
	for _, topic := range request.Topic {
		mapLock.Lock()
		channels, ok := topicToChannelsMap[topic.Name]
		if ok {
			topicToChannelsMap[topic.Name] = append(channels, c)
		} else {
			topicToChannelsMap[topic.Name] = []chan *pb.Message{c}
		}
		mapLock.Unlock()
	}
	return listenAndRespond(c, stream)
}

func listenAndRespond(c chan *pb.Message, stream pb.PubSub_SubscribeServer) error {
	for {
		msg:= <-c
		log.Println(msg.Content)
		err := stream.Send(&pb.SubscribeResponse{Msg: msg})
		if err != nil {
			log.Printf("failed to publish: %v\n", err)
			return err
		}
	}
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
