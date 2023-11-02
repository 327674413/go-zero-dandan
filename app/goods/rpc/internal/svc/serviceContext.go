package svc

import "go-zero-dandan/app/goods/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
	Mode   string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mode:   c.Mode,
	}
}
