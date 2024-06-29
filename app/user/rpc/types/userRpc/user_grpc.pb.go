// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: user.proto

package userRpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	User_GetUserByToken_FullMethodName    = "/user.user/getUserByToken"
	User_EditUserInfo_FullMethodName      = "/user.user/editUserInfo"
	User_RegByAccount_FullMethodName      = "/user.user/regByAccount"
	User_GetUserById_FullMethodName       = "/user.user/getUserById"
	User_GetUserPage_FullMethodName       = "/user.user/getUserPage"
	User_BindUnionUser_FullMethodName     = "/user.user/bindUnionUser"
	User_GetUserNormalInfo_FullMethodName = "/user.user/getUserNormalInfo"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetUserByToken(ctx context.Context, in *TokenReq, opts ...grpc.CallOption) (*UserMainInfo, error)
	EditUserInfo(ctx context.Context, in *EditUserInfoReq, opts ...grpc.CallOption) (*SuccResp, error)
	RegByAccount(ctx context.Context, in *RegByAccountReq, opts ...grpc.CallOption) (*LoginResp, error)
	GetUserById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserMainInfo, error)
	GetUserPage(ctx context.Context, in *GetUserPageReq, opts ...grpc.CallOption) (*GetUserPageResp, error)
	BindUnionUser(ctx context.Context, in *BindUnionUserReq, opts ...grpc.CallOption) (*BindUnionUserResp, error)
	GetUserNormalInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserNormalInfoResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserByToken(ctx context.Context, in *TokenReq, opts ...grpc.CallOption) (*UserMainInfo, error) {
	out := new(UserMainInfo)
	err := c.cc.Invoke(ctx, User_GetUserByToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) EditUserInfo(ctx context.Context, in *EditUserInfoReq, opts ...grpc.CallOption) (*SuccResp, error) {
	out := new(SuccResp)
	err := c.cc.Invoke(ctx, User_EditUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RegByAccount(ctx context.Context, in *RegByAccountReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, User_RegByAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserMainInfo, error) {
	out := new(UserMainInfo)
	err := c.cc.Invoke(ctx, User_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserPage(ctx context.Context, in *GetUserPageReq, opts ...grpc.CallOption) (*GetUserPageResp, error) {
	out := new(GetUserPageResp)
	err := c.cc.Invoke(ctx, User_GetUserPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BindUnionUser(ctx context.Context, in *BindUnionUserReq, opts ...grpc.CallOption) (*BindUnionUserResp, error) {
	out := new(BindUnionUserResp)
	err := c.cc.Invoke(ctx, User_BindUnionUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserNormalInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserNormalInfoResp, error) {
	out := new(GetUserNormalInfoResp)
	err := c.cc.Invoke(ctx, User_GetUserNormalInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetUserByToken(context.Context, *TokenReq) (*UserMainInfo, error)
	EditUserInfo(context.Context, *EditUserInfoReq) (*SuccResp, error)
	RegByAccount(context.Context, *RegByAccountReq) (*LoginResp, error)
	GetUserById(context.Context, *IdReq) (*UserMainInfo, error)
	GetUserPage(context.Context, *GetUserPageReq) (*GetUserPageResp, error)
	BindUnionUser(context.Context, *BindUnionUserReq) (*BindUnionUserResp, error)
	GetUserNormalInfo(context.Context, *GetUserInfoReq) (*GetUserNormalInfoResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUserByToken(context.Context, *TokenReq) (*UserMainInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByToken not implemented")
}
func (UnimplementedUserServer) EditUserInfo(context.Context, *EditUserInfoReq) (*SuccResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserInfo not implemented")
}
func (UnimplementedUserServer) RegByAccount(context.Context, *RegByAccountReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegByAccount not implemented")
}
func (UnimplementedUserServer) GetUserById(context.Context, *IdReq) (*UserMainInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServer) GetUserPage(context.Context, *GetUserPageReq) (*GetUserPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPage not implemented")
}
func (UnimplementedUserServer) BindUnionUser(context.Context, *BindUnionUserReq) (*BindUnionUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindUnionUser not implemented")
}
func (UnimplementedUserServer) GetUserNormalInfo(context.Context, *GetUserInfoReq) (*GetUserNormalInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserNormalInfo not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUserByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserByToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByToken(ctx, req.(*TokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_EditUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EditUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_EditUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EditUserInfo(ctx, req.(*EditUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RegByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegByAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RegByAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegByAccount(ctx, req.(*RegByAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserPage(ctx, req.(*GetUserPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BindUnionUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindUnionUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BindUnionUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_BindUnionUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BindUnionUser(ctx, req.(*BindUnionUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserNormalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserNormalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserNormalInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserNormalInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUserByToken",
			Handler:    _User_GetUserByToken_Handler,
		},
		{
			MethodName: "editUserInfo",
			Handler:    _User_EditUserInfo_Handler,
		},
		{
			MethodName: "regByAccount",
			Handler:    _User_RegByAccount_Handler,
		},
		{
			MethodName: "getUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "getUserPage",
			Handler:    _User_GetUserPage_Handler,
		},
		{
			MethodName: "bindUnionUser",
			Handler:    _User_BindUnionUser_Handler,
		},
		{
			MethodName: "getUserNormalInfo",
			Handler:    _User_GetUserNormalInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}