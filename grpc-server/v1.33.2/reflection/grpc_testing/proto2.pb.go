// Copyright 2017 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.3.0
// source: reflection/grpc_testing/proto2.proto

package grpc_testing

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
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

type ToBeExtended struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	Foo *int32 `protobuf:"varint,1,req,name=foo" json:"foo,omitempty"`
}

func (x *ToBeExtended) Reset() {
	*x = ToBeExtended{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reflection_grpc_testing_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToBeExtended) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToBeExtended) ProtoMessage() {}

func (x *ToBeExtended) ProtoReflect() protoreflect.Message {
	mi := &file_reflection_grpc_testing_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToBeExtended.ProtoReflect.Descriptor instead.
func (*ToBeExtended) Descriptor() ([]byte, []int) {
	return file_reflection_grpc_testing_proto2_proto_rawDescGZIP(), []int{0}
}

var extRange_ToBeExtended = []protoiface.ExtensionRangeV1{
	{Start: 10, End: 30},
}

// Deprecated: Use ToBeExtended.ProtoReflect.Descriptor.ExtensionRanges instead.
func (*ToBeExtended) ExtensionRangeArray() []protoiface.ExtensionRangeV1 {
	return extRange_ToBeExtended
}

func (x *ToBeExtended) GetFoo() int32 {
	if x != nil && x.Foo != nil {
		return *x.Foo
	}
	return 0
}

var File_reflection_grpc_testing_proto2_proto protoreflect.FileDescriptor

var file_reflection_grpc_testing_proto2_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x22, 0x26, 0x0a, 0x0c, 0x54, 0x6f, 0x42, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x2a, 0x04, 0x08, 0x0a, 0x10, 0x1f, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
}

var (
	file_reflection_grpc_testing_proto2_proto_rawDescOnce sync.Once
	file_reflection_grpc_testing_proto2_proto_rawDescData = file_reflection_grpc_testing_proto2_proto_rawDesc
)

func file_reflection_grpc_testing_proto2_proto_rawDescGZIP() []byte {
	file_reflection_grpc_testing_proto2_proto_rawDescOnce.Do(func() {
		file_reflection_grpc_testing_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_reflection_grpc_testing_proto2_proto_rawDescData)
	})
	return file_reflection_grpc_testing_proto2_proto_rawDescData
}

var file_reflection_grpc_testing_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reflection_grpc_testing_proto2_proto_goTypes = []interface{}{
	(*ToBeExtended)(nil), // 0: grpc.testing.ToBeExtended
}
var file_reflection_grpc_testing_proto2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reflection_grpc_testing_proto2_proto_init() }
func file_reflection_grpc_testing_proto2_proto_init() {
	if File_reflection_grpc_testing_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reflection_grpc_testing_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToBeExtended); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reflection_grpc_testing_proto2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reflection_grpc_testing_proto2_proto_goTypes,
		DependencyIndexes: file_reflection_grpc_testing_proto2_proto_depIdxs,
		MessageInfos:      file_reflection_grpc_testing_proto2_proto_msgTypes,
	}.Build()
	File_reflection_grpc_testing_proto2_proto = out.File
	file_reflection_grpc_testing_proto2_proto_rawDesc = nil
	file_reflection_grpc_testing_proto2_proto_goTypes = nil
	file_reflection_grpc_testing_proto2_proto_depIdxs = nil
}