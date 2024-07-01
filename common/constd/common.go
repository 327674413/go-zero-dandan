package constd

const (
	ModeDev = "dev" //开发环境
	ModePro = "pro" //生产环境

)
const (
	ResultFinish  = 0  //成功，且同步完成
	ResultTasking = 1  //成功，提交任务，不确定是否完成
	ResultFail    = -1 //失败
)
const (
	SysRoleEmUser   = "user"
	SysRoleEmAdmin  = "admin"
	SysRoleEmSystem = "system"
)
const (
	PhoneAreaEmChina = "86"
)
const (
	PlatClasEmMall = 1
)

const (
	SexEmUnknow = 0 //未知
	SexEmMan    = 1 //男
	SexEmWoman  = 2 //女
)
