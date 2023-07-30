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
		softDeletable bool
	}
)

// NewMessageSmsTempModel returns a model for the database table.
func NewMessageSmsTempModel(conn sqlx.SqlConn, platId ...int64) MessageSmsTempModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customMessageSmsTempModel{
		defaultMessageSmsTempModel: newMessageSmsTempModel(conn, platid),
		softDeletable:              true, //是否启用软删除
	}
}
