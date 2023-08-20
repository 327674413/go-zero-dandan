// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.19.4
// source: user.proto

package pb

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

type UserMainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UnionId   int64  `protobuf:"varint,2,opt,name=UnionId,proto3" json:"UnionId,omitempty"`
	StateEm   int64  `protobuf:"varint,3,opt,name=StateEm,proto3" json:"StateEm,omitempty"`
	Account   string `protobuf:"bytes,4,opt,name=Account,proto3" json:"Account,omitempty"`
	Nickname  string `protobuf:"bytes,5,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=Phone,proto3" json:"Phone,omitempty"`
	PhoneArea string `protobuf:"bytes,7,opt,name=PhoneArea,proto3" json:"PhoneArea,omitempty"`
	SexEm     int64  `protobuf:"varint,8,opt,name=SexEm,proto3" json:"SexEm,omitempty"`
	Email     string `protobuf:"bytes,9,opt,name=Email,proto3" json:"Email,omitempty"`
	Avatar    string `protobuf:"bytes,10,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
}

func (x *UserMainInfo) Reset() {
	*x = UserMainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMainInfo) ProtoMessage() {}

func (x *UserMainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMainInfo.ProtoReflect.Descriptor instead.
func (*UserMainInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserMainInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserMainInfo) GetUnionId() int64 {
	if x != nil {
		return x.UnionId
	}
	return 0
}

func (x *UserMainInfo) GetStateEm() int64 {
	if x != nil {
		return x.StateEm
	}
	return 0
}

func (x *UserMainInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserMainInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserMainInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserMainInfo) GetPhoneArea() string {
	if x != nil {
		return x.PhoneArea
	}
	return ""
}

func (x *UserMainInfo) GetSexEm() int64 {
	if x != nil {
		return x.SexEm
	}
	return 0
}

func (x *UserMainInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserMainInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type EditUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Nickname     *string `protobuf:"bytes,2,opt,name=Nickname,proto3,oneof" json:"Nickname,omitempty"`
	SexEm        *int64  `protobuf:"varint,3,opt,name=SexEm,proto3,oneof" json:"SexEm,omitempty"`
	Email        *string `protobuf:"bytes,4,opt,name=Email,proto3,oneof" json:"Email,omitempty"`
	Avatar       *string `protobuf:"bytes,5,opt,name=Avatar,proto3,oneof" json:"Avatar,omitempty"`
	GraduateFrom *string `protobuf:"bytes,6,opt,name=GraduateFrom,proto3,oneof" json:"GraduateFrom,omitempty"`
	BirthDate    *string `protobuf:"bytes,7,opt,name=BirthDate,proto3,oneof" json:"BirthDate,omitempty"`
}

func (x *EditUserInfoReq) Reset() {
	*x = EditUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserInfoReq) ProtoMessage() {}

func (x *EditUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserInfoReq.ProtoReflect.Descriptor instead.
func (*EditUserInfoReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *EditUserInfoReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditUserInfoReq) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *EditUserInfoReq) GetSexEm() int64 {
	if x != nil && x.SexEm != nil {
		return *x.SexEm
	}
	return 0
}

func (x *EditUserInfoReq) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *EditUserInfoReq) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

func (x *EditUserInfoReq) GetGraduateFrom() string {
	if x != nil && x.GraduateFrom != nil {
		return *x.GraduateFrom
	}
	return ""
}

func (x *EditUserInfoReq) GetBirthDate() string {
	if x != nil && x.BirthDate != nil {
		return *x.BirthDate
	}
	return ""
}

type UserCronyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserId      int64  `protobuf:"varint,1,opt,name=OwnerUserId,proto3" json:"OwnerUserId,omitempty"`
	TargetUserId     int64  `protobuf:"varint,2,opt,name=TargetUserId,proto3" json:"TargetUserId,omitempty"`
	Remark           string `protobuf:"bytes,3,opt,name=Remark,proto3" json:"Remark,omitempty"`
	TypeEm           int64  `protobuf:"varint,4,opt,name=TypeEm,proto3" json:"TypeEm,omitempty"`
	NameNote         string `protobuf:"bytes,5,opt,name=NameNote,proto3" json:"NameNote,omitempty"`
	CreateAt         int64  `protobuf:"varint,6,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	TargetUserName   string `protobuf:"bytes,7,opt,name=TargetUserName,proto3" json:"TargetUserName,omitempty"`
	TargetUserAvatar string `protobuf:"bytes,8,opt,name=TargetUserAvatar,proto3" json:"TargetUserAvatar,omitempty"`
	TagIds           string `protobuf:"bytes,9,opt,name=TagIds,proto3" json:"TagIds,omitempty"`
	GroupId          int64  `protobuf:"varint,10,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	GroupName        string `protobuf:"bytes,11,opt,name=GroupName,proto3" json:"GroupName,omitempty"`
	Id               int64  `protobuf:"varint,12,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *UserCronyInfo) Reset() {
	*x = UserCronyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCronyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCronyInfo) ProtoMessage() {}

