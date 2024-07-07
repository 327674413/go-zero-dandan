package model

var _ AssetNetdiskFileModel = (*customAssetNetdiskFileModel)(nil)
var softDeletableAssetNetdiskFile = true

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
	// 自定义方法加在customAssetNetdiskFileModel上
)
