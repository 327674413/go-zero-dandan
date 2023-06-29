package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserMainModel = (*customUserMainModel)(nil)

type (
	// UserMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMainModel.
	UserMainModel interface {
		userMainModel
	}

	customUserMainModel struct {
		*defaultUserMainModel
		SoftDeletable bool
	}
)

// NewUserMainModel returns a model for the database table.
func NewUserMainModel(conn sqlx.SqlConn, platId ...int64) UserMainModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customUserMainModel{
		defaultUserMainModel: newUserMainModel(conn, platid),
		SoftDeletable:        true, //是否启用软删除
	}
}
