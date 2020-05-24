package benchmarks

import (
	"testing"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

var (
	n     = 100
	want  = n * (n + 1) / 2
	topic = &pb.Topic{Name: "sum"}
)

func createPubSubConn(b *testing.B) *grpc.ClientConn {
	conn, err := grpc.Dial(
		":7476",
		grpc.WithTimeout(10*time.Second),
		grpc.WithInsecure())

	if err != nil {
		b.Fatalf("cannot create connection to cluster: %s", err)
	}

	return conn
}
