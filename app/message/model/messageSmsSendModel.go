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
		softDeletable bool
	}
)

// NewMessageSmsSendModel returns a model for the database table.
func NewMessageSmsSendModel(conn sqlx.SqlConn, platId ...string) MessageSmsSendModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customMessageSmsSendModel{
		defaultMessageSmsSendModel: newMessageSmsSendModel(conn, platid),
		softDeletable:              true, //是否启用软删除
	}
}
