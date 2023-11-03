package svc

import (
	"go-zero-dandan/app/message/mq/internal/config"
	"go-zero-dandan/common/es"
)

type ServiceContext struct {
	Config config.Config
	Es     *es.Es
}

func NewServiceContext(c config.Config) *ServiceContext {
	es := es.MustNewEs(&es.Config{
		Addresses: c.EsConf.Addresses,
		Username:  c.EsConf.Username,
		Password:  c.EsConf.Password,
	})
	return &ServiceContext{
		Config: c,
		Es:     es,
	}
}
