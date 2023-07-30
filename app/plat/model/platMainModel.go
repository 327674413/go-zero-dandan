package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PlatMainModel = (*customPlatMainModel)(nil)

type (
	// PlatMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatMainModel.
	PlatMainModel interface {
		platMainModel
	}

	customPlatMainModel struct {
		*defaultPlatMainModel
		softDeletable bool
	}
)

// NewPlatMainModel returns a model for the database table.
func NewPlatMainModel(conn sqlx.SqlConn, platId ...int64) PlatMainModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customPlatMainModel{
		defaultPlatMainModel: newPlatMainModel(conn, platid),
		softDeletable:        true, //是否启用软删除
	}
}
