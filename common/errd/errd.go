package errd

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/common/land"
)

type errd struct {
	Result bool   `json:"result"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (t *errd) Error() string {
	return fmt.Sprintf("result:%v, code: %d, msg: %s", t.Result, t.Code, t.Msg)
}
func FailCode(localize *i18n.Localizer, errCode int, tempData ...map[string]string) error {
	msg := ErrMsg(localize, errCode, tempData...)
	return &errd{Result: false, Code: errCode, Msg: msg}
}
func Fail(msg string, code ...int) error {
	apiCode := 400
	if len(code) > 0 {
		apiCode = code[0]
	}
	return &errd{Result: false, Code: apiCode, Msg: msg}
}
func FailI18n(localize *i18n.Localizer, temp string, tempData ...map[string]string) error {
	apiCode := 400
	return &errd{Result: false, Code: apiCode, Msg: land.Trans(localize, temp, tempData...)}
}
func FailI18nWithCode(localize *i18n.Localizer, temp string, code int, tempData ...map[string]string) error {
	return &errd{Result: false, Code: code, Msg: land.Trans(localize, temp, tempData...)}
}
