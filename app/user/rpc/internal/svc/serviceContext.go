package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/user/rpc/internal/config"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn
	Redis   *redisd.Redisd
	Mode    string
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "user")
	return &ServiceContext{
		Config:  c,
		SqlConn: sqlx.NewMysql(c.Db.DataSource),
		Redis:   redisdConn,
		Mode:    c.Mode,
	}
}
