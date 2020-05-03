package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/google/subcommands"
)

type subscribeCommand struct {
	s string
}

func (*subscribeCommand) Name() string {
	return "subscribe"
}

func (*subscribeCommand) Synopsis() string {
	return "Subscribe to messages from topics"
}

func (*subscribeCommand) Usage() string {
	return `subscribe <topic> <topic> ...

  Subcribe to messages from the given list of topics.

`
}

func (_ *subscribeCommand) SetFlags(f *flag.FlagSet) {}

func (c *subscribeCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	req := &pb.SubscribeRequest{}
	for _, t := range f.Args() {
		req.Topic = append(req.Topic, &pb.Topic{Name: t})
	}

	stream, err := createServiceStub(c.s).Subscribe(ctx, req)

	if err != nil {
		log.Fatalf("error ocurred while subscribing to topics: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed receiving message %v", err)
		}

		log.Println(res.Msg)
	}

	return subcommands.ExitSuccess
}
