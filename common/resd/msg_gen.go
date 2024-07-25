// Code generated by lang.go. DO NOT EDIT.
package resd

func init() {
	msg = make(map[int]string)
	msg[ErrAccountOrPassWrong] = "ErrAccountOrPassWrong"
	msg[ErrAuth] = "ErrAuth"
	msg[ErrAuthOperateState] = "ErrAuthOperateState"
	msg[ErrAuthOperateUser] = "ErrAuthOperateUser"
	msg[ErrAuthPlat] = "ErrAuthPlat"
	msg[ErrAuthUserNotLogin] = "ErrAuthUserNotLogin"
	msg[ErrConfigNotInit1] = "ErrConfigNotInit1"
	msg[ErrCopier] = "ErrCopier"
	msg[ErrCreateConversation] = "ErrCreateConversation"
	msg[ErrDataBiz] = "ErrDataBiz"
	msg[ErrDataExist1] = "ErrDataExist1"
	msg[ErrJsonDecode] = "ErrJsonDecode"
	msg[ErrJsonEncode] = "ErrJsonEncode"
	msg[ErrLimiterErr] = "ErrLimiterErr"
	msg[ErrMergeFileChunkNotFound] = "ErrMergeFileChunkNotFound"
	msg[ErrMongoDelete] = "ErrMongoDelete"
	msg[ErrMongoInsert] = "ErrMongoInsert"
	msg[ErrMongoSelect] = "ErrMongoSelect"
	msg[ErrMongoStrToId] = "ErrMongoStrToId"
	msg[ErrMongoUpdate] = "ErrMongoUpdate"
	msg[ErrMqPush] = "ErrMqPush"
	msg[ErrMultipartUploadFileHashRequired] = "ErrMultipartUploadFileHashRequired"
	msg[ErrMultipartUploadNotComplete] = "ErrMultipartUploadNotComplete"
	msg[ErrMysql] = "ErrMysql"
	msg[ErrMysqlCommit] = "ErrMysqlCommit"
	msg[ErrMysqlDelete] = "ErrMysqlDelete"
	msg[ErrMysqlInsert] = "ErrMysqlInsert"
	msg[ErrMysqlPrepareUpdate] = "ErrMysqlPrepareUpdate"
	msg[ErrMysqlRollback] = "ErrMysqlRollback"
	msg[ErrMysqlSave] = "ErrMysqlSave"
	msg[ErrMysqlSelect] = "ErrMysqlSelect"
	msg[ErrMysqlStartTrans] = "ErrMysqlStartTrans"
	msg[ErrMysqlUpdate] = "ErrMysqlUpdate"
	msg[ErrNotFound1] = "ErrNotFound1"
	msg[ErrNotFoundUser] = "ErrNotFoundUser"
	msg[ErrNotSupportFileType] = "ErrNotSupportFileType"
	msg[ErrNotSupportImageType] = "ErrNotSupportImageType"
	msg[ErrNotSupportPhoneArea] = "ErrNotSupportPhoneArea"
	msg[ErrPlatClas] = "ErrPlatClas"
	msg[ErrPlatId] = "ErrPlatId"
	msg[ErrPlatInvalid] = "ErrPlatInvalid"
	msg[ErrRedis] = "ErrRedis"
	msg[ErrRedisDec] = "ErrRedisDec"
	msg[ErrRedisGet] = "ErrRedisGet"
	msg[ErrRedisGetUserToken] = "ErrRedisGetUserToken"
	msg[ErrRedisInc] = "ErrRedisInc"
	msg[ErrRedisKeyNil] = "ErrRedisKeyNil"
	msg[ErrRedisSet] = "ErrRedisSet"
	msg[ErrRedisSetUserLoginState] = "ErrRedisSetUserLoginState"
	msg[ErrRedisSetVerifyCode] = "ErrRedisSetVerifyCode"
	msg[ErrRedisSetVerifyCodeInterval] = "ErrRedisSetVerifyCodeInterval"
	msg[ErrReqFieldEmpty1] = "ErrReqFieldEmpty1"
	msg[ErrReqFieldRequired1] = "ErrReqFieldRequired1"
	msg[ErrReqGetPhoneVerifyCodeDayLimit] = "ErrReqGetPhoneVerifyCodeDayLimit"
	msg[ErrReqGetPhoneVerifyCodeHourLimit] = "ErrReqGetPhoneVerifyCodeHourLimit"
	msg[ErrReqGetPhoneVerifyCodeWait] = "ErrReqGetPhoneVerifyCodeWait"
	msg[ErrReqKeyRequired] = "ErrReqKeyRequired"
	msg[ErrReqParam] = "ErrReqParam"
	msg[ErrReqParamFormat1] = "ErrReqParamFormat1"
	msg[ErrReqPhone] = "ErrReqPhone"
	msg[ErrReqRateLimit] = "ErrReqRateLimit"
	msg[ErrReqWait] = "ErrReqWait"
	msg[ErrRpcMissMeta] = "ErrRpcMissMeta"
	msg[ErrRpcResDecode] = "ErrRpcResDecode"
	msg[ErrRpcService] = "ErrRpcService"
	msg[ErrSocialAlreadyBlackMe] = "ErrSocialAlreadyBlackMe"
	msg[ErrSocialAlreadyFriend] = "ErrSocialAlreadyFriend"
	msg[ErrSocialNotAddSelf] = "ErrSocialNotAddSelf"
	msg[ErrSys] = "ErrSys"
	msg[ErrTrdSmsSend] = "ErrTrdSmsSend"
	msg[ErrUploadFileFail] = "ErrUploadFileFail"
	msg[ErrUploadFileSizeLimited1] = "ErrUploadFileSizeLimited1"
	msg[ErrUploadFileTypeLimited1] = "ErrUploadFileTypeLimited1"
	msg[ErrUploadImageSizeLimited1] = "ErrUploadImageSizeLimited1"
	msg[ErrUploadImageTypeLimited1] = "ErrUploadImageTypeLimited1"
	msg[ErrUserMainInfo] = "ErrUserMainInfo"
	msg[ErrVerifyCodeExpired] = "ErrVerifyCodeExpired"
	msg[ErrVerifyCodeWrong] = "ErrVerifyCodeWrong"
	msg[Ok] = "Ok"
	msg[OkAsync] = "OkAsync"
}

