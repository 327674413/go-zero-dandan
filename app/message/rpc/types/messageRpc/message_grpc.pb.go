// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: message.proto

package messageRpc

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
	Message_SendPhone_FullMethodName             = "/message.message/sendPhone"
	Message_SendPhoneAsync_FullMethodName        = "/message.message/sendPhoneAsync"
	Message_SendImChannelMsg_FullMethodName      = "/message.message/sendImChannelMsg"
	Message_SendImChannelMsgAsync_FullMethodName = "/message.message/sendImChannelMsgAsync"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	SendPhone(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*ResultResp, error)
	SendPhoneAsync(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*ResultResp, error)
	SendImChannelMsg(ctx context.Context, in *SendImChannelMsgReq, opts ...grpc.CallOption) (*ResultResp, error)
	SendImChannelMsgAsync(ctx context.Context, in *SendImChannelMsgReq, opts ...grpc.CallOption) (*ResultResp, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) SendPhone(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*ResultResp, error) {
	out := new(ResultResp)
	err := c.cc.Invoke(ctx, Message_SendPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendPhoneAsync(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*ResultResp, error) {
	out := new(ResultResp)
	err := c.cc.Invoke(ctx, Message_SendPhoneAsync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendImChannelMsg(ctx context.Context, in *SendImChannelMsgReq, opts ...grpc.CallOption) (*ResultResp, error) {
	out := new(ResultResp)
	err := c.cc.Invoke(ctx, Message_SendImChannelMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendImChannelMsgAsync(ctx context.Context, in *SendImChannelMsgReq, opts ...grpc.CallOption) (*ResultResp, error) {
	out := new(ResultResp)
	err := c.cc.Invoke(ctx, Message_SendImChannelMsgAsync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	SendPhone(context.Context, *SendPhoneReq) (*ResultResp, error)
	SendPhoneAsync(context.Context, *SendPhoneReq) (*ResultResp, error)
	SendImChannelMsg(context.Context, *SendImChannelMsgReq) (*ResultResp, error)
	SendImChannelMsgAsync(context.Context, *SendImChannelMsgReq) (*ResultResp, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) SendPhone(context.Context, *SendPhoneReq) (*ResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhone not implemented")
}
func (UnimplementedMessageServer) SendPhoneAsync(context.Context, *SendPhoneReq) (*ResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneAsync not implemented")
}
func (UnimplementedMessageServer) SendImChannelMsg(context.Context, *SendImChannelMsgReq) (*ResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendImChannelMsg not implemented")
}
func (UnimplementedMessageServer) SendImChannelMsgAsync(context.Context, *SendImChannelMsgReq) (*ResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendImChannelMsgAsync not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_SendPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendPhone(ctx, req.(*SendPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendPhoneAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendPhoneAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendPhoneAsync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendPhoneAsync(ctx, req.(*SendPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendImChannelMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendImChannelMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendImChannelMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendImChannelMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendImChannelMsg(ctx, req.(*SendImChannelMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendImChannelMsgAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendImChannelMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendImChannelMsgAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendImChannelMsgAsync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendImChannelMsgAsync(ctx, req.(*SendImChannelMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendPhone",
			Handler:    _Message_SendPhone_Handler,
		},
		{
			MethodName: "sendPhoneAsync",
			Handler:    _Message_SendPhoneAsync_Handler,
		},
		{
			MethodName: "sendImChannelMsg",
			Handler:    _Message_SendImChannelMsg_Handler,
		},
		{
			MethodName: "sendImChannelMsgAsync",
			Handler:    _Message_SendImChannelMsgAsync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}