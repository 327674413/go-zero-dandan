package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/asset/api/internal/config"
	"go-zero-dandan/app/asset/api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	LangMiddleware rest.Middleware

	Mode string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Mode:           c.Mode,
		LangMiddleware: middleware.NewLangMiddleware().Handle,
	}
}
