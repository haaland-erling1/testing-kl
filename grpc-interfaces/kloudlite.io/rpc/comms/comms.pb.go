// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: comms.proto

package comms

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

type VerificationEmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email             string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VerificationToken string `protobuf:"bytes,3,opt,name=verificationToken,proto3" json:"verificationToken,omitempty"`
}

func (x *VerificationEmailInput) Reset() {
	*x = VerificationEmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationEmailInput) ProtoMessage() {}

func (x *VerificationEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_comms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationEmailInput.ProtoReflect.Descriptor instead.
func (*VerificationEmailInput) Descriptor() ([]byte, []int) {
	return file_comms_proto_rawDescGZIP(), []int{0}
}

func (x *VerificationEmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerificationEmailInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VerificationEmailInput) GetVerificationToken() string {
	if x != nil {
		return x.VerificationToken
	}
	return ""
}

type WelcomeEmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WelcomeEmailInput) Reset() {
	*x = WelcomeEmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomeEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeEmailInput) ProtoMessage() {}

func (x *WelcomeEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_comms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeEmailInput.ProtoReflect.Descriptor instead.
func (*WelcomeEmailInput) Descriptor() ([]byte, []int) {
	return file_comms_proto_rawDescGZIP(), []int{1}
}

func (x *WelcomeEmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *WelcomeEmailInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PasswordResetEmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ResetToken string `protobuf:"bytes,3,opt,name=resetToken,proto3" json:"resetToken,omitempty"`
}

func (x *PasswordResetEmailInput) Reset() {
	*x = PasswordResetEmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordResetEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResetEmailInput) ProtoMessage() {}

func (x *PasswordResetEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_comms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResetEmailInput.ProtoReflect.Descriptor instead.
func (*PasswordResetEmailInput) Descriptor() ([]byte, []int) {
	return file_comms_proto_rawDescGZIP(), []int{2}
}

func (x *PasswordResetEmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PasswordResetEmailInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PasswordResetEmailInput) GetResetToken() string {
	if x != nil {
		return x.ResetToken
	}
	return ""
}

type AccountMemberInviteEmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName     string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	InvitationToken string `protobuf:"bytes,2,opt,name=invitationToken,proto3" json:"invitationToken,omitempty"`
	InvitedBy       string `protobuf:"bytes,3,opt,name=invitedBy,proto3" json:"invitedBy,omitempty"`
	Email           string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Name            string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccountMemberInviteEmailInput) Reset() {
	*x = AccountMemberInviteEmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountMemberInviteEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountMemberInviteEmailInput) ProtoMessage() {}

func (x *AccountMemberInviteEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_comms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountMemberInviteEmailInput.ProtoReflect.Descriptor instead.
func (*AccountMemberInviteEmailInput) Descriptor() ([]byte, []int) {
	return file_comms_proto_rawDescGZIP(), []int{3}
}

func (x *AccountMemberInviteEmailInput) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *AccountMemberInviteEmailInput) GetInvitationToken() string {
	if x != nil {
		return x.InvitationToken
	}
	return ""
}

func (x *AccountMemberInviteEmailInput) GetInvitedBy() string {
	if x != nil {
		return x.InvitedBy
	}
	return ""
}

func (x *AccountMemberInviteEmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountMemberInviteEmailInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProjectMemberInviteEmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName     string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	InvitationToken string `protobuf:"bytes,2,opt,name=invitationToken,proto3" json:"invitationToken,omitempty"`
	InvitedBy       string `protobuf:"bytes,3,opt,name=invitedBy,proto3" json:"invitedBy,omitempty"`
	Email           string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Name            string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProjectMemberInviteEmailInput) Reset() {
	*x = ProjectMemberInviteEmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectMemberInviteEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMemberInviteEmailInput) ProtoMessage() {}

func (x *ProjectMemberInviteEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_comms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectMemberInviteEmailInput.ProtoReflect.Descriptor instead.
func (*ProjectMemberInviteEmailInput) Descriptor() ([]byte, []int) {
	return file_comms_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectMemberInviteEmailInput) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ProjectMemberInviteEmailInput) GetInvitationToken() string {
	if x != nil {
		return x.InvitationToken
	}
	return ""
}

func (x *ProjectMemberInviteEmailInput) GetInvitedBy() string {
	if x != nil {
		return x.InvitedBy
	}
	return ""
}

func (x *ProjectMemberInviteEmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ProjectMemberInviteEmailInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SendContactUsEmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email        string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CompanyName  string `protobuf:"bytes,3,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Country      string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	MobileNumber int64  `protobuf:"varint,5,opt,name=mobileNumber,proto3" json:"mobileNumber,omitempty"`
	Message      string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendContactUsEmailInput) Reset() {
	*x = SendContactUsEmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendContactUsEmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendContactUsEmailInput) ProtoMessage() {}

func (x *SendContactUsEmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_comms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendContactUsEmailInput.ProtoReflect.Descriptor instead.
func (*SendContactUsEmailInput) Descriptor() ([]byte, []int) {
	return file_comms_proto_rawDescGZIP(), []int{5}
}

func (x *SendContactUsEmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendContactUsEmailInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SendContactUsEmailInput) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *SendContactUsEmailInput) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *SendContactUsEmailInput) GetMobileNumber() int64 {
	if x != nil {
		return x.MobileNumber
	}
	return 0
}

func (x *SendContactUsEmailInput) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_comms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_comms_proto_rawDescGZIP(), []int{6}
}

var File_comms_proto protoreflect.FileDescriptor

var file_comms_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a,
	0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x3d, 0x0a, 0x11, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63,
	0x0a, 0x17, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x1d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x1d, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xbd, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x55,
	0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x9e, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x6d,
	0x73, 0x12, 0x37, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x16, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x05,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x1c,
	0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x05, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x6c, 0x63, 0x6f,
	0x6d, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x2d, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x12, 0x35, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x55, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x55, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x18, 0x5a, 0x16, 0x6b, 0x6c, 0x6f, 0x75,
	0x64, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comms_proto_rawDescOnce sync.Once
	file_comms_proto_rawDescData = file_comms_proto_rawDesc
)

func file_comms_proto_rawDescGZIP() []byte {
	file_comms_proto_rawDescOnce.Do(func() {
		file_comms_proto_rawDescData = protoimpl.X.CompressGZIP(file_comms_proto_rawDescData)
	})
	return file_comms_proto_rawDescData
}

var file_comms_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_comms_proto_goTypes = []interface{}{
	(*VerificationEmailInput)(nil),        // 0: VerificationEmailInput
	(*WelcomeEmailInput)(nil),             // 1: WelcomeEmailInput
	(*PasswordResetEmailInput)(nil),       // 2: PasswordResetEmailInput
	(*AccountMemberInviteEmailInput)(nil), // 3: AccountMemberInviteEmailInput
	(*ProjectMemberInviteEmailInput)(nil), // 4: ProjectMemberInviteEmailInput
	(*SendContactUsEmailInput)(nil),       // 5: SendContactUsEmailInput
	(*Void)(nil),                          // 6: Void
}
var file_comms_proto_depIdxs = []int32{
	0, // 0: Comms.SendVerificationEmail:input_type -> VerificationEmailInput
	2, // 1: Comms.SendPasswordResetEmail:input_type -> PasswordResetEmailInput
	3, // 2: Comms.SendAccountMemberInviteEmail:input_type -> AccountMemberInviteEmailInput
	4, // 3: Comms.SendProjectMemberInviteEmail:input_type -> ProjectMemberInviteEmailInput
	1, // 4: Comms.SendWelcomeEmail:input_type -> WelcomeEmailInput
	1, // 5: Comms.SendWaitingEmail:input_type -> WelcomeEmailInput
	5, // 6: Comms.SendContactUsEmail:input_type -> SendContactUsEmailInput
	6, // 7: Comms.SendVerificationEmail:output_type -> Void
	6, // 8: Comms.SendPasswordResetEmail:output_type -> Void
	6, // 9: Comms.SendAccountMemberInviteEmail:output_type -> Void
	6, // 10: Comms.SendProjectMemberInviteEmail:output_type -> Void
	6, // 11: Comms.SendWelcomeEmail:output_type -> Void
	6, // 12: Comms.SendWaitingEmail:output_type -> Void
	6, // 13: Comms.SendContactUsEmail:output_type -> Void
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comms_proto_init() }
func file_comms_proto_init() {
	if File_comms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationEmailInput); i {
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
		file_comms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomeEmailInput); i {
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
		file_comms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordResetEmailInput); i {
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
		file_comms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountMemberInviteEmailInput); i {
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
		file_comms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectMemberInviteEmailInput); i {
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
		file_comms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendContactUsEmailInput); i {
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
		file_comms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_comms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comms_proto_goTypes,
		DependencyIndexes: file_comms_proto_depIdxs,
		MessageInfos:      file_comms_proto_msgTypes,
	}.Build()
	File_comms_proto = out.File
	file_comms_proto_rawDesc = nil
	file_comms_proto_goTypes = nil
	file_comms_proto_depIdxs = nil
}
