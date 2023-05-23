package svc

import (
	"go-zero-dandan/common/model"
	"go-zero-dandan/user/api/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		DB:     model.DB,
	}
}
