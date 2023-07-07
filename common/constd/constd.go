package constd

const (
	ModeDev = "dev" //开发环境
	ModePro = "pro" //生产环境

)
const (
	PhoneAreaEmChina = "86"
)
const (
	PlatClasEmMall = 1
)
const (
	UserStateEmNormal = 20 //正常
)
const (
	SexEmUnknow = 0 //未知
	SexEmMan    = 1 //男
	SexEmWoman  = 2 //女
)

const (
	RedisKeyUserToken = "userToken"
)

const (
	AssetModeLocal  = 1
	AssetModeMinio  = 2
	AssetModeOssAli = 3
	AssetModeOssTx  = 4

	AssetStateEmCreate  = 0 //提交上传
	AssetStateEmProcess = 1 //上传进行中，未完
	AssetStateEmFinish  = 2 //上传完成
)
