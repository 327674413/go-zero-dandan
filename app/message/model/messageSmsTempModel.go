package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MessageSmsTempModel = (*customMessageSmsTempModel)(nil)

type (
	// MessageSmsTempModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageSmsTempModel.
	MessageSmsTempModel interface {
		messageSmsTempModel
	}

	customMessageSmsTempModel struct {
		*defaultMessageSmsTempModel
		SoftDeletable bool
	}
)

// NewMessageSmsTempModel returns a model for the database table.
func NewMessageSmsTempModel(conn sqlx.SqlConn) MessageSmsTempModel {
	return &customMessageSmsTempModel{
		defaultMessageSmsTempModel: newMessageSmsTempModel(conn),
		SoftDeletable:              true, //是否启用软删除
	}
}
