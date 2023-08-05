package resd

// 2成功类，5异常类不适合展示给用户，6校验类，7第三方对接类
const (
	Ok                              = 200   //成功，直接完成
	OkAsync                         = 201   //成功，但属于异步交易
	Auth                            = 400   //权限异常
	AuthPlatErr                     = 40101 // 无效应用
	AuthUserNotLogin                = 40201 // 用户未登录
	SysErr                          = 500   //系统异常
	UserMainInfoErr                 = 50001 //系统异常
	MysqlErr                        = 501   //mysql异常
	MysqlInsertErr                  = 50101
	MysqlDeleteErr                  = 50102
	MysqlUpdateErr                  = 50103
	MysqlSelectErr                  = 50104
	MysqlStartTransErr              = 50106
	MysqlCommitErr                  = 50107
	MysqlRollbackErr                = 50108
	RedisErr                        = 502 //redis异常
	RedisSetErr                     = 50201
	RedisGetErr                     = 50202
	RedisIncErr                     = 50203
	RedisDecErr                     = 50204
	RedisSetVerifyCodeErr           = 502010001
	RedisSetVerifyCodeIntervalErr   = 502010002
	RedisSetUserLoginStateErr       = 502010003
	RedisGetUserTokenErr            = 502020001
	RpcResDecodeErr                 = 510
	ReqKeyRequired                  = 600 //未提供主键
	ReqFieldRequired                = 601 //未提供比必填字段
	ReqParamErr                     = 602 //请求参数不正确
	ConfigNotInit1                  = 603 //未配置参数
	PlatInvalid                     = 604 //无效应用
	PlatIdErr                       = 60400001
	PlatClasErr                     = 60400002
	ReqWait                         = 605      //请求太频繁
	NotFound                        = 606      //信息不存在
	DataExist1                      = 607      //数据已存在
	NotSupportPhoneArea             = 60800001 //暂不支持手机号
	NotSupportImageType             = 60800002 //图片格式不支持
	NotSupportFileType              = 60800003
	UploadImageTypeLimited1         = 60800004
	UploadFileTypeLimited1          = 60800005
	MergeFileChunkNotFound          = 60901    // 未找到合并的文件分片
	AccountOrPassWrong              = 62000001 //登录校验失败
	VerifyCodeWrong                 = 62000002
	VerifyCodeExpired               = 62000003
	UploadImageSizeLimited1         = 62000004 //图片超出大小
	UploadFileSizeLimited1          = 62000005 //图片超出大小
	UploadFileFail                  = 62000010 //图片超出大小
	MultipartUploadFileHashRequired = 62000030
	MultipartUploadNotComplete      = 62000031
	NotFoundUser                    = 60600001 //用户不存在
	ReqPhoneErr                     = 60200001
	ReqGetPhoneVerifyCodeWait       = 60500001 //请求太频繁
	ReqGetPhoneVerifyCodeDayLimit   = 60500002
	ReqGetPhoneVerifyCodeHourLimit  = 60500003

	TrdSmsSendErr = 700
)
