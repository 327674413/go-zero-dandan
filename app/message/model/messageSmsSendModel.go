package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MessageSmsSendModel = (*customMessageSmsSendModel)(nil)

type (
	// MessageSmsSendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageSmsSendModel.
	MessageSmsSendModel interface {
		messageSmsSendModel
	}

	customMessageSmsSendModel struct {
		*defaultMessageSmsSendModel
		SoftDeletable bool
	}
)

// NewMessageSmsSendModel returns a model for the database table.
func NewMessageSmsSendModel(conn sqlx.SqlConn) MessageSmsSendModel {
	return &customMessageSmsSendModel{
		defaultMessageSmsSendModel: newMessageSmsSendModel(conn),
		SoftDeletable:              true, //是否启用软删除
	}
}
