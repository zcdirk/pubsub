all: 

init:
	git config core.hooksPath .githooks

dependencies:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get ./...

pb: proto/*.proto
	protoc proto/*.proto --go_out=plugins=grpc:.

test: pb
	go test -v ./...

server: pb main.go
	go build -v -o pubsub-server .

client: pb client/main.go
	go build -v -o pubsub-client ./client

docker:
	docker build -t pubsub .

clean:
	rm -f proto/*.pb.go
	rm -f pubsub-*
