// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: console.proto

package console

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountSetupIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *AccountSetupIn) Reset() {
	*x = AccountSetupIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountSetupIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountSetupIn) ProtoMessage() {}

func (x *AccountSetupIn) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountSetupIn.ProtoReflect.Descriptor instead.
func (*AccountSetupIn) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{0}
}

func (x *AccountSetupIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type AccountSetupVoid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountSetupVoid) Reset() {
	*x = AccountSetupVoid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountSetupVoid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountSetupVoid) ProtoMessage() {}

func (x *AccountSetupVoid) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountSetupVoid.ProtoReflect.Descriptor instead.
func (*AccountSetupVoid) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{1}
}

type AppIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *AppIn) Reset() {
	*x = AppIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppIn) ProtoMessage() {}

func (x *AppIn) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppIn.ProtoReflect.Descriptor instead.
func (*AppIn) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{2}
}

func (x *AppIn) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type AppOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AppOut) Reset() {
	*x = AppOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppOut) ProtoMessage() {}

func (x *AppOut) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppOut.ProtoReflect.Descriptor instead.
func (*AppOut) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{3}
}

func (x *AppOut) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type MSvcIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsvcId string `protobuf:"bytes,1,opt,name=msvc_id,json=msvcId,proto3" json:"msvc_id,omitempty"`
}

func (x *MSvcIn) Reset() {
	*x = MSvcIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSvcIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSvcIn) ProtoMessage() {}

func (x *MSvcIn) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSvcIn.ProtoReflect.Descriptor instead.
func (*MSvcIn) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{4}
}

func (x *MSvcIn) GetMsvcId() string {
	if x != nil {
		return x.MsvcId
	}
	return ""
}

type MSvcOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MSvcOut) Reset() {
	*x = MSvcOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSvcOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSvcOut) ProtoMessage() {}

func (x *MSvcOut) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSvcOut.ProtoReflect.Descriptor instead.
func (*MSvcOut) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{5}
}

func (x *MSvcOut) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProjectIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=projectId,proto3" json:"projectId,omitempty"`
}

func (x *ProjectIn) Reset() {
	*x = ProjectIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectIn) ProtoMessage() {}

func (x *ProjectIn) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectIn.ProtoReflect.Descriptor instead.
func (*ProjectIn) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectIn) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ProjectOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProjectOut) Reset() {
	*x = ProjectOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectOut) ProtoMessage() {}

func (x *ProjectOut) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectOut.ProtoReflect.Descriptor instead.
func (*ProjectOut) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectOut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetupClusterVoid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetupClusterVoid) Reset() {
	*x = SetupClusterVoid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupClusterVoid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupClusterVoid) ProtoMessage() {}

func (x *SetupClusterVoid) ProtoReflect() protoreflect.Message {
	mi := &file_console_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupClusterVoid.ProtoReflect.Descriptor instead.
func (*SetupClusterVoid) Descriptor() ([]byte, []int) {
	return file_console_proto_rawDescGZIP(), []int{8}
}

var File_console_proto protoreflect.FileDescriptor

var file_console_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x75, 0x70, 0x49, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x75, 0x70, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x1e,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x32,
	0x0a, 0x06, 0x41, 0x70, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x21, 0x0a, 0x06, 0x4d, 0x53, 0x76, 0x63, 0x49, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x73, 0x76, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x73, 0x76, 0x63, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x07, 0x4d, 0x53, 0x76, 0x63, 0x4f, 0x75, 0x74,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x56, 0x6f, 0x69, 0x64, 0x32, 0xa7, 0x01, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0a, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4f,
	0x75, 0x74, 0x12, 0x19, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x12, 0x06, 0x2e, 0x41,
	0x70, 0x70, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x70, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x22, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x53, 0x76, 0x63, 0x12, 0x07,
	0x2e, 0x4d, 0x53, 0x76, 0x63, 0x49, 0x6e, 0x1a, 0x08, 0x2e, 0x4d, 0x53, 0x76, 0x63, 0x4f, 0x75,
	0x74, 0x12, 0x32, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x6b, 0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69,
	0x74, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_console_proto_rawDescOnce sync.Once
	file_console_proto_rawDescData = file_console_proto_rawDesc
)

func file_console_proto_rawDescGZIP() []byte {
	file_console_proto_rawDescOnce.Do(func() {
		file_console_proto_rawDescData = protoimpl.X.CompressGZIP(file_console_proto_rawDescData)
	})
	return file_console_proto_rawDescData
}

var file_console_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_console_proto_goTypes = []interface{}{
	(*AccountSetupIn)(nil),   // 0: AccountSetupIn
	(*AccountSetupVoid)(nil), // 1: AccountSetupVoid
	(*AppIn)(nil),            // 2: AppIn
	(*AppOut)(nil),           // 3: AppOut
	(*MSvcIn)(nil),           // 4: MSvcIn
	(*MSvcOut)(nil),          // 5: MSvcOut
	(*ProjectIn)(nil),        // 6: ProjectIn
	(*ProjectOut)(nil),       // 7: ProjectOut
	(*SetupClusterVoid)(nil), // 8: SetupClusterVoid
	(*anypb.Any)(nil),        // 9: google.protobuf.Any
}
var file_console_proto_depIdxs = []int32{
	9, // 0: AppOut.data:type_name -> google.protobuf.Any
	9, // 1: MSvcOut.data:type_name -> google.protobuf.Any
	6, // 2: Console.GetProjectName:input_type -> ProjectIn
	2, // 3: Console.GetApp:input_type -> AppIn
	4, // 4: Console.GetManagedSvc:input_type -> MSvcIn
	0, // 5: Console.SetupAccount:input_type -> AccountSetupIn
	7, // 6: Console.GetProjectName:output_type -> ProjectOut
	3, // 7: Console.GetApp:output_type -> AppOut
	5, // 8: Console.GetManagedSvc:output_type -> MSvcOut
	1, // 9: Console.SetupAccount:output_type -> AccountSetupVoid
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_console_proto_init() }
func file_console_proto_init() {
	if File_console_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_console_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountSetupIn); i {
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
		file_console_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountSetupVoid); i {
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
		file_console_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppIn); i {
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
		file_console_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppOut); i {
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
		file_console_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSvcIn); i {
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
		file_console_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSvcOut); i {
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
		file_console_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectIn); i {
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
		file_console_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectOut); i {
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
		file_console_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupClusterVoid); i {
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
			RawDescriptor: file_console_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_console_proto_goTypes,
		DependencyIndexes: file_console_proto_depIdxs,
		MessageInfos:      file_console_proto_msgTypes,
	}.Build()
	File_console_proto = out.File
	file_console_proto_rawDesc = nil
	file_console_proto_goTypes = nil
	file_console_proto_depIdxs = nil
}
