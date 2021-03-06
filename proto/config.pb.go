// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: proto/config.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServerConfig_ReplicationMode int32

const (
	ServerConfig_SINGLE_MACHINE ServerConfig_ReplicationMode = 0
	ServerConfig_MASTER_SLAVE   ServerConfig_ReplicationMode = 1
	ServerConfig_RAFT           ServerConfig_ReplicationMode = 2
)

// Enum value maps for ServerConfig_ReplicationMode.
var (
	ServerConfig_ReplicationMode_name = map[int32]string{
		0: "SINGLE_MACHINE",
		1: "MASTER_SLAVE",
		2: "RAFT",
	}
	ServerConfig_ReplicationMode_value = map[string]int32{
		"SINGLE_MACHINE": 0,
		"MASTER_SLAVE":   1,
		"RAFT":           2,
	}
)

func (x ServerConfig_ReplicationMode) Enum() *ServerConfig_ReplicationMode {
	p := new(ServerConfig_ReplicationMode)
	*p = x
	return p
}

func (x ServerConfig_ReplicationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerConfig_ReplicationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_proto_enumTypes[0].Descriptor()
}

func (ServerConfig_ReplicationMode) Type() protoreflect.EnumType {
	return &file_proto_config_proto_enumTypes[0]
}

func (x ServerConfig_ReplicationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerConfig_ReplicationMode.Descriptor instead.
func (ServerConfig_ReplicationMode) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

type ServerConfig_MasterSlaveConfig_NodeMode int32

const (
	ServerConfig_MasterSlaveConfig_MASTER ServerConfig_MasterSlaveConfig_NodeMode = 0
	ServerConfig_MasterSlaveConfig_SLAVE  ServerConfig_MasterSlaveConfig_NodeMode = 1
)

// Enum value maps for ServerConfig_MasterSlaveConfig_NodeMode.
var (
	ServerConfig_MasterSlaveConfig_NodeMode_name = map[int32]string{
		0: "MASTER",
		1: "SLAVE",
	}
	ServerConfig_MasterSlaveConfig_NodeMode_value = map[string]int32{
		"MASTER": 0,
		"SLAVE":  1,
	}
)

func (x ServerConfig_MasterSlaveConfig_NodeMode) Enum() *ServerConfig_MasterSlaveConfig_NodeMode {
	p := new(ServerConfig_MasterSlaveConfig_NodeMode)
	*p = x
	return p
}

func (x ServerConfig_MasterSlaveConfig_NodeMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerConfig_MasterSlaveConfig_NodeMode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_proto_enumTypes[1].Descriptor()
}

func (ServerConfig_MasterSlaveConfig_NodeMode) Type() protoreflect.EnumType {
	return &file_proto_config_proto_enumTypes[1]
}

func (x ServerConfig_MasterSlaveConfig_NodeMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerConfig_MasterSlaveConfig_NodeMode.Descriptor instead.
func (ServerConfig_MasterSlaveConfig_NodeMode) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0, 0, 0}
}

