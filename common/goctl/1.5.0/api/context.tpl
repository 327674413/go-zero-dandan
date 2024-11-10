package svc

import (
	{{.configImport}}
	"go-zero-dandan/app/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/common/interceptor"
)

type ServiceContext struct {
	Config {{.config}}
	UserRpc user.User
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
    UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	return &ServiceContext{
		Config: c,
		UserRpc: UserRpc,
		{{.middlewareAssignment}}

	}
}
