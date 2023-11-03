package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/goods/rpc/internal/config"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config  config.Config
	Mode    string
	Redis   *redisd.Redisd
	SqlConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis := redis.MustNewRedis(c.RedisConf)
	redisd := redisd.NewRedisd(redis, "goods")
	return &ServiceContext{
		Config:  c,
		Mode:    c.Mode,
		Redis:   redisd,
		SqlConn: sqlx.NewMysql(c.DB.DataSource),
	}
}
