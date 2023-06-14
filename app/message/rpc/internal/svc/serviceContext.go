package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/message/rpc/internal/config"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn
	Redis   *redisd.Redisd
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "message")
	return &ServiceContext{
		Config:  c,
		SqlConn: sqlx.NewMysql(c.Db.DataSource),
		Redis:   redisdConn,
	}
}
