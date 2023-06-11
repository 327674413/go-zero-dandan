package errd

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/common/land"
)

var msg map[int]string

func init() {
	msg = make(map[int]string)
	msg[Ok] = "Success"
	msg[OkAsync] = "SuccessAsync"
	msg[Auth] = "Auth"
	msg[Err] = "Error"
	msg[ReqFieldRequired] = "ReqFieldRequired"
	msg[ReqPhoneError] = "ReqPhoneError"
}

func ErrMsg(localize *i18n.Localizer, errCode int, tempData ...map[string]string) string {
	if code, ok := msg[errCode]; ok {
		return land.Trans(localize, code, tempData...)
	} else {
		return land.Trans(localize, msg[Err], tempData...)
	}
}
