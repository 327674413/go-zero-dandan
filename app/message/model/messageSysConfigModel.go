package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MessageSysConfigModel = (*customMessageSysConfigModel)(nil)

type (
	// MessageSysConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageSysConfigModel.
	MessageSysConfigModel interface {
		messageSysConfigModel
	}

	customMessageSysConfigModel struct {
		*defaultMessageSysConfigModel
		SoftDeletable bool
	}
)

// NewMessageSysConfigModel returns a model for the database table.
func NewMessageSysConfigModel(conn sqlx.SqlConn) MessageSysConfigModel {
	return &customMessageSysConfigModel{
		defaultMessageSysConfigModel: newMessageSysConfigModel(conn),
		SoftDeletable:                true, //是否启用软删除
	}
}
