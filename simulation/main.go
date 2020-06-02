package main

import (
	"context"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

var (
	word  = strings.Split("cs244b-2020-spring-pubsub", "")
	topic = &pb.Topic{Name: "topic"}

	step    = flag.Int("step", 50, "Number of clients added to PubSub each time")
	lower   = flag.Int("lower", 0, "Lower bound of number of clients to be handled")
	upper   = flag.Int("upper", 2000, "Upper bound number of clients to be handled")
	timeout = flag.String("timeout", "10s", "Timeout for each operation")
	output  = flag.String("output", "result.csv", "Path to export statistics result.")
)

type clientResult struct {
	incorrect bool
}

type batchResult struct {
	size      int
	succeeded int
	incorrect int
	timeOut   int
	duration    time.Duration
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
				return
			}

			if res.Msg.Content != w {
				log.Printf("incorrect message: got %s, want %s", res.Msg.Content, w)
				c.incorrect = true
				break
			}
		}
		r <- c
	}
}

func collectCurrentIterationResult(r batchResult, o *csv.Writer) {
	log.Printf("number of clients: %d", r.size)
	log.Printf("succeeded: %d", r.succeeded)
	log.Printf("incorrect: %d", r.incorrect)
	log.Printf("timed out: %d", r.timeOut)
	log.Printf("total time used: %s", r.duration)

	err := o.Write([]string{
		strconv.Itoa(r.size),
		strconv.Itoa(r.succeeded),
		strconv.Itoa(r.timeOut),
		r.duration.String(),
	})
	if err != nil {
		log.Fatalf("cannot write to csv: %s", err)
	}

	o.Flush()
	if o.Error() != nil {
		log.Fatalf("cannot write to csv: %s", err)
	}
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

	f, err := os.Create(*output)
	defer f.Close()
	if err != nil {
		log.Fatalf("cannot create output file: %s", err)
	}

	o := csv.NewWriter(f)
	o.Write([]string{
		"size",
		"succeeded",
		"timeout",
		"duration",
	})

	r := make(chan *clientResult)

	curr := *lower
	for i := 0; i < curr; i++ {
		go launchClient(sub[i%len(sub)], r)
	}

	log.Printf("%d cients have been pre-initiated", curr)

	for curr < *upper {
		for i := curr; i < curr + *step; i++ {
			go launchClient(sub[i%len(sub)], r)
		}

		curr += *step
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
		b.duration = end.Sub(start)

		collectCurrentIterationResult(b, o)

		time.Sleep(2 * time.Second)
	}	
}
