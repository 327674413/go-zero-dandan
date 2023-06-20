package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/app/user/api/internal/config"
	"go-zero-dandan/app/user/api/internal/middleware"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config         config.Config
	SqlConn        sqlx.SqlConn
	MessageRpc     message.Message
	LangMiddleware rest.Middleware
	Redis          *redisd.Redisd
	Mode           string
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "user")

	return &ServiceContext{
		Config:         c,
		SqlConn:        sqlx.NewMysql(c.DB.DataSource),
		MessageRpc:     message.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
		LangMiddleware: middleware.NewLangMiddleware().Handle,
		Redis:          redisdConn,
		Mode:           c.Mode,
	}
}
