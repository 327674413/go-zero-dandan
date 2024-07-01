package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/rpc/im"
	"go-zero-dandan/app/message/rpc/internal/config"
	"go-zero-dandan/common/interceptor"
	"go-zero-dandan/common/queued"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config    config.Config
	SqlConn   sqlx.SqlConn
	Redis     *redisd.Redisd
	Mode      string
	SmsPusher *queued.Producer
	ImRpc     im.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "message")
	smsPusher, err := queued.NewProducer(c.KqSmsPusher.Addrs)
	ImRpc := im.NewIm(zrpc.MustNewClient(c.ImRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		SqlConn:   sqlx.NewMysql(c.Db.DataSource),
		Redis:     redisdConn,
		Mode:      c.Mode,
		SmsPusher: smsPusher,
		ImRpc:     ImRpc,
	}
}
