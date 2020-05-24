all: 

init:
	git config core.hooksPath .githooks

dependencies:
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/golang/protobuf/protoc-gen-go

pb:
	protoc proto/*.proto --go_out=plugins=grpc:.

mock:
	mockgen -destination mock/pubsub_mock.go github.com/cs244b-2020-spring-pubsub/pubsub/proto PubSub_SubscribeServer,PubSubServer,PubSubClient

test: mock
	go test -v ./...

server: pb
	go build -v -o bin/pubsub .

client: pb
	go build -v -o bin/pubsubctl ./pubsubctl

docker:
	docker build -t pubsub .

single-machine-up: 
	docker-compose -f simulation/single-machine.yaml up -d

single-machine-down:
	docker-compose -f simulation/single-machine.yaml down

master-slave-up:
	docker-compose -f simulation/master-slave.yaml up -d

master-slave-down:
	docker-compose -f simulation/master-slave.yaml down

clean:
	rm -rf bin
	rm -rf mock
