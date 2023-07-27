package resd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type danError struct {
	Result     bool     `json:"result"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
	temps      []string `json:"-"`
	callerSkip int
}

func (t *danError) Error() string {
	return fmt.Sprintf("%s", t.Msg)
}
func (t *danError) SetTemps(temps []string) *danError {
	t.temps = temps
	return t
}
func (t *danError) GetTemps() []string {
	return t.temps
}
func newDanErr(msg string, errCode int, tempStr ...string) *danError {
	res := &danError{
		Result: false,
		Code:   errCode,
		Msg:    msg,
	}
	if len(tempStr) > 0 {
		res.SetTemps(tempStr)
	}
	return res
}

// Error 对error进行封装返回
func Error(err error, errorCode ...int) error {
	skip := 1
	code := SysErr
	if e, ok := err.(*danError); ok {
		e.callerSkip = e.callerSkip + 1
		skip = e.callerSkip + skip
		code = e.Code
	}
	logx.WithCallerSkip(skip).Error(err)
	if len(errorCode) > 0 {
		code = errorCode[0]
	}
	return newDanErr(err.Error(), code)

}

// ErrorCtx 对error进行封装返回,带上下文
func ErrorCtx(ctx context.Context, err error, errorCode ...int) error {
	skip := 1
	code := SysErr
	if e, ok := err.(*danError); ok {
		e.callerSkip = e.callerSkip + 1
		skip = e.callerSkip + skip
		code = e.Code
	}
	logx.WithCallerSkip(skip).WithContext(ctx).Error(ctx, msg)
	if len(errorCode) > 0 {
		code = errorCode[0]
	}
	return newDanErr(err.Error(), code)
}

// ErrorWithTemp 对error进行封装返回，附带模版变量
func ErrorWithTemp(err error, errorCode int, temps ...string) error {
	skip := 1
	if e, ok := err.(*danError); ok {
		e.callerSkip = e.callerSkip + 1
		skip = e.callerSkip + skip
	}
	logx.WithCallerSkip(skip).Error(err)
	return newDanErr(err.Error(), errorCode, temps...)
}

// ErrorWithTempCtx 对error进行封装返回，附带模版变量，带上下文
func ErrorWithTempCtx(ctx context.Context, err error, errorCode int, temps ...string) error {
	skip := 1
	if e, ok := err.(*danError); ok {
		e.callerSkip = e.callerSkip + 1
		skip = e.callerSkip + skip
	}
	logx.WithCallerSkip(skip).WithContext(ctx).Error(ctx, msg)
	return newDanErr(err.Error(), errorCode, temps...)
}

// NewErr 创建新的error
func NewErr(msg string, errorCode ...int) error {
	logx.WithCallerSkip(1).Error(errors.New(msg))
	if len(errorCode) > 0 {
		return newDanErr(msg, errorCode[0])
	}
	return newDanErr(msg, SysErr)
}

// NewErrCtx 创建新的error，带上下文
func NewErrCtx(ctx context.Context, msg string, errorCode ...int) error {
	code := SysErr
	logx.WithCallerSkip(1).WithContext(ctx).Error(errors.New(msg))
	if len(errorCode) > 0 {
		code = errorCode[0]
	}
	return newDanErr(msg, code)
}

// NewErrWithTemp 创建新的error，带模版
func NewErrWithTemp(msg string, errorCode int, temps ...string) error {
	logx.WithCallerSkip(1).Error(errors.New(msg))
	return newDanErr(msg, errorCode, temps...)
}

// NewErrWithTempCtx 创建新的error，带模版，带上下文
func NewErrWithTempCtx(ctx context.Context, msg string, errorCode int, temps ...string) error {
	logx.WithCallerSkip(1).WithContext(ctx).Error(errors.New(msg))
	return newDanErr(msg, errorCode, temps...)
}
func RpcDecodeErr(rpcError error) (int, string) {
	if r, ok := status.FromError(rpcError); ok { // grpc err错误
		return int(r.Code()), r.Message()
	} else {
		return -1, "rpc err decode error"
	}
}
func RpcEncodeMsgErr(text string, errCode ...int) error {
	code := 400
	if len(errCode) > 0 {
		code = errCode[0]
	}
	return status.Error(codes.Code(code), "msg:"+text)
}
func RpcEncodeSysErr(text string, errCode ...int) error {
	code := 500
	if len(errCode) > 0 {
		code = errCode[0]
	}
	return status.Error(codes.Code(code), "msg:"+text)
}
func RpcEncodeTempErr(errCode int, tempData ...[]string) error {
	var tempD []string
	if len(tempData) > 0 {
		tempD = tempData[0]
	}
	text, _ := json.Marshal(tempD)
	return status.Error(codes.Code(errCode), "tmp:"+string(text))
}
func RpcFail(localize *i18n.Localizer, rpcError error) error {
	errCode, text := RpcDecodeErr(rpcError)
	msgType := text[:4]
	if msgType == "msg:" {
		return newDanErr(text[4:], errCode)
	} else if msgType == "tmp:" {
		var tempData []string
		err := json.Unmarshal([]byte(text[4:]), &tempData)
		if err == nil {
			return ApiFail(localize, rpcError)
		}

	}
	return newDanErr(text, errCode)
}

// ApiFail 按错误内容返回错误信息
func ApiFail(localize *i18n.Localizer, failErr error) error {
	if err, ok := failErr.(*danError); ok {
		err.Msg = Msg(localize, err.Code, err.GetTemps())
		return err
	} else {
		return failErr
	}

}