func (x *UserCronyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCronyInfo.ProtoReflect.Descriptor instead.
func (*UserCronyInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserCronyInfo) GetOwnerUserId() int64 {
	if x != nil {
		return x.OwnerUserId
	}
	return 0
}

func (x *UserCronyInfo) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

func (x *UserCronyInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *UserCronyInfo) GetTypeEm() int64 {
	if x != nil {
		return x.TypeEm
	}
	return 0
}

func (x *UserCronyInfo) GetNameNote() string {
	if x != nil {
		return x.NameNote
	}
	return ""
}

func (x *UserCronyInfo) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *UserCronyInfo) GetTargetUserName() string {
	if x != nil {
		return x.TargetUserName
	}
	return ""
}

func (x *UserCronyInfo) GetTargetUserAvatar() string {
	if x != nil {
		return x.TargetUserAvatar
	}
	return ""
}

func (x *UserCronyInfo) GetTagIds() string {
	if x != nil {
		return x.TagIds
	}
	return ""
}

func (x *UserCronyInfo) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UserCronyInfo) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *UserCronyInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserCronyListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*UserCronyInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total *int64           `protobuf:"varint,2,opt,name=total,proto3,oneof" json:"total,omitempty"`
}

func (x *GetUserCronyListResp) Reset() {
	*x = GetUserCronyListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserCronyListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserCronyListResp) ProtoMessage() {}

func (x *GetUserCronyListResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserCronyListResp.ProtoReflect.Descriptor instead.
func (*GetUserCronyListResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserCronyListResp) GetList() []*UserCronyInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *GetUserCronyListResp) GetTotal() int64 {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return 0
}

