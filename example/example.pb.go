// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: example.proto

package example_example

import (
	_ "github.com/catalystsquad/protoc-gen-go-weaviate/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Thing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: fake:"{uuid}"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" fake:"{uuid}"`
	// @gotags: fake:"{price:0.00,1000.00}"
	ADouble float64 `protobuf:"fixed64,2,opt,name=a_double,json=aDouble,proto3" json:"a_double,omitempty" fake:"{price:0.00,1000.00}"`
	// @gotags: fake:"{price:0.00,1000.00}"
	AFloat float32 `protobuf:"fixed32,3,opt,name=a_float,json=aFloat,proto3" json:"a_float,omitempty" fake:"{price:0.00,1000.00}"`
	// @gotags: fake:"{int32}"
	AnInt32 int32 `protobuf:"varint,4,opt,name=an_int32,json=anInt32,proto3" json:"an_int32,omitempty" fake:"{int32}"`
	// @gotags: fake:"{int64}"
	AnInt64 int64 `protobuf:"varint,5,opt,name=an_int64,json=anInt64,proto3" json:"an_int64,omitempty" fake:"{int64}"`
	// @gotags: fake:"{bool}"
	ABool bool `protobuf:"varint,14,opt,name=a_bool,json=aBool,proto3" json:"a_bool,omitempty" fake:"{bool}"`
	// @gotags: fake:"{hackerphrase}"
	AString string `protobuf:"bytes,15,opt,name=a_string,json=aString,proto3" json:"a_string,omitempty" fake:"{hackerphrase}"`
	// @gotags: fake:"skip"
	ABytes []byte `protobuf:"bytes,16,opt,name=a_bytes,json=aBytes,proto3" json:"a_bytes,omitempty" fake:"skip"`
	// @gotags: fake:"{hackerphrase}"
	RepeatedScalarField []string `protobuf:"bytes,17,rep,name=repeated_scalar_field,json=repeatedScalarField,proto3" json:"repeated_scalar_field,omitempty" fake:"{hackerphrase}"`
	// @gotags: fake:"skip"
	OptionalScalarField *string `protobuf:"bytes,18,opt,name=optional_scalar_field,json=optionalScalarField,proto3,oneof" json:"optional_scalar_field,omitempty" fake:"skip"`
	// @gotags: fake:"skip"
	AssociatedThing *Thing2 `protobuf:"bytes,19,opt,name=associated_thing,json=associatedThing,proto3" json:"associated_thing,omitempty" fake:"skip"`
	// @gotags: fake:"skip"
	OptionalAssociatedThing *Thing2 `protobuf:"bytes,20,opt,name=optional_associated_thing,json=optionalAssociatedThing,proto3,oneof" json:"optional_associated_thing,omitempty" fake:"skip"`
	// @gotags: fake:"skip"
	RepeatedMessages []*Thing2 `protobuf:"bytes,21,rep,name=repeated_messages,json=repeatedMessages,proto3" json:"repeated_messages,omitempty" fake:"skip"`
	// @gotags: fake:"skip"
	ATimestamp *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=a_timestamp,json=aTimestamp,proto3" json:"a_timestamp,omitempty" fake:"skip"`
	// @gotags: fake:"skip"
	AnIgnoredField string `protobuf:"bytes,23,opt,name=an_ignored_field,json=anIgnoredField,proto3" json:"an_ignored_field,omitempty" fake:"skip"`
}

func (x *Thing) Reset() {
	*x = Thing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thing) ProtoMessage() {}

func (x *Thing) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thing.ProtoReflect.Descriptor instead.
func (*Thing) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *Thing) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Thing) GetADouble() float64 {
	if x != nil {
		return x.ADouble
	}
	return 0
}

func (x *Thing) GetAFloat() float32 {
	if x != nil {
		return x.AFloat
	}
	return 0
}

func (x *Thing) GetAnInt32() int32 {
	if x != nil {
		return x.AnInt32
	}
	return 0
}

func (x *Thing) GetAnInt64() int64 {
	if x != nil {
		return x.AnInt64
	}
	return 0
}

func (x *Thing) GetABool() bool {
	if x != nil {
		return x.ABool
	}
	return false
}

func (x *Thing) GetAString() string {
	if x != nil {
		return x.AString
	}
	return ""
}

func (x *Thing) GetABytes() []byte {
	if x != nil {
		return x.ABytes
	}
	return nil
}

func (x *Thing) GetRepeatedScalarField() []string {
	if x != nil {
		return x.RepeatedScalarField
	}
	return nil
}

func (x *Thing) GetOptionalScalarField() string {
	if x != nil && x.OptionalScalarField != nil {
		return *x.OptionalScalarField
	}
	return ""
}

func (x *Thing) GetAssociatedThing() *Thing2 {
	if x != nil {
		return x.AssociatedThing
	}
	return nil
}

func (x *Thing) GetOptionalAssociatedThing() *Thing2 {
	if x != nil {
		return x.OptionalAssociatedThing
	}
	return nil
}

func (x *Thing) GetRepeatedMessages() []*Thing2 {
	if x != nil {
		return x.RepeatedMessages
	}
	return nil
}

func (x *Thing) GetATimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ATimestamp
	}
	return nil
}

func (x *Thing) GetAnIgnoredField() string {
	if x != nil {
		return x.AnIgnoredField
	}
	return ""
}

type Thing2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: fake:"{uuid}"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" fake:"{uuid}"`
	// @gotags: fake:"{name}"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" fake:"{name}"`
}

func (x *Thing2) Reset() {
	*x = Thing2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thing2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thing2) ProtoMessage() {}

func (x *Thing2) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thing2.ProtoReflect.Descriptor instead.
func (*Thing2) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *Thing2) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Thing2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77,
	0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab,
	0x05, 0x0a, 0x05, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x61, 0x6e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6e, 0x5f, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x11, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x37, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x13, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x10, 0x61, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x32, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x12, 0x4d, 0x0a, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x32, 0x48, 0x01, 0x52, 0x17, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x88,
	0x01, 0x01, 0x12, 0x39, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x32, 0x52, 0x10, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3b, 0x0a,
	0x0b, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x30, 0x0a, 0x10, 0x61, 0x6e,
	0x5f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x38, 0x01, 0x52, 0x0e, 0x61, 0x6e,
	0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x06, 0xba, 0xb9,
	0x19, 0x02, 0x08, 0x01, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x1c,
	0x0a, 0x1a, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x34, 0x0a, 0x06,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02,
	0x08, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_proto_goTypes = []interface{}{
	(*Thing)(nil),                 // 0: test.Thing
	(*Thing2)(nil),                // 1: test.Thing2
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_example_proto_depIdxs = []int32{
	1, // 0: test.Thing.associated_thing:type_name -> test.Thing2
	1, // 1: test.Thing.optional_associated_thing:type_name -> test.Thing2
	1, // 2: test.Thing.repeated_messages:type_name -> test.Thing2
	2, // 3: test.Thing.a_timestamp:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thing); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thing2); i {
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
	file_example_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}