const (
	ErrAccountOrPassWrong              = 62000001
	ErrAuth                            = 400
	ErrAuthOperateState                = 40302
	ErrAuthOperateUser                 = 40301
	ErrAuthPlat                        = 40101
	ErrAuthUserNotLogin                = 40201
	ErrConfigNotInit1                  = 603
	ErrCopier                          = 521
	ErrCreateConversation              = 63000004
	ErrDataBiz                         = 660
	ErrDataExist1                      = 607
	ErrJsonDecode                      = 52002
	ErrJsonEncode                      = 52001
	ErrLimiterErr                      = 40502
	ErrMergeFileChunkNotFound          = 60901
	ErrMongoDelete                     = 50302
	ErrMongoInsert                     = 50301
	ErrMongoSelect                     = 50304
	ErrMongoStrToId                    = 50305
	ErrMongoUpdate                     = 50302
	ErrMqPush                          = 504
	ErrMultipartUploadFileHashRequired = 62000030
	ErrMultipartUploadNotComplete      = 62000031
	ErrMysql                           = 501
	ErrMysqlCommit                     = 50107
	ErrMysqlDelete                     = 50102
	ErrMysqlInsert                     = 50101
	ErrMysqlPrepareUpdate              = 50109
	ErrMysqlRollback                   = 50108
	ErrMysqlSave                       = 50104
	ErrMysqlSelect                     = 50105
	ErrMysqlStartTrans                 = 50106
	ErrMysqlUpdate                     = 50103
	ErrNotFound1                       = 606
	ErrNotFoundUser                    = 60600001
	ErrNotSupportFileType              = 60800003
	ErrNotSupportImageType             = 60800002
	ErrNotSupportPhoneArea             = 60800001
	ErrPlatClas                        = 60400002
	ErrPlatId                          = 60400001
	ErrPlatInvalid                     = 604
	ErrRedis                           = 502
	ErrRedisDec                        = 50204
	ErrRedisGet                        = 50202
	ErrRedisGetUserToken               = 502020001
	ErrRedisInc                        = 50203
	ErrRedisKeyNil                     = 50205
	ErrRedisSet                        = 50201
	ErrRedisSetUserLoginState          = 502010003
	ErrRedisSetVerifyCode              = 502010001
	ErrRedisSetVerifyCodeInterval      = 502010002
	ErrReqFieldEmpty1                  = 60101
	ErrReqFieldRequired1               = 601
	ErrReqGetPhoneVerifyCodeDayLimit   = 60500002
	ErrReqGetPhoneVerifyCodeHourLimit  = 60500003
	ErrReqGetPhoneVerifyCodeWait       = 60500001
	ErrReqKeyRequired                  = 600
	ErrReqParam                        = 602
	ErrReqParamFormat1                 = 60201
	ErrReqPhone                        = 60200001
	ErrReqRateLimit                    = 40501
	ErrReqWait                         = 605
	ErrRpcMissMeta                     = 51001
	ErrRpcResDecode                    = 510
	ErrRpcService                      = 50002
	ErrSocialAlreadyBlackMe            = 63000002
	ErrSocialAlreadyFriend             = 63000001
	ErrSocialNotAddSelf                = 63000003
	ErrSys                             = 500
	ErrTrdSmsSend                      = 700
	ErrUploadFileFail                  = 62000010
	ErrUploadFileSizeLimited1          = 62000005
	ErrUploadFileTypeLimited1          = 60800005
	ErrUploadImageSizeLimited1         = 62000004
	ErrUploadImageTypeLimited1         = 60800004
	ErrUserMainInfo                    = 50001
	ErrVerifyCodeExpired               = 62000003
	ErrVerifyCodeWrong                 = 62000002
	Ok                                 = 200
	OkAsync                            = 201
)
const (
	VarAccount          = "VarAccount"
	VarData             = "VarData"
	VarId               = "VarId"
	VarImage            = "VarImage"
	VarMessageSysConfig = "VarMessageSysConfig"
	VarPassword         = "VarPassword"
	VarPhoneNumber      = "VarPhoneNumber"
	VarSmsTemp          = "VarSmsTemp"
	VarStateEm          = "VarStateEm"
	VarUpoladTask       = "VarUpoladTask"
	VarUrl              = "VarUrl"
)