// PubSub server config
type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PubSub service port
	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// PubSub service replication mode
	ReplicationMode ServerConfig_ReplicationMode `protobuf:"varint,2,opt,name=replication_mode,json=replicationMode,proto3,enum=pubsub.ServerConfig_ReplicationMode" json:"replication_mode,omitempty"`
	// Config fragment for PubSub if
	// the service is replicated through
	// master-slave mode.
	MasterSlaveConfig *ServerConfig_MasterSlaveConfig `protobuf:"bytes,3,opt,name=master_slave_config,json=masterSlaveConfig,proto3" json:"master_slave_config,omitempty"`
	// Config fragment for PubSub if the service is replicated through raft mode.
	RaftConfig *ServerConfig_RaftConfig `protobuf:"bytes,4,opt,name=raft_config,json=raftConfig,proto3" json:"raft_config,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConfig) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServerConfig) GetReplicationMode() ServerConfig_ReplicationMode {
	if x != nil {
		return x.ReplicationMode
	}
	return ServerConfig_SINGLE_MACHINE
}

func (x *ServerConfig) GetMasterSlaveConfig() *ServerConfig_MasterSlaveConfig {
	if x != nil {
		return x.MasterSlaveConfig
	}
	return nil
}

func (x *ServerConfig) GetRaftConfig() *ServerConfig_RaftConfig {
	if x != nil {
		return x.RaftConfig
	}
	return nil
}

type ServerConfig_MasterSlaveConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// In master-salve mode, a node can
	// be either master or a slave.
	Mode ServerConfig_MasterSlaveConfig_NodeMode `protobuf:"varint,1,opt,name=mode,proto3,enum=pubsub.ServerConfig_MasterSlaveConfig_NodeMode" json:"mode,omitempty"`
	// Address of its master node
	// if current node is a slave.
	MasterAddress string `protobuf:"bytes,2,opt,name=master_address,json=masterAddress,proto3" json:"master_address,omitempty"`
	// Port on which master node hosts
	// PubSub service.
	MasterPort int32 `protobuf:"varint,3,opt,name=master_port,json=masterPort,proto3" json:"master_port,omitempty"`
	// Timeout of connection from slave
	// to master.
	MasterTimeout string `protobuf:"bytes,4,opt,name=master_timeout,json=masterTimeout,proto3" json:"master_timeout,omitempty"`
}

func (x *ServerConfig_MasterSlaveConfig) Reset() {
	*x = ServerConfig_MasterSlaveConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig_MasterSlaveConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig_MasterSlaveConfig) ProtoMessage() {}

func (x *ServerConfig_MasterSlaveConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig_MasterSlaveConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig_MasterSlaveConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerConfig_MasterSlaveConfig) GetMode() ServerConfig_MasterSlaveConfig_NodeMode {
	if x != nil {
		return x.Mode
	}
	return ServerConfig_MasterSlaveConfig_MASTER
}

func (x *ServerConfig_MasterSlaveConfig) GetMasterAddress() string {
	if x != nil {
		return x.MasterAddress
	}
	return ""
}

func (x *ServerConfig_MasterSlaveConfig) GetMasterPort() int32 {
	if x != nil {
		return x.MasterPort
	}
	return 0
}

func (x *ServerConfig_MasterSlaveConfig) GetMasterTimeout() string {
	if x != nil {
		return x.MasterTimeout
	}
	return ""
}

type ServerConfig_RaftConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of itself
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Peer address
	Peers []string `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	// Timeout of peer connection.
	PeerTimeout string `protobuf:"bytes,3,opt,name=peer_timeout,json=peerTimeout,proto3" json:"peer_timeout,omitempty"`
	// Interval of heartbeat.
	HeartbeatInterval string `protobuf:"bytes,4,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
}

func (x *ServerConfig_RaftConfig) Reset() {
	*x = ServerConfig_RaftConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig_RaftConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig_RaftConfig) ProtoMessage() {}

func (x *ServerConfig_RaftConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig_RaftConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig_RaftConfig) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ServerConfig_RaftConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServerConfig_RaftConfig) GetPeers() []string {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *ServerConfig_RaftConfig) GetPeerTimeout() string {
	if x != nil {
		return x.PeerTimeout
	}
	return ""
}

func (x *ServerConfig_RaftConfig) GetHeartbeatInterval() string {
	if x != nil {
		return x.HeartbeatInterval
	}
	return ""
}

var File_proto_config_proto protoreflect.FileDescriptor

var file_proto_config_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x22, 0xc4, 0x05, 0x0a,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x4f, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x56, 0x0a, 0x13, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x61,
	0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x6c, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40, 0x0a, 0x0b, 0x72, 0x61,
	0x66, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xea, 0x01, 0x0a,
	0x11, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x43, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6c, 0x61,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x21, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x4c, 0x41, 0x56, 0x45, 0x10, 0x01, 0x1a, 0x84, 0x01, 0x0a, 0x0a, 0x52, 0x61,
	0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x2d, 0x0a, 0x12, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x22, 0x41, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x4d, 0x41,
	0x43, 0x48, 0x49, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x41, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x53, 0x4c, 0x41, 0x56, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x41, 0x46,
	0x54, 0x10, 0x02, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_config_proto_rawDescOnce sync.Once
	file_proto_config_proto_rawDescData = file_proto_config_proto_rawDesc
)

func file_proto_config_proto_rawDescGZIP() []byte {
	file_proto_config_proto_rawDescOnce.Do(func() {
		file_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_proto_rawDescData)
	})
	return file_proto_config_proto_rawDescData
}

var file_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_config_proto_goTypes = []interface{}{
	(ServerConfig_ReplicationMode)(0),            // 0: pubsub.ServerConfig.ReplicationMode
	(ServerConfig_MasterSlaveConfig_NodeMode)(0), // 1: pubsub.ServerConfig.MasterSlaveConfig.NodeMode
	(*ServerConfig)(nil),                         // 2: pubsub.ServerConfig
	(*ServerConfig_MasterSlaveConfig)(nil),       // 3: pubsub.ServerConfig.MasterSlaveConfig
	(*ServerConfig_RaftConfig)(nil),              // 4: pubsub.ServerConfig.RaftConfig
}
var file_proto_config_proto_depIdxs = []int32{
	0, // 0: pubsub.ServerConfig.replication_mode:type_name -> pubsub.ServerConfig.ReplicationMode
	3, // 1: pubsub.ServerConfig.master_slave_config:type_name -> pubsub.ServerConfig.MasterSlaveConfig
	4, // 2: pubsub.ServerConfig.raft_config:type_name -> pubsub.ServerConfig.RaftConfig
	1, // 3: pubsub.ServerConfig.MasterSlaveConfig.mode:type_name -> pubsub.ServerConfig.MasterSlaveConfig.NodeMode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_config_proto_init() }
func file_proto_config_proto_init() {
	if File_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig_MasterSlaveConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig_RaftConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_proto_goTypes,
		DependencyIndexes: file_proto_config_proto_depIdxs,
		EnumInfos:         file_proto_config_proto_enumTypes,
		MessageInfos:      file_proto_config_proto_msgTypes,
	}.Build()
	File_proto_config_proto = out.File
	file_proto_config_proto_rawDesc = nil
	file_proto_config_proto_goTypes = nil
	file_proto_config_proto_depIdxs = nil
}
