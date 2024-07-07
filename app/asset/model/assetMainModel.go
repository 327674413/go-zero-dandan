package model

var _ AssetMainModel = (*customAssetMainModel)(nil)
var softDeletableAssetMain = true

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
	// 自定义方法加在customAssetMainModel上
)
