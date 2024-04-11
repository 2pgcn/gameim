// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: protocol/queue.proto

package protocol

import (
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

type Type int32

const (
	Type_APP   Type = 0
	Type_ROOM  Type = 1
	Type_PUSH  Type = 2
	Type_CLOSE Type = 3
	Type_ACK   Type = 4
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "APP",
		1: "ROOM",
		2: "PUSH",
		3: "CLOSE",
		4: "ACK",
	}
	Type_value = map[string]int32{
		"APP":   0,
		"ROOM":  1,
		"PUSH":  2,
		"CLOSE": 3,
		"ACK":   4,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_queue_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_protocol_queue_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_protocol_queue_proto_rawDescGZIP(), []int{0}
}

// comet后续消息格式
type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Type   `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.Type" json:"type,omitempty"`
	ToId   string `protobuf:"bytes,2,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	SendId string `protobuf:"bytes,3,opt,name=send_id,json=sendId,proto3" json:"send_id,omitempty"`
	Msg    []byte `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_protocol_queue_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_APP
}

func (x *Msg) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *Msg) GetSendId() string {
	if x != nil {
		return x.SendId
	}
	return ""
}

func (x *Msg) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

// reply回复消息
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  *Msg  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_protocol_queue_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_protocol_queue_proto protoreflect.FileDescriptor

var file_protocol_queue_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x22, 0x69, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x74,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3c, 0x0a, 0x05, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x37, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f,
	0x4f, 0x4d, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b,
	0x10, 0x04, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x32, 0x70, 0x67, 0x63, 0x6e, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_queue_proto_rawDescOnce sync.Once
	file_protocol_queue_proto_rawDescData = file_protocol_queue_proto_rawDesc
)

func file_protocol_queue_proto_rawDescGZIP() []byte {
	file_protocol_queue_proto_rawDescOnce.Do(func() {
		file_protocol_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_queue_proto_rawDescData)
	})
	return file_protocol_queue_proto_rawDescData
}

var file_protocol_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_queue_proto_goTypes = []interface{}{
	(Type)(0),     // 0: protocol.Type
	(*Msg)(nil),   // 1: protocol.Msg
	(*Reply)(nil), // 2: protocol.Reply
}
var file_protocol_queue_proto_depIdxs = []int32{
	0, // 0: protocol.Msg.type:type_name -> protocol.Type
	1, // 1: protocol.Reply.msg:type_name -> protocol.Msg
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_queue_proto_init() }
func file_protocol_queue_proto_init() {
	if File_protocol_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_protocol_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_protocol_queue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_queue_proto_goTypes,
		DependencyIndexes: file_protocol_queue_proto_depIdxs,
		EnumInfos:         file_protocol_queue_proto_enumTypes,
		MessageInfos:      file_protocol_queue_proto_msgTypes,
	}.Build()
	File_protocol_queue_proto = out.File
	file_protocol_queue_proto_rawDesc = nil
	file_protocol_queue_proto_goTypes = nil
	file_protocol_queue_proto_depIdxs = nil
}
