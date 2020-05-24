package benchmarks

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

var (
	word  = strings.Split("pubsub", "")
	topic = &pb.Topic{Name: "topic"}
)

func runPubSubBenchmark(b *testing.B, svr []string) {
	conn := make([]*grpc.ClientConn, len(svr))
	for i, s := range svr {
		c, err := grpc.Dial(
			s,
			grpc.WithTimeout(10*time.Second),
			grpc.WithInsecure())

		if err != nil {
			b.Fatalf("cannot create connection to cluster: %s", err)
		}

		conn[i] = c
	}

	wg := &sync.WaitGroup{}
	wg.Add(b.N)

	for s := 0; s < b.N; s++ {
		go func(sub pb.PubSubClient) {
			defer wg.Done()

			stream, err := sub.Subscribe(context.Background(), &pb.SubscribeRequest{Topic: []*pb.Topic{topic}})
			if err != nil {
				panic(err)
			}

			for _, w := range word {
				res, err := stream.Recv()
				if err != nil {
					panic(err)
				}

				if res.Msg.Content != w {
					panic(fmt.Errorf("data was compromised in process, want: %s, but got: %s", w, res.Msg.Content))
				}
			}
		}(pb.NewPubSubClient(conn[s%len(svr)]))
	}

	time.Sleep(5 * time.Second)

	for p, x := range word {
		pub := pb.NewPubSubClient(conn[p%len(svr)])

		if _, err := pub.Publish(context.Background(), &pb.PublishRequest{
			Topic: topic,
			Msg:   &pb.Message{Content: x},
		}); err != nil {
			b.Fatalf("cannot publish to topic: %s", err)
		}
	}

	wg.Wait()
}
