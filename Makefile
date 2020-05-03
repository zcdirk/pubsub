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

server: pb
	go build -v -o pubsub-server .

client: pb
	go build -v -o pubsubctl ./ctl

docker:
	docker build -t pubsub .

clean:
	rm -f proto/*.pb.go
	rm -f pubsub-*
	rm -f pubsubctl
