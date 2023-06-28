package resd

// 2成功类，5异常类不适合展示给用户，6校验类，7第三方对接类
const (
	Ok                             = 200 //成功，直接完成
	OkAsync                        = 201 //成功，但属于异步交易
	Auth                           = 400 //权限异常
	Err                            = 500 //系统异常
	MysqlErr                       = 501 //mysql异常
	MysqlInsertErr                 = 50101
	MysqlDeleteErr                 = 50102
	MysqlUpdateErr                 = 50103
	MysqlSelectErr                 = 50104
	MysqlStartTransErr             = 50105
	MysqlCommitErr                 = 50106
	MysqlRollbackErr               = 50107
	RedisErr                       = 502 //redis异常
	RedisSetErr                    = 50201
	RedisGetErr                    = 50202
	RedisIncErr                    = 50203
	RedisDecErr                    = 50204
	RedisSetVerifyCodeErr          = 502010001
	RedisSetVerifyCodeIntervalErr  = 502010002
	ReqKeyRequired                 = 600 //未提供主键
	ReqFieldRequired               = 601 //未提供比必填字段
	ReqParamErr                    = 602 //请求参数不正确
	ConfigNotInit                  = 603 //未配置参数
	PlatInvalid                    = 604 //无效应用
	PlatIdErr                      = 60400001
	PlatClasErr                    = 60400002
	ReqWait                        = 605      //请求太频繁
	NotFound                       = 606      //信息不存在
	NotSupportPhoneArea            = 60700001 //暂不支持手机号
	AccountOrPassWrong             = 62000001 //登录校验失败
	VerifyCodeWrong                = 62000002
	VerifyCodeExpired              = 62000003
	NotFoundUser                   = 60600001 //用户不存在
	ReqPhoneErr                    = 60200001
	ReqGetPhoneVerifyCodeWait      = 60500001 //请求太频繁
	ReqGetPhoneVerifyCodeDayLimit  = 60500002
	ReqGetPhoneVerifyCodeHourLimit = 60500003

	TrdSmsSendErr = 700
)
