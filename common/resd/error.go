package resd

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
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
		//skip = e.callerSkip + skip //目前打算每层都调用，所以不哦那个增加了
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
	if r, ok := status.FromError(err); ok { // grpc err错误
		fmt.Println("1", r.Code())
		fmt.Println("2", r.Message())
		fmt.Println("3", r.Err())
		fmt.Println("4", r.String())
	} else {
		fmt.Println("不是rpc")
	}
	skip := 1
	code := SysErr
	if e, ok := err.(*danError); ok {
		e.callerSkip = e.callerSkip + 1
		//skip = e.callerSkip + skip
		code = e.Code
	}
	logx.WithCallerSkip(skip).WithContext(ctx).Error(err)
	if len(errorCode) > 0 {
		code = errorCode[0]
	}
	return newDanErr(err.Error(), code)
}

// ErrorWithTemp 对error进行封装返回，附带模版变量
func ErrorWithTemp(err error, errorCode int, temps ...string) error {
	skip := 1
	if e, ok := err.(*danError); ok {
		//e.callerSkip = e.callerSkip + 1
		skip = e.callerSkip + skip
	}
	logx.WithCallerSkip(skip).Error(err)
	return newDanErr(err.Error(), errorCode, temps...)
}

// ErrorWithTempCtx 对error进行封装返回，附带模版变量，带上下文
func ErrorWithTempCtx(ctx context.Context, err error, errorCode int, temps ...string) error {
	skip := 1
	if e, ok := err.(*danError); ok {
		//e.callerSkip = e.callerSkip + 1
		skip = e.callerSkip + skip
	}
	logx.WithCallerSkip(skip).WithContext(ctx).Error(err)
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

// RpcErrDecode 解码
func RpcErrDecode(rpcError error) error {
	if r, ok := status.FromError(rpcError); ok { // grpc err错误
		res := strings.Split(r.Message(), " ,,,,,, ")
		if len(res) == 0 {
			return NewErr("rpc错误内容为空", RpcResDecodeErr)
		}
		msg := res[0]
		res = res[1:]
		return NewErrWithTemp(msg, int(r.Code()), res...)
	} else {
		return NewErr("rpc错误解码失败", RpcResDecodeErr)
	}
}

// RpcErrEncode rpc结果报错时用此方法返回，配合RpcError解析，新错误 以及 多模版都通过新建NewErr后再传入
func RpcErrEncode(err error) error {
	if err == nil {
		return nil
	}
	code := SysErr
	res := make([]string, 0)
	if e, ok := err.(*danError); ok {
		e.callerSkip = e.callerSkip + 1
		//skip = e.callerSkip + skip //目前打算每层都调用，所以不哦那个增加了
		code = e.Code
		res = append(res, e.Msg)
		res = append(res, e.temps...)
	} else {
		res = append(res, err.Error())
	}
	return status.Error(codes.Code(code), strings.Join(res, " ,,,,,, "))
}

func AssertErr(failErr error) (*danError, bool) {
	if err, ok := failErr.(*danError); ok {
		return err, ok
	} else {
		return nil, false
	}
}

/*
// ApiFail 按错误内容返回错误信息
func ApiFail(localize *i18n.Localizer, failErr error) error {
	if err, ok := failErr.(*danError); ok {
		err.Msg = Msg(localize, err.Code, err.GetTemps())
		return err
	} else {
		return failErr
	}

}
*/
