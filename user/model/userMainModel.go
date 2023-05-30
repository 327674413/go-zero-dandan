package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserMainModel = (*customUserMainModel)(nil)

type (
	// UserMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMainModel.
	UserMainModel interface {
		userMainModel
	}

	customUserMainModel struct {
		*defaultUserMainModel
	}
)

// NewUserMainModel returns a model for the database table.
func NewUserMainModel(conn sqlx.SqlConn) UserMainModel {
	return &customUserMainModel{
		defaultUserMainModel: newUserMainModel(conn),
	}
}
