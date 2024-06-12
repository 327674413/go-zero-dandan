package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AssetNetdiskFileModel = (*customAssetNetdiskFileModel)(nil)

type (
	// AssetNetdiskFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssetNetdiskFileModel.
	AssetNetdiskFileModel interface {
		assetNetdiskFileModel
	}

	customAssetNetdiskFileModel struct {
		*defaultAssetNetdiskFileModel
		softDeletable bool
	}
)

// NewAssetNetdiskFileModel returns a model for the database table.
func NewAssetNetdiskFileModel(conn sqlx.SqlConn, platId ...string) AssetNetdiskFileModel {
	var platid string
	if len(platId) > 0 {
		platid = platId[0]
	} else {
		platid = ""
	}
	return &customAssetNetdiskFileModel{
		defaultAssetNetdiskFileModel: newAssetNetdiskFileModel(conn, platid),
		softDeletable:                true, //是否启用软删除
	}
}
