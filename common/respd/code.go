package respd

// 2成功类，5异常类不适合展示给用户，6校验类，7第三方对接类
const (
	Ok                            = 200 //成功，直接完成
	OkAsync                       = 201 //成功，但属于异步交易
	Auth                          = 400 //权限异常
	Err                           = 500 //系统异常
	MysqlErr                      = 501 //mysql异常
	RedisErr                      = 502 //redis异常
	RedisSetErr                   = 50201
	RedisGetErr                   = 50202
	RedisIncErr                   = 50203
	RedisDecErr                   = 50204
	RedisSetVerifyCodeErr         = 502010001
	RedisSetVerifyCodeIntervalErr = 502010002
	ReqKeyRequired                = 600 //未提供主键
	ReqFieldRequired              = 601 //未提供比必填字段
	PlatConfigNotInit             = 603 //未设置配置
	PlatInvalid                   = 604 //无效appid、secret
	ReqPhoneErr                   = 60200001

	TrdSmsSendErr = 700
)
