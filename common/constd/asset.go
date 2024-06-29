package constd

const (
	AssetModeLocal  = 1
	AssetModeMinio  = 2
	AssetModeAliOss = 3
	AssetModeTxCos  = 4

	AssetStateEmCreate  = 0 //提交上传
	AssetStateEmProcess = 1 //上传进行中，未完
	AssetStateEmFinish  = 2 //上传完成
)
