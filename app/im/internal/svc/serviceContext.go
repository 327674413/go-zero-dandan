package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/internal/config"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn
	UserRpc user.User
	Redis   *redisd.Redisd
	Mode    string
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "user")
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc))
	return &ServiceContext{
		Config: c,
		//SqlConn:             sqlx.NewMysql(c.DB.DataSource),
		UserRpc: UserRpc,
		Redis:   redisdConn,
		Mode:    c.Mode,
	}
}
