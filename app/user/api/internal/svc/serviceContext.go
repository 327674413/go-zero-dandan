package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/user/api/internal/config"
	"go-zero-dandan/app/user/model"
)

type ServiceContext struct {
	Config        config.Config
	UserMainModel model.UserMainModel
	UserInfoModel model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:        c,
		UserMainModel: model.NewUserMainModel(sqlx.NewMysql(c.DB.DataSource)),
		UserInfoModel: model.NewUserInfoModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
