package interceptor

import (
	"context"
	"go-zero-dandan/common/resd"
	"google.golang.org/grpc"
)

// RpcClientInterceptor rpc客户端拦截器，解析返回值
func RpcClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		// 将rpc的error转成我自己的error
		if err != nil {
			err = resd.RpcErrDecode(err)
		} else {
			//如果正常返回会进这里，是否可以统一封装变成中间件还未知，目前也没该需求
		}
		return err
	}
}
