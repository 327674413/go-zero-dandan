package svc

import (
	{{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}
	{{.middleware}}
	Mode    string
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mode:    c.Mode,
		{{.middlewareAssignment}}
	}
}
