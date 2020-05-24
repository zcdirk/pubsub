package benchmarks

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
)

func BenchmarkSingleMachine(b *testing.B) {
	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	want := n * (n + 1) / 2

	for i := 0; i < b.N; i++ {
		go func() {
			sub := pb.NewPubSubClient(createPubSubConn(b))

			stream, err := sub.Subscribe(context.Background(), &pb.SubscribeRequest{Topic: []*pb.Topic{topic}})
			if err != nil {
				b.Fatalf("error connect to server: %s", err)
			}

			got := 0
			for j := 0; j < n; j++ {
				res, err := stream.Recv()
				if err != nil {
					b.Fatalf("error receiving message from server: %s", err)
				}

				x, _ := strconv.Atoi(res.Msg.Content)
				got += x
			}

			if got != want {
				b.Fatalf("data was lost in process: wanted %d, but got %d", want, got)
			}

			wg.Done()
		}()
	}

	time.Sleep(5 * time.Second)

	b.StartTimer()

	pub := pb.NewPubSubClient(createPubSubConn(b))
	for i := 1; i <= n; i++ {
		if _, err := pub.Publish(context.Background(), &pb.PublishRequest{
			Topic: topic,
			Msg:   &pb.Message{Content: strconv.Itoa(i)},
		}); err != nil {
			b.Fatalf("cannot publish to topic: %s", err)
		}
	}

	wg.Wait()

	b.StopTimer()
}
