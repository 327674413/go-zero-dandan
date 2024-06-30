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
		softDeletable bool
	}
)

// NewMessageSysConfigModel returns a model for the database table.
func NewMessageSysConfigModel(conn sqlx.SqlConn, platId ...string) MessageSysConfigModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customMessageSysConfigModel{
		defaultMessageSysConfigModel: newMessageSysConfigModel(conn, platid),
		softDeletable:                true, //是否启用软删除
	}
}
