package benchmarks

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

var (
	n     = 10
	want  = n * (n - 1) / 2
	topic = &pb.Topic{Name: "sum"}
)

func createPubSubConn(b *testing.B, svr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		svr,
		grpc.WithTimeout(10*time.Second),
		grpc.WithInsecure())

	if err != nil {
		b.Fatalf("cannot create connection to cluster: %s", err)
	}

	return conn
}

func createPubSubBenchmark(b *testing.B, svr []string) {
	conn := make([]*grpc.ClientConn, len(svr))
	for i, s := range svr {
		conn[i] = createPubSubConn(b, s)
	}

	ec := make(chan error)
	for i := 0; i < b.N; i++ {
		sub := pb.NewPubSubClient(conn[i%len(svr)])
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

	for i := 0; i < n; i++ {
		pub := pb.NewPubSubClient(conn[i%len(svr)])
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
