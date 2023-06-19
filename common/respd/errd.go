package respd

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/common/land"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Errd struct {
	Result bool   `json:"result"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}
type Okd struct {
	Result bool `json:"result"`
	Code   int  `json:"code"`
	Data   any  `json:"data"`
}

func (t *Errd) Error() string {
	return fmt.Sprintf("result:%v, code: %d, msg: %s", t.Result, t.Code, t.Msg)
}
func RpcDecodeErr(rpcError error) (int, string) {
	if r, ok := status.FromError(rpcError); ok { // grpc err错误
		return int(r.Code()), r.Message()
	} else {
		return -1, "rpc err decode error"
	}
}
func RpcEncodeMsgErr(msg string, errCode ...int) error {
	code := 400
	if len(errCode) > 0 {
		code = errCode[0]
	}
	return status.Error(codes.Code(code), "msg:"+msg)
}
func RpcEncodeSysErr(msg string, errCode ...int) error {
	code := 500
	if len(errCode) > 0 {
		code = errCode[0]
	}
	return status.Error(codes.Code(code), "msg:"+msg)
}
func RpcEncodeTempErr(errCode int, tempData []string) error {
	msg, _ := json.Marshal(tempData)
	return status.Error(codes.Code(errCode), "tmp:"+string(msg))
}
func RpcFail(localize *i18n.Localizer, rpcError error) error {
	errCode, msg := RpcDecodeErr(rpcError)
	msgType := msg[:4]
	if msgType == "msg:" {
		return Fail(msg[4:], errCode)
	} else if msgType == "tmp:" {
		var tempData []string
		err := json.Unmarshal([]byte(msg[4:]), &tempData)
		if err == nil {
			return FailCode(localize, errCode, tempData)
		}

	}
	return Fail(msg, errCode)
}
func FailCode(localize *i18n.Localizer, errCode int, tempData ...[]string) error {
	var tempD []string
	if len(tempData) > 0 {
		tempD = tempData[0]
	}
	msg := Msg(localize, errCode, tempD)
	return &Errd{Result: false, Code: errCode, Msg: msg}
}
func Fail(msg string, code ...int) error {
	apiCode := 400
	if len(code) > 0 {
		apiCode = code[0]
	}
	return &Errd{Result: false, Code: apiCode, Msg: msg}
}
func Succ(data any) *Okd {
	return &Okd{Result: true, Code: Ok, Data: data}
}
func SuccAsync(data any) *Okd {
	return &Okd{Result: true, Code: OkAsync, Data: data}
}
func FailI18n(localize *i18n.Localizer, temp string, tempData ...map[string]string) error {
	apiCode := 400
	return &Errd{Result: false, Code: apiCode, Msg: land.Trans(localize, temp, tempData...)}
}
func FailI18nWithCode(localize *i18n.Localizer, temp string, code int, tempData ...map[string]string) error {
	return &Errd{Result: false, Code: code, Msg: land.Trans(localize, temp, tempData...)}
}
