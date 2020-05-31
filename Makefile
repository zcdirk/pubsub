all: 

init:
	git config core.hooksPath .githooks

dependencies:
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/golang/protobuf/protoc-gen-go

pb: dependencies
	protoc proto/*.proto --go_out=plugins=grpc:.

mock: pb
	mockgen -destination mock/pubsub_mock.go github.com/cs244b-2020-spring-pubsub/pubsub/proto PubSub_SubscribeServer,PubSubServer,PubSubClient

test: mock
	go test -v ./...

server: pb main.go
	go build -v -o bin/pubsub .

pubsubctl: pb pubsubctl/main.go
	go build -v -o bin/pubsubctl ./pubsubctl

simulation: pb simulation/main.go
	go build -v -o bin/simulation ./simulation

docker:
	docker build -t pubsub .

single-machine-up: 
	docker-compose -f simulation/single-machine.yaml up -d

single-machine-down:
	docker-compose -f simulation/single-machine.yaml down

single-machine-restart:
	docker-compose -f simulation/single-machine.yaml restart

master-slave-up:
	docker-compose -f simulation/master-slave.yaml up -d

master-slave-down:
	docker-compose -f simulation/master-slave.yaml down

master-slave-restart:
	docker-compose -f simulation/master-slave.yaml restart

raft-up:
	docker-compose -f simulation/raft.yaml up -d

raft-down:
	docker-compose -f simulation/raft.yaml down

raft-restart:
	docker-compose -f simulation/raft.yaml restart

clean:
	rm -rf bin
	rm -rf mock
