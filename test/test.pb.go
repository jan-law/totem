// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v1.0.0
// source: github.com/kralicky/totem/test/test.proto

package test

import (
	totem "github.com/kralicky/totem"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnError bool `protobuf:"varint,1,opt,name=ReturnError,proto3" json:"ReturnError,omitempty"`
}

func (x *ErrorRequest) Reset() {
	*x = ErrorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorRequest) ProtoMessage() {}

func (x *ErrorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorRequest.ProtoReflect.Descriptor instead.
func (*ErrorRequest) Descriptor() ([]byte, []int) {
	return file_github_com_kralicky_totem_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorRequest) GetReturnError() bool {
	if x != nil {
		return x.ReturnError
	}
	return false
}

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_github_com_kralicky_totem_test_test_proto_rawDescGZIP(), []int{1}
}

func (x *Number) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=Str,proto3" json:"Str,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_github_com_kralicky_totem_test_test_proto_rawDescGZIP(), []int{2}
}

func (x *String) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type Bytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Bytes) Reset() {
	*x = Bytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytes) ProtoMessage() {}

func (x *Bytes) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bytes.ProtoReflect.Descriptor instead.
func (*Bytes) Descriptor() ([]byte, []int) {
	return file_github_com_kralicky_totem_test_test_proto_rawDescGZIP(), []int{3}
}

func (x *Bytes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Operands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (x *Operands) Reset() {
	*x = Operands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operands) ProtoMessage() {}

func (x *Operands) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kralicky_totem_test_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operands.ProtoReflect.Descriptor instead.
func (*Operands) Descriptor() ([]byte, []int) {
	return file_github_com_kralicky_totem_test_test_proto_rawDescGZIP(), []int{4}
}

func (x *Operands) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Operands) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_github_com_kralicky_totem_test_test_proto protoreflect.FileDescriptor

var file_github_com_kralicky_totem_test_test_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61,
	0x6c, 0x69, 0x63, 0x6b, 0x79, 0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x6c, 0x69,
	0x63, 0x6b, 0x79, 0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x6c, 0x69, 0x63, 0x6b, 0x79, 0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x30, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x1e, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x74, 0x72, 0x22,
	0x1b, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x08,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x41, 0x12, 0x0c, 0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x42, 0x32, 0x30, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0a,
	0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0a, 0x2e, 0x74, 0x6f, 0x74,
	0x65, 0x6d, 0x2e, 0x52, 0x50, 0x43, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x50, 0x43, 0x28, 0x01, 0x30, 0x01, 0x32, 0x2e, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x49, 0x6e, 0x63, 0x12, 0x0c, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x2e, 0x0a, 0x09, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x44, 0x65, 0x63, 0x12, 0x0c, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x2a, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22,
	0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x32, 0x40, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x36, 0x0a, 0x06,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x8a, 0xf1,
	0x04, 0x02, 0x08, 0x01, 0x32, 0x2f, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79,
	0x12, 0x23, 0x0a, 0x03, 0x4d, 0x75, 0x6c, 0x12, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x3c, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0x28, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x1a, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x20, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x6c,
	0x69, 0x63, 0x6b, 0x79, 0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_kralicky_totem_test_test_proto_rawDescOnce sync.Once
	file_github_com_kralicky_totem_test_test_proto_rawDescData = file_github_com_kralicky_totem_test_test_proto_rawDesc
)

func file_github_com_kralicky_totem_test_test_proto_rawDescGZIP() []byte {
	file_github_com_kralicky_totem_test_test_proto_rawDescOnce.Do(func() {
		file_github_com_kralicky_totem_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_kralicky_totem_test_test_proto_rawDescData)
	})
	return file_github_com_kralicky_totem_test_test_proto_rawDescData
}

var file_github_com_kralicky_totem_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_kralicky_totem_test_test_proto_goTypes = []interface{}{
	(*ErrorRequest)(nil),  // 0: test.ErrorRequest
	(*Number)(nil),        // 1: test.Number
	(*String)(nil),        // 2: test.String
	(*Bytes)(nil),         // 3: test.Bytes
	(*Operands)(nil),      // 4: test.Operands
	(*totem.RPC)(nil),     // 5: totem.RPC
	(*emptypb.Empty)(nil), // 6: google.protobuf.Empty
}
var file_github_com_kralicky_totem_test_test_proto_depIdxs = []int32{
	5, // 0: test.Test.TestStream:input_type -> totem.RPC
	1, // 1: test.Increment.Inc:input_type -> test.Number
	1, // 2: test.Decrement.Dec:input_type -> test.Number
	2, // 3: test.Hash.Hash:input_type -> test.String
	2, // 4: test.Notify.Notify:input_type -> test.String
	4, // 5: test.Multiply.Mul:input_type -> test.Operands
	0, // 6: test.Error.Error:input_type -> test.ErrorRequest
	3, // 7: test.Echo.Echo:input_type -> test.Bytes
	5, // 8: test.Test.TestStream:output_type -> totem.RPC
	1, // 9: test.Increment.Inc:output_type -> test.Number
	1, // 10: test.Decrement.Dec:output_type -> test.Number
	2, // 11: test.Hash.Hash:output_type -> test.String
	6, // 12: test.Notify.Notify:output_type -> google.protobuf.Empty
	1, // 13: test.Multiply.Mul:output_type -> test.Number
	6, // 14: test.Error.Error:output_type -> google.protobuf.Empty
	3, // 15: test.Echo.Echo:output_type -> test.Bytes
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_kralicky_totem_test_test_proto_init() }
func file_github_com_kralicky_totem_test_test_proto_init() {
	if File_github_com_kralicky_totem_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_kralicky_totem_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorRequest); i {
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
		file_github_com_kralicky_totem_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_github_com_kralicky_totem_test_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
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
		file_github_com_kralicky_totem_test_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bytes); i {
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
		file_github_com_kralicky_totem_test_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operands); i {
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
			RawDescriptor: file_github_com_kralicky_totem_test_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   8,
		},
		GoTypes:           file_github_com_kralicky_totem_test_test_proto_goTypes,
		DependencyIndexes: file_github_com_kralicky_totem_test_test_proto_depIdxs,
		MessageInfos:      file_github_com_kralicky_totem_test_test_proto_msgTypes,
	}.Build()
	File_github_com_kralicky_totem_test_test_proto = out.File
	file_github_com_kralicky_totem_test_test_proto_rawDesc = nil
	file_github_com_kralicky_totem_test_test_proto_goTypes = nil
	file_github_com_kralicky_totem_test_test_proto_depIdxs = nil
}
