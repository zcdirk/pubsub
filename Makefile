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
	go build -v -o bin/pubsub .

client: pb
	go build -v -o bin/pubsubctl ./pubsubctl

docker:
	docker build -t docker.pkg.github.com/cs244b-2020-spring-pubsub/pubsub/server .

clean: pb
	rm -rf bin
	go mod tidy
