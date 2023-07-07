package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AssetMainModel = (*customAssetMainModel)(nil)

type (
	// AssetMainModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssetMainModel.
	AssetMainModel interface {
		assetMainModel
	}

	customAssetMainModel struct {
		*defaultAssetMainModel
		softDeletable bool
	}
)

// NewAssetMainModel returns a model for the database table.
func NewAssetMainModel(conn sqlx.SqlConn, platId ...int64) AssetMainModel {
	var platid int64
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = 0
	}
	return &customAssetMainModel{
		defaultAssetMainModel: newAssetMainModel(conn, platid),
		softDeletable:         true, //是否启用软删除
	}
}
