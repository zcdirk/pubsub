package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/google/subcommands"
)

type publishCommand struct {
	s string
	t pb.Topic
}

func (*publishCommand) Name() string {
	return "publish"
}

func (*publishCommand) Synopsis() string {
	return "Publish message to topic"
}

func (*publishCommand) Usage() string {
	return `publish --topic=<topic> <message>

  Publish message to the topic

`
}

func (c *publishCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.t.Name, "topic", "", "topic where to publish the message")
}

func (c *publishCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	req := &pb.PublishRequest{}
	req.Topic = &c.t
	req.Msg = &pb.Message{Content: f.Args()[0]}

	res, err := createServiceStub(c.s).Publish(ctx, req)

	if err != nil {
		log.Fatalf("error ocurred while publishing message: %v", err)
	}

	switch res.Status.(type) {
	case *pb.PublishResponse_Failure_:
		log.Fatalf("publish request failed: %v", res.GetFailure().Reason)
	}

	return subcommands.ExitSuccess
}
