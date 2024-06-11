package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/plat/rpc/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		SqlConn: sqlx.NewMysql(c.DB.DataSource),
	}
}
