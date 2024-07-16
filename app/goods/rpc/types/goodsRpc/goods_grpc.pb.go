// Code generated by goctl. DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: goods.proto

package goodsRpc

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
	Goods_GetOne_FullMethodName             = "/goods.goods/GetOne"
	Goods_GetPage_FullMethodName            = "/goods.goods/GetPage"
	Goods_GetHotPageByCursor_FullMethodName = "/goods.goods/GetHotPageByCursor"
)

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	GetOne(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*GoodsInfo, error)
	GetPage(ctx context.Context, in *GetPageReq, opts ...grpc.CallOption) (*GetPageResp, error)
	GetHotPageByCursor(ctx context.Context, in *GetHotPageByCursorReq, opts ...grpc.CallOption) (*GetPageByCursorResp, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetOne(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*GoodsInfo, error) {
	out := new(GoodsInfo)
	err := c.cc.Invoke(ctx, Goods_GetOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetPage(ctx context.Context, in *GetPageReq, opts ...grpc.CallOption) (*GetPageResp, error) {
	out := new(GetPageResp)
	err := c.cc.Invoke(ctx, Goods_GetPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetHotPageByCursor(ctx context.Context, in *GetHotPageByCursorReq, opts ...grpc.CallOption) (*GetPageByCursorResp, error) {
	out := new(GetPageByCursorResp)
	err := c.cc.Invoke(ctx, Goods_GetHotPageByCursor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	GetOne(context.Context, *IdReq) (*GoodsInfo, error)
	GetPage(context.Context, *GetPageReq) (*GetPageResp, error)
	GetHotPageByCursor(context.Context, *GetHotPageByCursorReq) (*GetPageByCursorResp, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) GetOne(context.Context, *IdReq) (*GoodsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedGoodsServer) GetPage(context.Context, *GetPageReq) (*GetPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (UnimplementedGoodsServer) GetHotPageByCursor(context.Context, *GetHotPageByCursorReq) (*GetPageByCursorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotPageByCursor not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetOne(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetPage(ctx, req.(*GetPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetHotPageByCursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotPageByCursorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetHotPageByCursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetHotPageByCursor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetHotPageByCursor(ctx, req.(*GetHotPageByCursorReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOne",
			Handler:    _Goods_GetOne_Handler,
		},
		{
			MethodName: "GetPage",
			Handler:    _Goods_GetPage_Handler,
		},
		{
			MethodName: "GetHotPageByCursor",
			Handler:    _Goods_GetHotPageByCursor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}
