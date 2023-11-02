package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/goods/api/internal/config"
	"go-zero-dandan/app/goods/api/internal/middleware"
	"go-zero-dandan/common/redisd"
)

type ServiceContext struct {
	Config             config.Config
	LangMiddleware     rest.Middleware
	UserInfoMiddleware rest.Middleware
	SqlConn            sqlx.SqlConn
	Redis              *redisd.Redisd
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "goods")
	return &ServiceContext{
		Config:             c,
		SqlConn:            sqlx.NewMysql(c.DB.DataSource),
		Redis:              redisdConn,
		LangMiddleware:     middleware.NewLangMiddleware().Handle,
		UserInfoMiddleware: middleware.NewUserInfoMiddleware().Handle,
	}
}
