// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.1
// source: packet.proto

package packet

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

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ByCode []byte `protobuf:"bytes,1,opt,name=byCode,proto3" json:"byCode,omitempty"` //1byte. packet code 0x89
	Size   []byte `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`     //2byte. packet size
	ByType []byte `protobuf:"bytes,3,opt,name=byType,proto3" json:"byType,omitempty"` //2byte. packet type
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetByCode() []byte {
	if x != nil {
		return x.ByCode
	}
	return nil
}

func (x *Header) GetSize() []byte {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *Header) GetByType() []byte {
	if x != nil {
		return x.ByType
	}
	return nil
}

type SCConnectSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SCConnectSuccess) Reset() {
	*x = SCConnectSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCConnectSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCConnectSuccess) ProtoMessage() {}

func (x *SCConnectSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCConnectSuccess.ProtoReflect.Descriptor instead.
func (*SCConnectSuccess) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{1}
}

type CSConnectSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CSConnectSuccess) Reset() {
	*x = CSConnectSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSConnectSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSConnectSuccess) ProtoMessage() {}

func (x *CSConnectSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSConnectSuccess.ProtoReflect.Descriptor instead.
func (*CSConnectSuccess) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{2}
}

var File_packet_proto protoreflect.FileDescriptor

var file_packet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x75, 0x72, 0x70, 0x67, 0x22, 0x4c, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x62, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x43, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packet_proto_rawDescOnce sync.Once
	file_packet_proto_rawDescData = file_packet_proto_rawDesc
)

func file_packet_proto_rawDescGZIP() []byte {
	file_packet_proto_rawDescOnce.Do(func() {
		file_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_packet_proto_rawDescData)
	})
	return file_packet_proto_rawDescData
}

var file_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_packet_proto_goTypes = []interface{}{
	(*Header)(nil),           // 0: urpg.Header
	(*SCConnectSuccess)(nil), // 1: urpg.SCConnectSuccess
	(*CSConnectSuccess)(nil), // 2: urpg.CSConnectSuccess
}
var file_packet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_packet_proto_init() }
func file_packet_proto_init() {
	if File_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_packet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCConnectSuccess); i {
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
		file_packet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSConnectSuccess); i {
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
			RawDescriptor: file_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packet_proto_goTypes,
		DependencyIndexes: file_packet_proto_depIdxs,
		MessageInfos:      file_packet_proto_msgTypes,
	}.Build()
	File_packet_proto = out.File
	file_packet_proto_rawDesc = nil
	file_packet_proto_goTypes = nil
	file_packet_proto_depIdxs = nil
}