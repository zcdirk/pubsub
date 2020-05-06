package service

import (
	mock_proto "github.com/cs244b-2020-spring-pubsub/pubsub/mock"
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPersistAndPublishMessage(t *testing.T) {
	topicToServersMap = make(map[string][]pb.PubSub_SubscribeServer)
	messageHistoryQueue = nil

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream1 := mock_proto.NewMockPubSub_SubscribeServer(ctrl)
	stream2 := mock_proto.NewMockPubSub_SubscribeServer(ctrl)

	msg1 := pb.Message{Content: "Message1"}
	msg2 := pb.Message{Content: "Message2"}
	msg3 := pb.Message{Content: "Message3"}

	stream1.EXPECT().Send(&pb.SubscribeResponse{Msg: &msg1}).Return(nil)
	stream1.EXPECT().Send(&pb.SubscribeResponse{Msg: &msg2}).Return(nil)
	stream2.EXPECT().Send(&pb.SubscribeResponse{Msg: &msg3}).Return(nil)

	PersistTopics(stream1, []*pb.Topic{{Name: "Topic1"}})
	PersistTopics(stream2, []*pb.Topic{{Name: "Topic2"}})

	PersistAndPublishMessage(&pb.Topic{Name: "Topic1"}, &msg1)
	PersistAndPublishMessage(&pb.Topic{Name: "Topic1"}, &msg2)
	PersistAndPublishMessage(&pb.Topic{Name: "Topic2"}, &msg3)

	if len(messageHistoryQueue) != 3 ||
		messageHistoryQueue[0].message.Content != msg1.Content ||
		messageHistoryQueue[1].message.Content != msg2.Content ||
		messageHistoryQueue[2].message.Content != msg3.Content {
		t.Error()
	}
}
