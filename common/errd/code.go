package errd

// 2成功类，5一场类，
const (
	Ok               = 200 //成功，直接完成
	OkAsync          = 201 //成功，但属于异步交易
	Auth             = 400 //权限异常
	Err              = 500 //系统异常
	ReqKeyRequired   = 600 //未提供主键
	ReqFieldRequired = 601 //未提供比必填字段

	ReqPhoneError = 60200001
)
