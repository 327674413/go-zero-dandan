package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsMainModel = (*customGoodsMainModel)(nil)

type (
	// GoodsMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsMainModel.
	GoodsMainModel interface {
		goodsMainModel
	}

	customGoodsMainModel struct {
		*defaultGoodsMainModel
		softDeletable bool
	}
)

// NewGoodsMainModel returns a model for the database table.
func NewGoodsMainModel(conn sqlx.SqlConn, platId ...string) GoodsMainModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customGoodsMainModel{
		defaultGoodsMainModel: newGoodsMainModel(conn, platid),
		softDeletable:         true, //是否启用软删除
	}
}
