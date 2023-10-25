// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.19.4
// source: wechat.proto

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

type AuthByCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *AuthByCodeReq) Reset() {
	*x = AuthByCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthByCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByCodeReq) ProtoMessage() {}

func (x *AuthByCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByCodeReq.ProtoReflect.Descriptor instead.
func (*AuthByCodeReq) Descriptor() ([]byte, []int) {
	return file_wechat_proto_rawDescGZIP(), []int{0}
}

func (x *AuthByCodeReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AuthByCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	Openid      string `protobuf:"bytes,2,opt,name=Openid,proto3" json:"Openid,omitempty"`
	Unionid     string `protobuf:"bytes,3,opt,name=Unionid,proto3" json:"Unionid,omitempty"`
}

func (x *AuthByCodeResp) Reset() {
	*x = AuthByCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthByCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByCodeResp) ProtoMessage() {}

func (x *AuthByCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_wechat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByCodeResp.ProtoReflect.Descriptor instead.
func (*AuthByCodeResp) Descriptor() ([]byte, []int) {
	return file_wechat_proto_rawDescGZIP(), []int{1}
}

func (x *AuthByCodeResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthByCodeResp) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *AuthByCodeResp) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

var File_wechat_proto protoreflect.FileDescriptor

var file_wechat_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x22, 0x23, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x64, 0x0a, 0x0e, 0x41,
	0x75, 0x74, 0x68, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x6e, 0x69, 0x6f, 0x6e,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x69,
	0x64, 0x32, 0x48, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0f, 0x77, 0x78, 0x70,
	0x75, 0x62, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x77,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wechat_proto_rawDescOnce sync.Once
	file_wechat_proto_rawDescData = file_wechat_proto_rawDesc
)

func file_wechat_proto_rawDescGZIP() []byte {
	file_wechat_proto_rawDescOnce.Do(func() {
		file_wechat_proto_rawDescData = protoimpl.X.CompressGZIP(file_wechat_proto_rawDescData)
	})
	return file_wechat_proto_rawDescData
}

var file_wechat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wechat_proto_goTypes = []interface{}{
	(*AuthByCodeReq)(nil),  // 0: wechat.AuthByCodeReq
	(*AuthByCodeResp)(nil), // 1: wechat.AuthByCodeResp
}
var file_wechat_proto_depIdxs = []int32{
	0, // 0: wechat.user.wxpubAuthByCode:input_type -> wechat.AuthByCodeReq
	1, // 1: wechat.user.wxpubAuthByCode:output_type -> wechat.AuthByCodeResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wechat_proto_init() }
func file_wechat_proto_init() {
	if File_wechat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wechat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthByCodeReq); i {
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
		file_wechat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthByCodeResp); i {
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
			RawDescriptor: file_wechat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wechat_proto_goTypes,
		DependencyIndexes: file_wechat_proto_depIdxs,
		MessageInfos:      file_wechat_proto_msgTypes,
	}.Build()
	File_wechat_proto = out.File
	file_wechat_proto_rawDesc = nil
	file_wechat_proto_goTypes = nil
	file_wechat_proto_depIdxs = nil
}