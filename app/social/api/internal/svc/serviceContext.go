package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/social/api/internal/config"
	"go-zero-dandan/app/social/api/internal/middleware"
	"go-zero-dandan/app/social/rpc/social"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/interceptor"
)

type ServiceContext struct {
	Config              config.Config
	UserInfoMiddleware  rest.Middleware
	UserTokenMiddleware rest.Middleware
	LangMiddleware      rest.Middleware
	SocialRpc           social.Social
	UserRpc             user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	SocialRpc := social.NewSocial(zrpc.MustNewClient(c.SocialRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))

	return &ServiceContext{
		Config:              c,
		UserInfoMiddleware:  middleware.NewUserInfoMiddleware(UserRpc).Handle,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		LangMiddleware:      middleware.NewLangMiddleware().Handle,
		UserRpc:             UserRpc,
		SocialRpc:           SocialRpc,
	}
}
