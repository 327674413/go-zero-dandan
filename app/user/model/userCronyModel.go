package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserCronyModel = (*customUserCronyModel)(nil)

type (
	// UserCronyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCronyModel.
	UserCronyModel interface {
		userCronyModel
	}

	customUserCronyModel struct {
		*defaultUserCronyModel
		softDeletable bool
	}
)

// NewUserCronyModel returns a model for the database table.
func NewUserCronyModel(conn sqlx.SqlConn, platId ...int64) UserCronyModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customUserCronyModel{
		defaultUserCronyModel: newUserCronyModel(conn, platid),
		softDeletable:         true, //是否启用软删除
	}
}
