package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/google/subcommands"
	"google.golang.org/grpc"
)

var (
	address = flag.String("address", "localhost", "pubsub server address")
	port    = flag.Int("port", 7476, "pubsub server port")
	timeout = flag.String("timeout", "10s", "pubsub server connection time out")
)

func createServiceStub(server string) pb.PubSubClient {
	connTimeout, err := time.ParseDuration(*timeout)
	if err != nil {
		log.Fatalf("invalid timeout definition: %s", *timeout)
	}

	conn, err := grpc.Dial(server,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(connTimeout))

	if err != nil {
		log.Fatalf("cannot connect to server [%s]: %v", server, err)
	}

	return pb.NewPubSubClient(conn)
}

func main() {
	flag.Parse()

	server := fmt.Sprintf("%s:%d", *address, *port)

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&publishCommand{s: server}, "")
	subcommands.Register(&subscribeCommand{s: server}, "")

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
