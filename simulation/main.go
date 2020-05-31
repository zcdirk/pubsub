package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

var (
	word  = strings.Split("cs244b-2020-spring-pubsub", "")
	topic = &pb.Topic{Name: "topic"}

	step    = flag.Int("step", 100, "Number of clients added to PubSub each time")
	upper   = flag.Int("upper", 5000, "Maximum number of clients to be handled")
	timeout = flag.String("timeout", "2s", "Timeout for each operation")
)

type clientResult struct {
	incorrect bool
}

type batchResult struct {
	size      int
	succeeded int
	incorrect int
	timeOut   int
	length    time.Duration
}

func launchClient(sub pb.PubSubClient, r chan *clientResult) {
	stream, err := sub.Subscribe(context.Background(), &pb.SubscribeRequest{Topic: []*pb.Topic{topic}})
	if err != nil {
		log.Fatalf("cannot connect to server: %s", err)
	}

	for {
		c := &clientResult{}

		for _, w := range word {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("cannot read stream: %s", err)
			}

			if res.Msg.Content != w {
				c.incorrect = true
				break
			}
		}

		r <- c
	}
}

func collectCurrentIterationResult(r batchResult) {
	log.Printf("number of clients: %d", r.size)
	log.Printf("succeeded: %d", r.succeeded)
	log.Printf("incorrect: %d", r.incorrect)
	log.Printf("timed out: %d", r.timeOut)
	log.Printf("total time used: %s", r.length)
}

func main() {
	flag.Parse()

	svr := flag.Args()
	conn := make([]*grpc.ClientConn, len(svr))
	for i, s := range svr {
		c, err := grpc.Dial(
			s,
			grpc.WithTimeout(10*time.Second),
			grpc.WithInsecure())
		defer c.Close()

		if err != nil {
			log.Fatalf("cannot create connection to cluster: %s", err)
		}

		conn[i] = c
	}

	pub := make([]pb.PubSubClient, len(conn))
	sub := make([]pb.PubSubClient, len(conn))
	for i, cc := range conn {
		pub[i] = pb.NewPubSubClient(cc)
		sub[i] = pb.NewPubSubClient(cc)
	}

	to, err := time.ParseDuration(*timeout)
	if err != nil {
		log.Fatalf("cannot parse timeout: %s", err)
	}

	curr := *step
	r := make(chan *clientResult)
	for curr <= *upper {
		for i := 0; i < *step; i++ {
			go launchClient(sub[i%len(sub)], r)
		}

		time.Sleep(2 * time.Second)

		for i, x := range word {
			if _, err := pub[i%len(pub)].Publish(context.Background(), &pb.PublishRequest{
				Topic: topic,
				Msg:   &pb.Message{Content: x},
			}); err != nil {
				log.Fatalf("cannot publish to topic: %s", err)
			}
		}

		b := batchResult{}
		b.size = curr
		start := time.Now()
		for i := 0; i < curr; i++ {
			select {
			case c := <-r:
				b.succeeded++
				if c.incorrect {
					b.incorrect++
				}
			case <-time.After(to):
				b.timeOut++
				start = start.Add(to)
			}
		}
		end := time.Now()
		b.length = end.Sub(start)

		collectCurrentIterationResult(b)

		curr += *step

		time.Sleep(2 * time.Second)
	}
}