type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	PlatId int64 `protobuf:"varint,2,opt,name=PlatId,proto3" json:"PlatId,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *IdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IdReq) GetPlatId() int64 {
	if x != nil {
		return x.PlatId
	}
	return 0
}

type TokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenReq) Reset() {
	*x = TokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq) ProtoMessage() {}

func (x *TokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq.ProtoReflect.Descriptor instead.
func (*TokenReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *TokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SuccResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *SuccResp) Reset() {
	*x = SuccResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccResp) ProtoMessage() {}

func (x *SuccResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccResp.ProtoReflect.Descriptor instead.
func (*SuccResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *SuccResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetUserCronyListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatId        *int64  `protobuf:"varint,1,opt,name=PlatId,proto3,oneof" json:"PlatId,omitempty"`
	IsNeedTotal   *int64  `protobuf:"varint,2,opt,name=IsNeedTotal,proto3,oneof" json:"IsNeedTotal,omitempty"`
	OwnerUserId   *int64  `protobuf:"varint,3,opt,name=OwnerUserId,proto3,oneof" json:"OwnerUserId,omitempty"`
	OwnerUserName *string `protobuf:"bytes,4,opt,name=OwnerUserName,proto3,oneof" json:"OwnerUserName,omitempty"`
	GroupId       *int64  `protobuf:"varint,5,opt,name=GroupId,proto3,oneof" json:"GroupId,omitempty"`
	TypeEms       *string `protobuf:"bytes,6,opt,name=TypeEms,proto3,oneof" json:"TypeEms,omitempty"`
	AddStartTime  *string `protobuf:"bytes,7,opt,name=AddStartTime,proto3,oneof" json:"AddStartTime,omitempty"`
	AddEndTime    *string `protobuf:"bytes,8,opt,name=AddEndTime,proto3,oneof" json:"AddEndTime,omitempty"`
}

func (x *GetUserCronyListReq) Reset() {
	*x = GetUserCronyListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserCronyListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserCronyListReq) ProtoMessage() {}

func (x *GetUserCronyListReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserCronyListReq.ProtoReflect.Descriptor instead.
func (*GetUserCronyListReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserCronyListReq) GetPlatId() int64 {
	if x != nil && x.PlatId != nil {
		return *x.PlatId
	}
	return 0
}

func (x *GetUserCronyListReq) GetIsNeedTotal() int64 {
	if x != nil && x.IsNeedTotal != nil {
		return *x.IsNeedTotal
	}
	return 0
}

func (x *GetUserCronyListReq) GetOwnerUserId() int64 {
	if x != nil && x.OwnerUserId != nil {
		return *x.OwnerUserId
	}
	return 0
}

func (x *GetUserCronyListReq) GetOwnerUserName() string {
	if x != nil && x.OwnerUserName != nil {
		return *x.OwnerUserName
	}
	return ""
}

func (x *GetUserCronyListReq) GetGroupId() int64 {
	if x != nil && x.GroupId != nil {
		return *x.GroupId
	}
	return 0
}

func (x *GetUserCronyListReq) GetTypeEms() string {
	if x != nil && x.TypeEms != nil {
		return *x.TypeEms
	}
	return ""
}

func (x *GetUserCronyListReq) GetAddStartTime() string {
	if x != nil && x.AddStartTime != nil {
		return *x.AddStartTime
	}
	return ""
}

func (x *GetUserCronyListReq) GetAddEndTime() string {
	if x != nil && x.AddEndTime != nil {
		return *x.AddEndTime
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x80, 0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x41, 0x72, 0x65,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x78, 0x45, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x53, 0x65, 0x78, 0x45, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0xac, 0x02, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x53, 0x65,
	0x78, 0x45, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x53, 0x65, 0x78,
	0x45, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a,
	0x0c, 0x47, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0c, 0x47, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x09, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x53, 0x65, 0x78, 0x45, 0x6d,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x47, 0x72, 0x61, 0x64, 0x75, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x22, 0xf1, 0x02, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x6f,
	0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x67, 0x49,
	0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x61, 0x67, 0x49, 0x64, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x6f, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x6f, 0x6e, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2f,
	0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x22,
	0x20, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x1e, 0x0a, 0x08, 0x53, 0x75, 0x63, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0xac, 0x03, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x6f,
	0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x50, 0x6c, 0x61,
	0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x49, 0x73, 0x4e, 0x65, 0x65, 0x64,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0b, 0x49,
	0x73, 0x4e, 0x65, 0x65, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a,
	0x0b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x02, 0x52, 0x0b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0d, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x04, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x05, 0x52, 0x07, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0c, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x41, 0x64,
	0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x49, 0x73, 0x4e, 0x65, 0x65,
	0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6d, 0x73,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0xbe, 0x01, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0e, 0x67, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x35, 0x0a, 0x0c, 0x65, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x6f, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x6f, 0x6e, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x6f, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(*UserMainInfo)(nil),         // 0: user.UserMainInfo
	(*EditUserInfoReq)(nil),      // 1: user.EditUserInfoReq
	(*UserCronyInfo)(nil),        // 2: user.UserCronyInfo
	(*GetUserCronyListResp)(nil), // 3: user.getUserCronyListResp
	(*IdReq)(nil),                // 4: user.IdReq
	(*TokenReq)(nil),             // 5: user.TokenReq
	(*SuccResp)(nil),             // 6: user.SuccResp
	(*GetUserCronyListReq)(nil),  // 7: user.getUserCronyListReq
}
var file_user_proto_depIdxs = []int32{
	2, // 0: user.getUserCronyListResp.list:type_name -> user.UserCronyInfo
	5, // 1: user.user.getUserByToken:input_type -> user.TokenReq
	1, // 2: user.user.editUserInfo:input_type -> user.EditUserInfoReq
	7, // 3: user.user.getUserCronyList:input_type -> user.getUserCronyListReq
	0, // 4: user.user.getUserByToken:output_type -> user.UserMainInfo
	6, // 5: user.user.editUserInfo:output_type -> user.SuccResp
	3, // 6: user.user.getUserCronyList:output_type -> user.getUserCronyListResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMainInfo); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserInfoReq); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCronyInfo); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserCronyListResp); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReq); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccResp); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserCronyListReq); i {
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
	file_user_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_user_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_user_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
