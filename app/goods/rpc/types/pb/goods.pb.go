// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.19.4
// source: goods.proto

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

type GoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Spec      string `protobuf:"bytes,3,opt,name=Spec,proto3" json:"Spec,omitempty"`
	Cover     string `protobuf:"bytes,4,opt,name=Cover,proto3" json:"Cover,omitempty"`
	SellPrice int64  `protobuf:"varint,5,opt,name=SellPrice,proto3" json:"SellPrice,omitempty"`
	StoreQty  int64  `protobuf:"varint,6,opt,name=StoreQty,proto3" json:"StoreQty,omitempty"`
	State     int64  `protobuf:"varint,7,opt,name=State,proto3" json:"State,omitempty"`
	IsSpecial int64  `protobuf:"varint,8,opt,name=IsSpecial,proto3" json:"IsSpecial,omitempty"`
	UnitId    int64  `protobuf:"varint,9,opt,name=UnitId,proto3" json:"UnitId,omitempty"`
	UnitName  string `protobuf:"bytes,10,opt,name=UnitName,proto3" json:"UnitName,omitempty"`
	PlatId    int64  `protobuf:"varint,11,opt,name=PlatId,proto3" json:"PlatId,omitempty"`
	ViewNum   int64  `protobuf:"varint,12,opt,name=ViewNum,proto3" json:"ViewNum,omitempty"`
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsInfo) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *GoodsInfo) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *GoodsInfo) GetSellPrice() int64 {
	if x != nil {
		return x.SellPrice
	}
	return 0
}

func (x *GoodsInfo) GetStoreQty() int64 {
	if x != nil {
		return x.StoreQty
	}
	return 0
}

func (x *GoodsInfo) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *GoodsInfo) GetIsSpecial() int64 {
	if x != nil {
		return x.IsSpecial
	}
	return 0
}

func (x *GoodsInfo) GetUnitId() int64 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *GoodsInfo) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *GoodsInfo) GetPlatId() int64 {
	if x != nil {
		return x.PlatId
	}
	return 0
}

func (x *GoodsInfo) GetViewNum() int64 {
	if x != nil {
		return x.ViewNum
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
		mi := &file_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[1]
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
	return file_goods_proto_rawDescGZIP(), []int{1}
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

type GetPageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int64  `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Size      int64  `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	Sort      string `protobuf:"bytes,3,opt,name=Sort,proto3" json:"Sort,omitempty"`
	PlatId    int64  `protobuf:"varint,4,opt,name=PlatId,proto3" json:"PlatId,omitempty"`
	TotalFlag int64  `protobuf:"varint,5,opt,name=TotalFlag,proto3" json:"TotalFlag,omitempty"`
}

func (x *GetPageReq) Reset() {
	*x = GetPageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageReq) ProtoMessage() {}

func (x *GetPageReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageReq.ProtoReflect.Descriptor instead.
func (*GetPageReq) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{2}
}

func (x *GetPageReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPageReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetPageReq) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetPageReq) GetPlatId() int64 {
	if x != nil {
		return x.PlatId
	}
	return 0
}

func (x *GetPageReq) GetTotalFlag() int64 {
	if x != nil {
		return x.TotalFlag
	}
	return 0
}

type GetHotPageByCursorReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int64 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Size   int64 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	PlatId int64 `protobuf:"varint,3,opt,name=PlatId,proto3" json:"PlatId,omitempty"`
	Cursor int64 `protobuf:"varint,4,opt,name=Cursor,proto3" json:"Cursor,omitempty"`
	LastId int64 `protobuf:"varint,5,opt,name=LastId,proto3" json:"LastId,omitempty"`
}

func (x *GetHotPageByCursorReq) Reset() {
	*x = GetHotPageByCursorReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotPageByCursorReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotPageByCursorReq) ProtoMessage() {}

func (x *GetHotPageByCursorReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotPageByCursorReq.ProtoReflect.Descriptor instead.
func (*GetHotPageByCursorReq) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{3}
}

func (x *GetHotPageByCursorReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetHotPageByCursorReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetHotPageByCursorReq) GetPlatId() int64 {
	if x != nil {
		return x.PlatId
	}
	return 0
}

func (x *GetHotPageByCursorReq) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *GetHotPageByCursorReq) GetLastId() int64 {
	if x != nil {
		return x.LastId
	}
	return 0
}

type GetPageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int64        `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Size    int64        `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	List    []*GoodsInfo `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	IsCache bool         `protobuf:"varint,4,opt,name=IsCache,proto3" json:"IsCache,omitempty"`
	Total   *int64       `protobuf:"varint,5,opt,name=Total,proto3,oneof" json:"Total,omitempty"`
	LastId  *int64       `protobuf:"varint,6,opt,name=LastId,proto3,oneof" json:"LastId,omitempty"`
	Cursor  *int64       `protobuf:"varint,7,opt,name=cursor,proto3,oneof" json:"cursor,omitempty"`
	IsEnd   *bool        `protobuf:"varint,8,opt,name=IsEnd,proto3,oneof" json:"IsEnd,omitempty"`
}

func (x *GetPageResp) Reset() {
	*x = GetPageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageResp) ProtoMessage() {}

func (x *GetPageResp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageResp.ProtoReflect.Descriptor instead.
func (*GetPageResp) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{4}
}

func (x *GetPageResp) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPageResp) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetPageResp) GetList() []*GoodsInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *GetPageResp) GetIsCache() bool {
	if x != nil {
		return x.IsCache
	}
	return false
}

func (x *GetPageResp) GetTotal() int64 {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return 0
}

func (x *GetPageResp) GetLastId() int64 {
	if x != nil && x.LastId != nil {
		return *x.LastId
	}
	return 0
}

func (x *GetPageResp) GetCursor() int64 {
	if x != nil && x.Cursor != nil {
		return *x.Cursor
	}
	return 0
}

func (x *GetPageResp) GetIsEnd() bool {
	if x != nil && x.IsEnd != nil {
		return *x.IsEnd
	}
	return false
}

type GetPageByCursorResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size    int64        `protobuf:"varint,1,opt,name=Size,proto3" json:"Size,omitempty"`
	IsCache bool         `protobuf:"varint,2,opt,name=IsCache,proto3" json:"IsCache,omitempty"`
	IsEnd   bool         `protobuf:"varint,3,opt,name=IsEnd,proto3" json:"IsEnd,omitempty"`
	LastId  int64        `protobuf:"varint,4,opt,name=LastId,proto3" json:"LastId,omitempty"`
	Cursor  int64        `protobuf:"varint,5,opt,name=cursor,proto3" json:"cursor,omitempty"`
	List    []*GoodsInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetPageByCursorResp) Reset() {
	*x = GetPageByCursorResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageByCursorResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageByCursorResp) ProtoMessage() {}

func (x *GetPageByCursorResp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageByCursorResp.ProtoReflect.Descriptor instead.
func (*GetPageByCursorResp) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{5}
}

func (x *GetPageByCursorResp) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetPageByCursorResp) GetIsCache() bool {
	if x != nil {
		return x.IsCache
	}
	return false
}

func (x *GetPageByCursorResp) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *GetPageByCursorResp) GetLastId() int64 {
	if x != nil {
		return x.LastId
	}
	return 0
}

func (x *GetPageByCursorResp) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *GetPageByCursorResp) GetList() []*GoodsInfo {
	if x != nil {
		return x.List
	}
	return nil
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
		mi := &file_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccResp) ProtoMessage() {}

func (x *SuccResp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[6]
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
	return file_goods_proto_rawDescGZIP(), []int{6}
}

func (x *SuccResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_goods_proto protoreflect.FileDescriptor

var file_goods_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x22, 0xad, 0x02, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x51, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x51, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x49, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69,
	0x65, 0x77, 0x4e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x65,
	0x77, 0x4e, 0x75, 0x6d, 0x22, 0x2f, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50,
	0x6c, 0x61, 0x74, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x46, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x42, 0x79, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x22,
	0x8f, 0x02, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x49, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x01, 0x52, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x02, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x49, 0x73, 0x45, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x05, 0x49,
	0x73, 0x45, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x73, 0x45, 0x6e,
	0x64, 0x22, 0xaf, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x42, 0x79, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x49, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x45, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4c,
	0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x24, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x08, 0x53, 0x75, 0x63, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0xb3, 0x01, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x28, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x0c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x74, 0x50, 0x61, 0x67, 0x65, 0x42, 0x79, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x42, 0x79, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_proto_rawDescOnce sync.Once
	file_goods_proto_rawDescData = file_goods_proto_rawDesc
)

func file_goods_proto_rawDescGZIP() []byte {
	file_goods_proto_rawDescOnce.Do(func() {
		file_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_proto_rawDescData)
	})
	return file_goods_proto_rawDescData
}

var file_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_goods_proto_goTypes = []interface{}{
	(*GoodsInfo)(nil),             // 0: goods.GoodsInfo
	(*IdReq)(nil),                 // 1: goods.IdReq
	(*GetPageReq)(nil),            // 2: goods.GetPageReq
	(*GetHotPageByCursorReq)(nil), // 3: goods.GetHotPageByCursorReq
	(*GetPageResp)(nil),           // 4: goods.GetPageResp
	(*GetPageByCursorResp)(nil),   // 5: goods.GetPageByCursorResp
	(*SuccResp)(nil),              // 6: goods.SuccResp
}
var file_goods_proto_depIdxs = []int32{
	0, // 0: goods.GetPageResp.list:type_name -> goods.GoodsInfo
	0, // 1: goods.GetPageByCursorResp.list:type_name -> goods.GoodsInfo
	1, // 2: goods.goods.GetOne:input_type -> goods.IdReq
	2, // 3: goods.goods.GetPage:input_type -> goods.GetPageReq
	3, // 4: goods.goods.GetHotPageByCursor:input_type -> goods.GetHotPageByCursorReq
	0, // 5: goods.goods.GetOne:output_type -> goods.GoodsInfo
	4, // 6: goods.goods.GetPage:output_type -> goods.GetPageResp
	5, // 7: goods.goods.GetHotPageByCursor:output_type -> goods.GetPageByCursorResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_goods_proto_init() }
func file_goods_proto_init() {
	if File_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfo); i {
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
		file_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageReq); i {
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
		file_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHotPageByCursorReq); i {
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
		file_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageResp); i {
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
		file_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageByCursorResp); i {
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
		file_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_goods_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_proto_goTypes,
		DependencyIndexes: file_goods_proto_depIdxs,
		MessageInfos:      file_goods_proto_msgTypes,
	}.Build()
	File_goods_proto = out.File
	file_goods_proto_rawDesc = nil
	file_goods_proto_goTypes = nil
	file_goods_proto_depIdxs = nil
}