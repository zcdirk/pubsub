package service

import (
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
)

var topicToServersMap = make(map[string][]pb.PubSub_SubscribeServer)

func PersistTopics(stream pb.PubSub_SubscribeServer, topics []*pb.Topic) error {
	for _, topic := range topics {
		value, ok := topicToServersMap[topic.Name]
		if ok {
			topicToServersMap[topic.Name] = append(value, stream)
		} else {
			topicToServersMap[topic.Name] = []pb.PubSub_SubscribeServer{stream}
		}
	}
	return nil
}

func GetServersForTopic(topicName string) []pb.PubSub_SubscribeServer {
	return topicToServersMap[topicName]
}
