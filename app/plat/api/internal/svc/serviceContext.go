package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/plat/api/internal/config"
	"go-zero-dandan/app/plat/api/internal/middleware"
	"go-zero-dandan/common/resd"
)

type ServiceContext struct {
	Config         config.Config
	LangMiddleware rest.Middleware
	SqlConn        sqlx.SqlConn
	I18n           *resd.I18n
}

func NewServiceContext(c config.Config) *ServiceContext {
	i18n, err := resd.NewI18n(&resd.I18nConfig{
		LangPathList: c.I18n.Langs,
		DefaultLang:  c.I18n.Default,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		I18n:           i18n,
		SqlConn:        sqlx.NewMysql(c.DB.DataSource),
		LangMiddleware: middleware.NewLangMiddleware().Handle,
	}
}
