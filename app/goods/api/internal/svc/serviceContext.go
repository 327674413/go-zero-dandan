package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/goods/api/internal/config"
	"go-zero-dandan/app/goods/api/internal/middleware"
	"go-zero-dandan/app/goods/rpc/goods"
	"go-zero-dandan/common/interceptor"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config             config.Config
	LangMiddleware     rest.Middleware
	UserInfoMiddleware rest.Middleware
	SqlConn            sqlx.SqlConn
	Redis              *redisd.Redisd
	GoodsRpc           goods.Goods
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "goods")

	return &ServiceContext{
		Config:             c,
		SqlConn:            sqlx.NewMysql(c.DB.DataSource),
		Redis:              redisdConn,
		LangMiddleware:     middleware.NewLangMiddleware().Handle,
		UserInfoMiddleware: middleware.NewUserInfoMiddleware().Handle,
		GoodsRpc:           goods.NewGoods(zrpc.MustNewClient(c.GoodsRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor()))),
	}
}
