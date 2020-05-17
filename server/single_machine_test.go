package server

import (
	"context"
	"testing"
	"time"

	mock "github.com/cs244b-2020-spring-pubsub/pubsub/mock"
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"github.com/golang/mock/gomock"
)

func TestSingleMachineServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	topic1 := pb.Topic{Name: "Topic1"}
	topic2 := pb.Topic{Name: "Topic2"}
	topic3 := pb.Topic{Name: "Topic3"}
	msg1 := pb.Message{Content: "Message1"}
	msg2 := pb.Message{Content: "Message2"}
	msg3 := pb.Message{Content: "Message3"}

	stream1 := mock.NewMockPubSub_SubscribeServer(ctrl)
	stream2 := mock.NewMockPubSub_SubscribeServer(ctrl)

	svr := NewSingleMachineServer()

	stream1.EXPECT().Send(&pb.SubscribeResponse{Msg: &msg1}).Return(nil)
	stream1.EXPECT().Send(&pb.SubscribeResponse{Msg: &msg2}).Return(nil)
	stream2.EXPECT().Send(&pb.SubscribeResponse{Msg: &msg2}).Return(nil)
	stream2.EXPECT().Send(&pb.SubscribeResponse{Msg: &msg3}).Return(nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go svr.Subscribe(&pb.SubscribeRequest{Topic: []*pb.Topic{&topic1}}, stream1)
	go svr.Subscribe(&pb.SubscribeRequest{Topic: []*pb.Topic{&topic2}}, stream1)
	go svr.Subscribe(&pb.SubscribeRequest{Topic: []*pb.Topic{&topic2}}, stream2)
	go svr.Subscribe(&pb.SubscribeRequest{Topic: []*pb.Topic{&topic3}}, stream2)

	// We need to wait some time between 2 actions to avoid concurrency issue
	time.Sleep(time.Second)

	svr.Publish(ctx, &pb.PublishRequest{Topic: &topic1, Msg: &msg1})
	time.Sleep(time.Second)

	svr.Publish(ctx, &pb.PublishRequest{Topic: &topic2, Msg: &msg2})
	time.Sleep(time.Second)

	svr.Publish(ctx, &pb.PublishRequest{Topic: &topic3, Msg: &msg3})
	time.Sleep(time.Second)
}
