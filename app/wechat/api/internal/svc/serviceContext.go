package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/wechat/api/internal/config"
	"go-zero-dandan/app/wechat/api/internal/middleware"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config         config.Config
	Redis          *redisd.Redisd
	MetaMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "wechat")
	return &ServiceContext{
		Config:         c,
		Redis:          redisdConn,
		MetaMiddleware: middleware.NewLangMiddleware().Handle,
	}
}
