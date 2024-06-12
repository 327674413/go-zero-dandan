package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserUnionModel = (*customUserUnionModel)(nil)

type (
	// UserUnionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserUnionModel.
	UserUnionModel interface {
		userUnionModel
	}

	customUserUnionModel struct {
		*defaultUserUnionModel
		softDeletable bool
	}
)

// NewUserUnionModel returns a model for the database table.
func NewUserUnionModel(conn sqlx.SqlConn, platId ...string) UserUnionModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customUserUnionModel{
		defaultUserUnionModel: newUserUnionModel(conn, platid),
		softDeletable:         true, //是否启用软删除
	}
}
