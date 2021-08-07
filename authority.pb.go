// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: authority.proto

package types

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

type Authority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Address:
	//	*Authority_V4
	//	*Authority_V6
	Address isAuthority_Address `protobuf_oneof:"address"`
	Port    uint32              `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Authority) Reset() {
	*x = Authority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authority) ProtoMessage() {}

func (x *Authority) ProtoReflect() protoreflect.Message {
	mi := &file_authority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authority.ProtoReflect.Descriptor instead.
func (*Authority) Descriptor() ([]byte, []int) {
	return file_authority_proto_rawDescGZIP(), []int{0}
}

func (m *Authority) GetAddress() isAuthority_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (x *Authority) GetV4() uint32 {
	if x, ok := x.GetAddress().(*Authority_V4); ok {
		return x.V4
	}
	return 0
}

func (x *Authority) GetV6() []byte {
	if x, ok := x.GetAddress().(*Authority_V6); ok {
		return x.V6
	}
	return nil
}

func (x *Authority) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type isAuthority_Address interface {
	isAuthority_Address()
}

type Authority_V4 struct {
	V4 uint32 `protobuf:"fixed32,1,opt,name=v4,proto3,oneof"`
}

type Authority_V6 struct {
	V6 []byte `protobuf:"bytes,2,opt,name=v6,proto3,oneof"`
}

func (*Authority_V4) isAuthority_Address() {}

func (*Authority_V6) isAuthority_Address() {}

var File_authority_proto protoreflect.FileDescriptor

var file_authority_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x62, 0x6c, 0x61, 0x68, 0x63, 0x64, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x4e, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a,
	0x02, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x48, 0x00, 0x52, 0x02, 0x76, 0x34, 0x12,
	0x10, 0x0a, 0x02, 0x76, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x02, 0x76,
	0x36, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x6c, 0x61, 0x68, 0x63, 0x64, 0x6e, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authority_proto_rawDescOnce sync.Once
	file_authority_proto_rawDescData = file_authority_proto_rawDesc
)

func file_authority_proto_rawDescGZIP() []byte {
	file_authority_proto_rawDescOnce.Do(func() {
		file_authority_proto_rawDescData = protoimpl.X.CompressGZIP(file_authority_proto_rawDescData)
	})
	return file_authority_proto_rawDescData
}

var file_authority_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_authority_proto_goTypes = []interface{}{
	(*Authority)(nil), // 0: blahcdn.types.Authority
}
var file_authority_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authority_proto_init() }
func file_authority_proto_init() {
	if File_authority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authority); i {
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
	file_authority_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Authority_V4)(nil),
		(*Authority_V6)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authority_proto_goTypes,
		DependencyIndexes: file_authority_proto_depIdxs,
		MessageInfos:      file_authority_proto_msgTypes,
	}.Build()
	File_authority_proto = out.File
	file_authority_proto_rawDesc = nil
	file_authority_proto_goTypes = nil
	file_authority_proto_depIdxs = nil
}
