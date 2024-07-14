package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/api/internal/config"
	"go-zero-dandan/app/im/api/internal/middleware"
	"go-zero-dandan/app/social/rpc/social"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/interceptor"
)

type ServiceContext struct {
	Config              config.Config
	MetaMiddleware      rest.Middleware
	UserInfoMiddleware  rest.Middleware
	UserTokenMiddleware rest.Middleware
	SocialRpc           social.Social
	UserRpc             user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	SocialRpc := social.NewSocial(zrpc.MustNewClient(c.SocialRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))

	return &ServiceContext{
		Config:              c,
		MetaMiddleware:      middleware.NewMetaMiddleware().Handle,
		UserInfoMiddleware:  middleware.NewUserInfoMiddleware(UserRpc).Handle,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		UserRpc:             UserRpc,
		SocialRpc:           SocialRpc,
	}
}
