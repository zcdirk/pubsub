package main

import (
	"flag"
	"io/ioutil"
	"log"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
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
}
