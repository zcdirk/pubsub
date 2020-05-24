package benchmarks

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
)

func BenchmarkSingleMachine(b *testing.B) {
	sub := pb.NewPubSubClient(createPubSubConn(b, ":7476"))
	ec := make(chan error)
	for i := 0; i < b.N; i++ {
		go func() {
			stream, err := sub.Subscribe(context.Background(), &pb.SubscribeRequest{Topic: []*pb.Topic{topic}})
			if err != nil {
				ec <- err
				return
			}

			got := 0
			for j := 0; j < n; j++ {
				res, err := stream.Recv()
				if err != nil {
					ec <- err
					return
				}

				x, _ := strconv.Atoi(res.Msg.Content)
				got += x
			}

			if got != want {
				ec <- fmt.Errorf("data was compromised in process, want: %d, but got: %d", want, got)
				return
			}

			ec <- nil
		}()
	}

	time.Sleep(5 * time.Second)

	b.StartTimer()

	pub := pb.NewPubSubClient(createPubSubConn(b, ":7476"))
	for i := 1; i <= n; i++ {
		if _, err := pub.Publish(context.Background(), &pb.PublishRequest{
			Topic: topic,
			Msg:   &pb.Message{Content: strconv.Itoa(i)},
		}); err != nil {
			b.Fatalf("cannot publish to topic: %s", err)
		}
	}

	for i := 0; i < b.N; i++ {
		if err := <-ec; err != nil {
			b.Fatalf("client error: %s", err)
		}
	}

	b.StopTimer()
}
