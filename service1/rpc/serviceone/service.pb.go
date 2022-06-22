// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: service1/rpc/service.proto

package serviceone

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

type Size struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inches int32 `protobuf:"varint,1,opt,name=inches,proto3" json:"inches,omitempty"` // must be > 0
}

func (x *Size) Reset() {
	*x = Size{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service1_rpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size) ProtoMessage() {}

func (x *Size) ProtoReflect() protoreflect.Message {
	mi := &file_service1_rpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size.ProtoReflect.Descriptor instead.
func (*Size) Descriptor() ([]byte, []int) {
	return file_service1_rpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *Size) GetInches() int32 {
	if x != nil {
		return x.Inches
	}
	return 0
}

type Hat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inches int32  `protobuf:"varint,1,opt,name=inches,proto3" json:"inches,omitempty"`
	Color  string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"` // anything but "invisible"
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`   // i.e. "bowler"
}

func (x *Hat) Reset() {
	*x = Hat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service1_rpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hat) ProtoMessage() {}

func (x *Hat) ProtoReflect() protoreflect.Message {
	mi := &file_service1_rpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hat.ProtoReflect.Descriptor instead.
func (*Hat) Descriptor() ([]byte, []int) {
	return file_service1_rpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *Hat) GetInches() int32 {
	if x != nil {
		return x.Inches
	}
	return 0
}

func (x *Hat) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Hat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_service1_rpc_service_proto protoreflect.FileDescriptor

var file_service1_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x7a, 0x61,
	0x69, 0x6e, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x22, 0x1e, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x03, 0x48, 0x61, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x69, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0x58, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x4a,
	0x0a, 0x07, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x7a, 0x61, 0x69, 0x6e,
	0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x31, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x1e, 0x2e, 0x7a, 0x61, 0x69,
	0x6e, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x2e, 0x48, 0x61, 0x74, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service1_rpc_service_proto_rawDescOnce sync.Once
	file_service1_rpc_service_proto_rawDescData = file_service1_rpc_service_proto_rawDesc
)

func file_service1_rpc_service_proto_rawDescGZIP() []byte {
	file_service1_rpc_service_proto_rawDescOnce.Do(func() {
		file_service1_rpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service1_rpc_service_proto_rawDescData)
	})
	return file_service1_rpc_service_proto_rawDescData
}

var file_service1_rpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service1_rpc_service_proto_goTypes = []interface{}{
	(*Size)(nil), // 0: zainokta.example.service1.Size
	(*Hat)(nil),  // 1: zainokta.example.service1.Hat
}
var file_service1_rpc_service_proto_depIdxs = []int32{
	0, // 0: zainokta.example.service1.ServiceOne.MakeHat:input_type -> zainokta.example.service1.Size
	1, // 1: zainokta.example.service1.ServiceOne.MakeHat:output_type -> zainokta.example.service1.Hat
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service1_rpc_service_proto_init() }
func file_service1_rpc_service_proto_init() {
	if File_service1_rpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service1_rpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Size); i {
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
		file_service1_rpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hat); i {
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
			RawDescriptor: file_service1_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service1_rpc_service_proto_goTypes,
		DependencyIndexes: file_service1_rpc_service_proto_depIdxs,
		MessageInfos:      file_service1_rpc_service_proto_msgTypes,
	}.Build()
	File_service1_rpc_service_proto = out.File
	file_service1_rpc_service_proto_rawDesc = nil
	file_service1_rpc_service_proto_goTypes = nil
	file_service1_rpc_service_proto_depIdxs = nil
}
