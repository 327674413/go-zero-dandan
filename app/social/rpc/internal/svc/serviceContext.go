package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/social/rpc/internal/config"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redisd.Redisd
	SqlConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis := redis.MustNewRedis(c.RedisConf)
	redisd := redisd.NewRedisd(redis, "social")
	return &ServiceContext{
		Config:  c,
		Redis:   redisd,
		SqlConn: sqlx.NewMysql(c.DB.DataSource),
	}
}
