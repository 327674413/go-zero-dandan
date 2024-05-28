package resd

var Msg map[int]string

func init() {
	Msg = make(map[int]string)
	Msg[Auth] = "Auth"
	Msg[AuthUserNotLogin] = "AuthUserNotLogin"
	Msg[AuthPlatErr] = "AuthPlatErr"
	Msg[AccountOrPassWrong] = "AccountOrPassWrong"
	Msg[ConfigNotInit1] = "ConfigNotInit"
	Msg[DataExist1] = "DataExist1"
	Msg[SysErr] = "SysErr"
	Msg[UserMainInfoErr] = "SysErr"
	Msg[UploadImageSizeLimited1] = "UploadImageSizeLimited1"
	Msg[UploadFileSizeLimited1] = "UploadFileSizeLimited1"
	Msg[MergeFileChunkNotFound] = "MergeFileChunkNotFound"
	Msg[NotFound] = "NotFound"
	Msg[NotFoundUser] = "NotFoundUser"
	Msg[NotSupportPhoneArea] = "NotSupportPhoneArea"
	Msg[NotSupportFileType] = "NotSupportFileType"
	Msg[NotSupportImageType] = "NotSupportImageType"
	Msg[MysqlErr] = "MysqlErr"
	Msg[MysqlInsertErr] = "MysqlInsertErr"
	Msg[MysqlDeleteErr] = "MysqlDeleteErr"
	Msg[MysqlUpdateErr] = "MysqlUpdateErr"
	Msg[MysqlSelectErr] = "MysqlSelectErr"
	Msg[MysqlStartTransErr] = "MysqlStartTransErr"
	Msg[MysqlCommitErr] = "MysqlCommitErr"
	Msg[MysqlRollbackErr] = "MysqlRollbackErr"
	Msg[MultipartUploadNotComplete] = "MultipartUploadNotComplete"
	Msg[MultipartUploadFileHashRequired] = "MultipartUploadFileHashRequired"
	Msg[Ok] = "Success"
	Msg[OkAsync] = "SuccessAsync"
	Msg[PlatClasErr] = "PlatClasErr"
	Msg[PlatIdErr] = "PlatIdErr"
	Msg[PlatInvalid] = "PlatInvalid"
	Msg[RedisErr] = "RedisErr"
	Msg[RedisIncErr] = "RedisIncErr"
	Msg[RedisDecErr] = "RedisDecErr"
	Msg[RedisSetErr] = "RedisSetErr"
	Msg[RedisSetVerifyCodeErr] = "RedisSetErr"
	Msg[RedisSetVerifyCodeIntervalErr] = "RedisSetErr"
	Msg[RedisSetUserLoginStateErr] = "RedisSetErr"
	Msg[RedisGetErr] = "RedisGetErr"
	Msg[RedisGetUserTokenErr] = "RedisGetErr"
	Msg[ReqFieldRequired1] = "ReqFieldRequired1"
	Msg[ReqGetPhoneVerifyCodeWait] = "ReqGetPhoneVerifyCodeWait"
	Msg[ReqGetPhoneVerifyCodeDayLimit] = "ReqGetPhoneVerifyCodeDayLimit"
	Msg[ReqGetPhoneVerifyCodeHourLimit] = "ReqGetPhoneVerifyCodeHourLimit"
	Msg[ReqKeyRequired] = "ReqKeyRequired"
	Msg[ReqParamErr] = "ReqParamErr"
	Msg[ReqPhoneErr] = "ReqPhoneErr"
	Msg[ReqWait] = "ReqWait"
	Msg[UploadFileFail] = "UploadFileFail"
	Msg[UploadImageTypeLimited1] = "UploadImageTypeLimited1"
	Msg[UploadFileTypeLimited1] = "UploadFileTypeLimited1"
	Msg[VerifyCodeWrong] = "VerifyCodeWrong"
	Msg[VerifyCodeExpired] = "VerifyCodeExpired"
	Msg[CopierErr] = "CopierErr"

}
