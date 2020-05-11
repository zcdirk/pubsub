all: 

dependencies:
	go get -u github.com/golang/protobuf/protoc-gen-go

pb: proto/*.proto
	protoc proto/*.proto --go_out=plugins=grpc:.

test: pb
	go test -v ./...

server: pb
	go build -v -o bin/pubsub .

client: pb
	go build -v -o bin/pubsubctl ./pubsubctl

docker:
	docker build -t pubsub .

clean: pb
	go mod tidy
	rm -rf bin
