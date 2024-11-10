package svc

import (
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/dify/api/internal/config"
	"go-zero-dandan/app/dify/api/internal/middleware"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/interceptor"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config                 config.Config
	Redis                  *redisd.Redisd
	UserRpc                user.User
	ReqRateLimitMiddleware rest.Middleware
	MetaMiddleware         rest.Middleware
	UserInfoMiddleware     rest.Middleware
	UserTokenMiddleware    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "dify")
	userRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	limiter := limit.NewPeriodLimit(c.ReqRateLimitByIpAgent.Seconds, c.ReqRateLimitByIpAgent.Quota, redisConn, c.ReqRateLimitByIpAgent.KeyPrefix)
	return &ServiceContext{
		Config:                 c,
		UserRpc:                userRpc,
		Redis:                  redisdConn,
		ReqRateLimitMiddleware: middleware.NewReqRateLimitMiddleware(limiter).Handle,
		MetaMiddleware:         middleware.NewMetaMiddleware().Handle,
		UserInfoMiddleware:     middleware.NewUserInfoMiddleware(userRpc).Handle,
		UserTokenMiddleware:    middleware.NewUserTokenMiddleware().Handle,
	}
}
