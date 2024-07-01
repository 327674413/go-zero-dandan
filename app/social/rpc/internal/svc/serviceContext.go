package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/rpc/im"
	"go-zero-dandan/app/social/rpc/internal/config"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/interceptor"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redisd.Redisd
	SqlConn sqlx.SqlConn
	UserRpc user.User
	ImRpc   im.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis := redis.MustNewRedis(c.RedisConf)
	redisd := redisd.NewRedisd(redis, "social")
	userRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	imRpc := im.NewIm(zrpc.MustNewClient(c.ImRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	return &ServiceContext{
		Config:  c,
		Redis:   redisd,
		SqlConn: sqlx.NewMysql(c.DB.DataSource),
		UserRpc: userRpc,
		ImRpc:   imRpc,
	}
}
