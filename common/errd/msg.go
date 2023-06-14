package errd

import (
	"fmt"
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
	msg[PlatConfigNotInit] = "PlatConfigNotInit"

}

func Msg(localize *i18n.Localizer, errCode int, tempData []string) string {
	m := make(map[string]string)
	for i, v := range tempData {
		key := "Field" + fmt.Sprint(i+1)
		m[key] = getMsg(localize, v)
	}
	if code, ok := msg[errCode]; ok {
		return land.Trans(localize, code, m)
	} else {
		return land.Trans(localize, msg[Err], m)
	}

}

func getMsg(localize *i18n.Localizer, tempCode string, tempData ...map[string]string) string {
	return land.Trans(localize, tempCode, tempData...)
}
