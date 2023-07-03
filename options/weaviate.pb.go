// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: options/weaviate.proto

package weaviate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WeaviateFileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generate bool `protobuf:"varint,1,opt,name=generate,proto3" json:"generate,omitempty"`
}

func (x *WeaviateFileOptions) Reset() {
	*x = WeaviateFileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_weaviate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeaviateFileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeaviateFileOptions) ProtoMessage() {}

func (x *WeaviateFileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_weaviate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeaviateFileOptions.ProtoReflect.Descriptor instead.
func (*WeaviateFileOptions) Descriptor() ([]byte, []int) {
	return file_options_weaviate_proto_rawDescGZIP(), []int{0}
}

func (x *WeaviateFileOptions) GetGenerate() bool {
	if x != nil {
		return x.Generate
	}
	return false
}

type WeaviateMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generate bool `protobuf:"varint,1,opt,name=generate,proto3" json:"generate,omitempty"`
}

func (x *WeaviateMessageOptions) Reset() {
	*x = WeaviateMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_weaviate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeaviateMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeaviateMessageOptions) ProtoMessage() {}

func (x *WeaviateMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_weaviate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeaviateMessageOptions.ProtoReflect.Descriptor instead.
func (*WeaviateMessageOptions) Descriptor() ([]byte, []int) {
	return file_options_weaviate_proto_rawDescGZIP(), []int{1}
}

func (x *WeaviateMessageOptions) GetGenerate() bool {
	if x != nil {
		return x.Generate
	}
	return false
}

type WeaviateFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ignore bool `protobuf:"varint,7,opt,name=ignore,proto3" json:"ignore,omitempty"`
}

func (x *WeaviateFieldOptions) Reset() {
	*x = WeaviateFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_weaviate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeaviateFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeaviateFieldOptions) ProtoMessage() {}

func (x *WeaviateFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_weaviate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeaviateFieldOptions.ProtoReflect.Descriptor instead.
func (*WeaviateFieldOptions) Descriptor() ([]byte, []int) {
	return file_options_weaviate_proto_rawDescGZIP(), []int{2}
}

func (x *WeaviateFieldOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

var file_options_weaviate_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*WeaviateFileOptions)(nil),
		Field:         52120,
		Name:          "weaviate.file_opts",
		Tag:           "bytes,52120,opt,name=file_opts",
		Filename:      "options/weaviate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*WeaviateMessageOptions)(nil),
		Field:         52120,
		Name:          "weaviate.opts",
		Tag:           "bytes,52120,opt,name=opts",
		Filename:      "options/weaviate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*WeaviateFieldOptions)(nil),
		Field:         52120,
		Name:          "weaviate.field",
		Tag:           "bytes,52120,opt,name=field",
		Filename:      "options/weaviate.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional weaviate.WeaviateFileOptions file_opts = 52120;
	E_FileOpts = &file_options_weaviate_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// ormable will cause orm code to be generated for this message/object
	//
	// optional weaviate.WeaviateMessageOptions opts = 52120;
	E_Opts = &file_options_weaviate_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional weaviate.WeaviateFieldOptions field = 52120;
	E_Field = &file_options_weaviate_proto_extTypes[2]
)

var File_options_weaviate_proto protoreflect.FileDescriptor

var file_options_weaviate_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61,
	0x74, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x13, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x22, 0x34, 0x0a, 0x16, 0x57, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x22, 0x2e, 0x0a,
	0x14, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x3a, 0x5a, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x98, 0x97, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x76,
	0x69, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x3a, 0x57, 0x0a, 0x04, 0x6f, 0x70, 0x74,
	0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x98, 0x97, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x65, 0x61,
	0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x6f, 0x70,
	0x74, 0x73, 0x3a, 0x55, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x98, 0x97, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x57, 0x65,
	0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x79, 0x73, 0x74,
	0x73, 0x71, 0x75, 0x61, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_weaviate_proto_rawDescOnce sync.Once
	file_options_weaviate_proto_rawDescData = file_options_weaviate_proto_rawDesc
)

func file_options_weaviate_proto_rawDescGZIP() []byte {
	file_options_weaviate_proto_rawDescOnce.Do(func() {
		file_options_weaviate_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_weaviate_proto_rawDescData)
	})
	return file_options_weaviate_proto_rawDescData
}

var file_options_weaviate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_options_weaviate_proto_goTypes = []interface{}{
	(*WeaviateFileOptions)(nil),         // 0: weaviate.WeaviateFileOptions
	(*WeaviateMessageOptions)(nil),      // 1: weaviate.WeaviateMessageOptions
	(*WeaviateFieldOptions)(nil),        // 2: weaviate.WeaviateFieldOptions
	(*descriptorpb.FileOptions)(nil),    // 3: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 5: google.protobuf.FieldOptions
}
var file_options_weaviate_proto_depIdxs = []int32{
	3, // 0: weaviate.file_opts:extendee -> google.protobuf.FileOptions
	4, // 1: weaviate.opts:extendee -> google.protobuf.MessageOptions
	5, // 2: weaviate.field:extendee -> google.protobuf.FieldOptions
	0, // 3: weaviate.file_opts:type_name -> weaviate.WeaviateFileOptions
	1, // 4: weaviate.opts:type_name -> weaviate.WeaviateMessageOptions
	2, // 5: weaviate.field:type_name -> weaviate.WeaviateFieldOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_weaviate_proto_init() }
func file_options_weaviate_proto_init() {
	if File_options_weaviate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_weaviate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeaviateFileOptions); i {
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
		file_options_weaviate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeaviateMessageOptions); i {
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
		file_options_weaviate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeaviateFieldOptions); i {
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
			RawDescriptor: file_options_weaviate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_options_weaviate_proto_goTypes,
		DependencyIndexes: file_options_weaviate_proto_depIdxs,
		MessageInfos:      file_options_weaviate_proto_msgTypes,
		ExtensionInfos:    file_options_weaviate_proto_extTypes,
	}.Build()
	File_options_weaviate_proto = out.File
	file_options_weaviate_proto_rawDesc = nil
	file_options_weaviate_proto_goTypes = nil
	file_options_weaviate_proto_depIdxs = nil
}
