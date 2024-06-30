package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/message/mq/mqClient"
	"go-zero-dandan/app/message/rpc/internal/config"
	"go-zero-dandan/common/queued"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config    config.Config
	SqlConn   sqlx.SqlConn
	Redis     *redisd.Redisd
	Mode      string
	SmsPusher *queued.Producer
	mqClient.ImSendCli
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "message")
	smsPusher, err := queued.NewProducer(c.KqSmsPusher.Addrs)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		SqlConn:   sqlx.NewMysql(c.Db.DataSource),
		Redis:     redisdConn,
		Mode:      c.Mode,
		SmsPusher: smsPusher,
		ImSendCli: mqClient.NewImSendCli(c.KqImPusher.Addrs, c.KqImPusher.Topic),
	}
}
