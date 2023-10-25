package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/message/rpc/internal/config"
	"go-zero-dandan/common/queued"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn
	Redis   *redisd.Redisd
	Mode    string
	Pusher  *queued.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "message")
	producer, err := queued.NewProducer(c.KafkaPusherConf.Addrs)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:  c,
		SqlConn: sqlx.NewMysql(c.Db.DataSource),
		Redis:   redisdConn,
		Mode:    c.Mode,
		Pusher:  producer,
	}
}
