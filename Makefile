all: 

protobuf: proto
	protoc proto/*.proto --go_out=plugins=grpc:.

server: protobuf main.go
	go build -o pubsub-server github.com/cs244b-2020-spring-pubsub/pubsub

client: protobuf client/main.go
	go build -o pubsub-client github.com/cs244b-2020-spring-pubsub/pubsub/client

test: protobuf
	go test github.com/cs244b-2020-spring-pubsub/pubsub/...

clean:
	rm -f proto/*.pb.go
	rm -f pubsub-server
	rm -f pubsub-client
