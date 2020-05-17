all: 

init:
	git config core.hooksPath .githooks

dependencies:
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/golang/protobuf/protoc-gen-go

pb:
	protoc proto/*.proto --go_out=plugins=grpc:.

mock: pb
	mockgen -destination mock/pubsub_mock.go github.com/cs244b-2020-spring-pubsub/pubsub/proto PubSub_SubscribeServer,PubSubServer,PubSubClient

test: mock
	go test -v ./...

server: pb
	go build -v -o bin/pubsub .

client: pb
	go build -v -o bin/pubsubctl ./pubsubctl

docker:
	docker build -t pubsub .

clean:
	rm -rf bin
	rm -rf mock
