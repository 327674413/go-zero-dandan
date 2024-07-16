package interceptor

import (
	"context"
	"go-zero-dandan/common/ctxd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"go-zero-dandan/common/utild"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

// RpcClientInterceptor rpc客户端拦截器，自动传数据用
func RpcClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//调用前处理
		// 从上下文中获取现有的metadata
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			// 如果不存在现有的metadata，创建一个新的
			md = metadata.New(nil)
		}
		meta, _ := ctx.Value(ctxd.KeyReqMeta).(*typed.ReqMeta)
		bstr, err := utild.StdToBase64(meta)
		if err == nil {
			md.Append(ctxd.KeyReqMeta, bstr)
		}
		newCtx := metadata.NewOutgoingContext(ctx, md)
		//调用rpc
		err = invoker(newCtx, method, req, reply, cc, opts...)
		// 调用后目前暂时不做处理
		return err
	}
}

// RpcServerInterceptor rpc服务端拦截器，解析数据用
func RpcServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 从上下文中获取 metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, resd.NewErrCtx(ctx, "missing metadata", resd.ErrRpcMissMeta)
		}
		reqMeta := &typed.ReqMeta{}
		//经过打印发现，这里的key竟然都会被变成小写
		if len(md[strings.ToLower(ctxd.KeyReqMeta)]) > 0 {
			metaJson := md[strings.ToLower(ctxd.KeyReqMeta)][0]
			utild.Base64ToStd(metaJson, &reqMeta)
		}
		// 将解析出的数据添加到新的上下文中
		newCtx := context.WithValue(ctx, ctxd.KeyReqMeta, reqMeta)
		// 调用 handler 处理请求
		return handler(newCtx, req)
	}
}
