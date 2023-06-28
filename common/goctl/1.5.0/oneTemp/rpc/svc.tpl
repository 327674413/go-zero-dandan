package svc

import {{.imports}}

type ServiceContext struct {
	Config config.Config
	Mode   string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:c,
		Mode:    c.Mode,
	}
}
