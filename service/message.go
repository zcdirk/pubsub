package service

import (
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"log"
)

type MessageHistoryEntry struct {
	message pb.Message
	topic   pb.Topic
}

var messageHistoryQueue []MessageHistoryEntry

func PersistAndPublishMessage(topic *pb.Topic, message *pb.Message) (*pb.PublishResponse, error) {
	// Write to history
	messageHistoryQueue = append(messageHistoryQueue, MessageHistoryEntry{message: *message, topic: *topic})

	// Publish
	for _, stream := range GetServersForTopic(topic.Name) {
		err := stream.Send(&pb.SubscribeResponse{Msg: message})
		if err != nil {
			log.Fatalf("failed to publish: %v", err)
		}
	}

	return &pb.PublishResponse{Status: pb.PublishResponse_OK}, nil
}
