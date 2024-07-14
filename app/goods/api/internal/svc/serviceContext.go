package svc

import (
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/goods/api/internal/config"
	"go-zero-dandan/app/goods/api/internal/middleware"
	"go-zero-dandan/app/goods/rpc/goods"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/interceptor"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config                 config.Config
	MetaMiddleware         rest.Middleware
	UserInfoMiddleware     rest.Middleware
	ReqRateLimitMiddleware rest.Middleware
	SqlConn                sqlx.SqlConn
	Redis                  *redisd.Redisd
	GoodsRpc               goods.Goods
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "goods")
	userRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	limiter := limit.NewPeriodLimit(c.ReqRateLimitByIpAgent.Seconds, c.ReqRateLimitByIpAgent.Quota, redisConn, c.ReqRateLimitByIpAgent.KeyPrefix)
	return &ServiceContext{
		Config:                 c,
		SqlConn:                sqlx.NewMysql(c.DB.DataSource),
		Redis:                  redisdConn,
		MetaMiddleware:         middleware.NewMetaMiddleware().Handle,
		UserInfoMiddleware:     middleware.NewUserInfoMiddleware(userRpc).Handle,
		ReqRateLimitMiddleware: middleware.NewReqRateLimitMiddleware(limiter).Handle,
		GoodsRpc:               goods.NewGoods(zrpc.MustNewClient(c.GoodsRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor()))),
	}
}
