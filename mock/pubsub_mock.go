// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cs244b-2020-spring-pubsub/pubsub/proto (interfaces: PubSub_SubscribeServer,PubSubServer,PubSubClient)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	pubsub "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockPubSub_SubscribeServer is a mock of PubSub_SubscribeServer interface
type MockPubSub_SubscribeServer struct {
	ctrl     *gomock.Controller
	recorder *MockPubSub_SubscribeServerMockRecorder
}

// MockPubSub_SubscribeServerMockRecorder is the mock recorder for MockPubSub_SubscribeServer
type MockPubSub_SubscribeServerMockRecorder struct {
	mock *MockPubSub_SubscribeServer
}

// NewMockPubSub_SubscribeServer creates a new mock instance
func NewMockPubSub_SubscribeServer(ctrl *gomock.Controller) *MockPubSub_SubscribeServer {
	mock := &MockPubSub_SubscribeServer{ctrl: ctrl}
	mock.recorder = &MockPubSub_SubscribeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubSub_SubscribeServer) EXPECT() *MockPubSub_SubscribeServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockPubSub_SubscribeServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockPubSub_SubscribeServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockPubSub_SubscribeServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockPubSub_SubscribeServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockPubSub_SubscribeServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockPubSub_SubscribeServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockPubSub_SubscribeServer) Send(arg0 *pubsub.SubscribeResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockPubSub_SubscribeServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPubSub_SubscribeServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockPubSub_SubscribeServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockPubSub_SubscribeServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockPubSub_SubscribeServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockPubSub_SubscribeServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockPubSub_SubscribeServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockPubSub_SubscribeServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockPubSub_SubscribeServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockPubSub_SubscribeServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockPubSub_SubscribeServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockPubSub_SubscribeServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockPubSub_SubscribeServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockPubSub_SubscribeServer)(nil).SetTrailer), arg0)
}

// MockPubSubServer is a mock of PubSubServer interface
type MockPubSubServer struct {
	ctrl     *gomock.Controller
	recorder *MockPubSubServerMockRecorder
}

// MockPubSubServerMockRecorder is the mock recorder for MockPubSubServer
type MockPubSubServerMockRecorder struct {
	mock *MockPubSubServer
}

// NewMockPubSubServer creates a new mock instance
func NewMockPubSubServer(ctrl *gomock.Controller) *MockPubSubServer {
	mock := &MockPubSubServer{ctrl: ctrl}
	mock.recorder = &MockPubSubServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubSubServer) EXPECT() *MockPubSubServerMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockPubSubServer) Publish(arg0 context.Context, arg1 *pubsub.PublishRequest) (*pubsub.PublishResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(*pubsub.PublishResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish
func (mr *MockPubSubServerMockRecorder) Publish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPubSubServer)(nil).Publish), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockPubSubServer) Subscribe(arg0 *pubsub.SubscribeRequest, arg1 pubsub.PubSub_SubscribeServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockPubSubServerMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockPubSubServer)(nil).Subscribe), arg0, arg1)
}

// MockPubSubClient is a mock of PubSubClient interface
type MockPubSubClient struct {
	ctrl     *gomock.Controller
	recorder *MockPubSubClientMockRecorder
}

// MockPubSubClientMockRecorder is the mock recorder for MockPubSubClient
type MockPubSubClientMockRecorder struct {
	mock *MockPubSubClient
}

// NewMockPubSubClient creates a new mock instance
func NewMockPubSubClient(ctrl *gomock.Controller) *MockPubSubClient {
	mock := &MockPubSubClient{ctrl: ctrl}
	mock.recorder = &MockPubSubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubSubClient) EXPECT() *MockPubSubClientMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockPubSubClient) Publish(arg0 context.Context, arg1 *pubsub.PublishRequest, arg2 ...grpc.CallOption) (*pubsub.PublishResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*pubsub.PublishResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish
func (mr *MockPubSubClientMockRecorder) Publish(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPubSubClient)(nil).Publish), varargs...)
}

// Subscribe mocks base method
func (m *MockPubSubClient) Subscribe(arg0 context.Context, arg1 *pubsub.SubscribeRequest, arg2 ...grpc.CallOption) (pubsub.PubSub_SubscribeClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(pubsub.PubSub_SubscribeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockPubSubClientMockRecorder) Subscribe(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockPubSubClient)(nil).Subscribe), varargs...)
}
