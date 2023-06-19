package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/plat/api/internal/config"
	"go-zero-dandan/app/plat/api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	LangMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		LangMiddleware: middleware.NewLangMiddleware().Handle,
	}
}
