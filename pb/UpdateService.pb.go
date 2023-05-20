// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: UpdateService.proto

package telegrampb

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

type ElementType int32

const (
	ElementType_topic   ElementType = 0
	ElementType_test    ElementType = 1
	ElementType_lecture ElementType = 2
)

// Enum value maps for ElementType.
var (
	ElementType_name = map[int32]string{
		0: "topic",
		1: "test",
		2: "lecture",
	}
	ElementType_value = map[string]int32{
		"topic":   0,
		"test":    1,
		"lecture": 2,
	}
)

func (x ElementType) Enum() *ElementType {
	p := new(ElementType)
	*p = x
	return p
}

func (x ElementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ElementType) Descriptor() protoreflect.EnumDescriptor {
	return file_UpdateService_proto_enumTypes[0].Descriptor()
}

func (ElementType) Type() protoreflect.EnumType {
	return &file_UpdateService_proto_enumTypes[0]
}

func (x ElementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ElementType.Descriptor instead.
func (ElementType) EnumDescriptor() ([]byte, []int) {
	return file_UpdateService_proto_rawDescGZIP(), []int{0}
}

type ActionType int32

const (
	ActionType_create ActionType = 0
	ActionType_update ActionType = 1
	ActionType_delete ActionType = 2
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "create",
		1: "update",
		2: "delete",
	}
	ActionType_value = map[string]int32{
		"create": 0,
		"update": 1,
		"delete": 2,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_UpdateService_proto_enumTypes[1].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_UpdateService_proto_enumTypes[1]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_UpdateService_proto_rawDescGZIP(), []int{1}
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course      string      `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	Element     string      `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
	ElementType ElementType `protobuf:"varint,3,opt,name=elementType,proto3,enum=kill.me.palas.grpc.ElementType" json:"elementType,omitempty"`
	ActionType  ActionType  `protobuf:"varint,4,opt,name=actionType,proto3,enum=kill.me.palas.grpc.ActionType" json:"actionType,omitempty"`
	CourseId    int64       `protobuf:"varint,5,opt,name=courseId,proto3" json:"courseId,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_UpdateService_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateRequest) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *UpdateRequest) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

func (x *UpdateRequest) GetElementType() ElementType {
	if x != nil {
		return x.ElementType
	}
	return ElementType_topic
}

func (x *UpdateRequest) GetActionType() ActionType {
	if x != nil {
		return x.ActionType
	}
	return ActionType_create
}

func (x *UpdateRequest) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_UpdateService_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_UpdateService_proto protoreflect.FileDescriptor

var file_UpdateService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x6d, 0x65, 0x2e, 0x70,
	0x61, 0x6c, 0x61, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0xe0, 0x01, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a,
	0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x6c,
	0x61, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x6d, 0x65, 0x2e, 0x70,
	0x61, 0x6c, 0x61, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x2a, 0x2f, 0x0a, 0x0b, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x74, 0x65, 0x73, 0x74, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x10, 0x02, 0x2a, 0x30, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x10, 0x02, 0x32, 0x60, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x21, 0x2e, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x6d, 0x65, 0x2e, 0x70, 0x61, 0x6c, 0x61, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x6d, 0x65, 0x2e, 0x70, 0x61,
	0x6c, 0x61, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UpdateService_proto_rawDescOnce sync.Once
	file_UpdateService_proto_rawDescData = file_UpdateService_proto_rawDesc
)

func file_UpdateService_proto_rawDescGZIP() []byte {
	file_UpdateService_proto_rawDescOnce.Do(func() {
		file_UpdateService_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdateService_proto_rawDescData)
	})
	return file_UpdateService_proto_rawDescData
}

var file_UpdateService_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_UpdateService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_UpdateService_proto_goTypes = []interface{}{
	(ElementType)(0),       // 0: kill.me.palas.grpc.ElementType
	(ActionType)(0),        // 1: kill.me.palas.grpc.ActionType
	(*UpdateRequest)(nil),  // 2: kill.me.palas.grpc.UpdateRequest
	(*UpdateResponse)(nil), // 3: kill.me.palas.grpc.UpdateResponse
}
var file_UpdateService_proto_depIdxs = []int32{
	0, // 0: kill.me.palas.grpc.UpdateRequest.elementType:type_name -> kill.me.palas.grpc.ElementType
	1, // 1: kill.me.palas.grpc.UpdateRequest.actionType:type_name -> kill.me.palas.grpc.ActionType
	2, // 2: kill.me.palas.grpc.UpdateService.update:input_type -> kill.me.palas.grpc.UpdateRequest
	3, // 3: kill.me.palas.grpc.UpdateService.update:output_type -> kill.me.palas.grpc.UpdateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_UpdateService_proto_init() }
func file_UpdateService_proto_init() {
	if File_UpdateService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UpdateService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_UpdateService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_UpdateService_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_UpdateService_proto_goTypes,
		DependencyIndexes: file_UpdateService_proto_depIdxs,
		EnumInfos:         file_UpdateService_proto_enumTypes,
		MessageInfos:      file_UpdateService_proto_msgTypes,
	}.Build()
	File_UpdateService_proto = out.File
	file_UpdateService_proto_rawDesc = nil
	file_UpdateService_proto_goTypes = nil
	file_UpdateService_proto_depIdxs = nil
}
