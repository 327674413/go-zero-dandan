package respd

// 2成功类，5异常类不适合展示给用户，6校验类，7第三方对接类
const (
	Ok                = 200 //成功，直接完成
	OkAsync           = 201 //成功，但属于异步交易
	Auth              = 400 //权限异常
	Err               = 500 //系统异常
	ReqKeyRequired    = 600 //未提供主键
	ReqFieldRequired  = 601 //未提供比必填字段
	PlatConfigNotInit = 603 //未设置配置
	ReqPhoneError     = 60200001

	TrdSmsSendError = 700
)
