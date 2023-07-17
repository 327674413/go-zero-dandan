package resd

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/common/land"
)

var msg map[int]string

func init() {
	msg = make(map[int]string)
	msg[Auth] = "Auth"
	msg[AccountOrPassWrong] = "AccountOrPassWrong"
	msg[ConfigNotInit] = "ConfigNotInit"
	msg[DataExist1] = "DataExist1"
	msg[Err] = "Error"
	msg[ImageSizeLimited1] = "ImageSizeLimited1"
	msg[NotFound] = "NotFound"
	msg[NotFoundUser] = "NotFoundUser"
	msg[NotSupportPhoneArea] = "NotSupportPhoneArea"
	msg[NotSupportFileType] = "NotSupportFileType"
	msg[NotSupportImageType] = "NotSupportImageType"
	msg[MysqlErr] = "MysqlErr"
	msg[MysqlInsertErr] = "MysqlInsertErr"
	msg[MysqlDeleteErr] = "MysqlDeleteErr"
	msg[MysqlUpdateErr] = "MysqlUpdateErr"
	msg[MysqlSelectErr] = "MysqlSelectErr"
	msg[MysqlStartTransErr] = "MysqlStartTransErr"
	msg[MysqlCommitErr] = "MysqlCommitErr"
	msg[MysqlRollbackErr] = "MysqlRollbackErr"
	msg[Ok] = "Success"
	msg[OkAsync] = "SuccessAsync"
	msg[PlatClasErr] = "PlatClasErr"
	msg[PlatIdErr] = "PlatIdErr"
	msg[PlatInvalid] = "PlatInvalid"
	msg[RedisErr] = "RedisErr"
	msg[RedisIncErr] = "RedisIncErr"
	msg[RedisDecErr] = "RedisDecErr"
	msg[RedisSetErr] = "RedisSetErr"
	msg[RedisSetVerifyCodeErr] = "RedisSetErr"
	msg[RedisSetVerifyCodeIntervalErr] = "RedisSetErr"
	msg[RedisSetUserLoginStateErr] = "RedisSetErr"
	msg[RedisGetErr] = "RedisGetErr"
	msg[RedisGetUserTokenErr] = "RedisGetErr"
	msg[ReqFieldRequired] = "ReqFieldRequired"
	msg[ReqGetPhoneVerifyCodeWait] = "ReqGetPhoneVerifyCodeWait"
	msg[ReqGetPhoneVerifyCodeDayLimit] = "ReqGetPhoneVerifyCodeDayLimit"
	msg[ReqGetPhoneVerifyCodeHourLimit] = "ReqGetPhoneVerifyCodeHourLimit"
	msg[ReqKeyRequired] = "ReqKeyRequired"
	msg[ReqParamErr] = "ReqParamErr"
	msg[ReqPhoneErr] = "ReqPhoneErr"
	msg[ReqWait] = "ReqWait"
	msg[UploadFileFail] = "UploadFileFail"
	msg[VerifyCodeWrong] = "VerifyCodeWrong"
	msg[VerifyCodeExpired] = "VerifyCodeExpired"

}

func Msg(localize *i18n.Localizer, msgCode int, tempDataArr ...[]string) string {
	tempData := make([]string, 0)
	if len(tempDataArr) > 0 {
		tempData = tempDataArr[0]
	}
	m := make(map[string]string)
	for i, v := range tempData {
		key := "Field" + fmt.Sprint(i+1)
		m[key] = getMsg(localize, v)
	}
	if code, ok := msg[msgCode]; ok {
		return land.Trans(localize, code, m)
	} else {
		return land.Trans(localize, msg[Err], m)
	}

}

func getMsg(localize *i18n.Localizer, tempCode string, tempData ...map[string]string) string {
	return land.Trans(localize, tempCode, tempData...)
}
