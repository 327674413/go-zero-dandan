package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/plat/api/internal/config"
	"go-zero-dandan/app/plat/api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	MetaMiddleware rest.Middleware
	SqlConn        sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		SqlConn:        sqlx.NewMysql(c.DB.DataSource),
		MetaMiddleware: middleware.NewMetaMiddleware().Handle,
	}
}
