package resd

import (
	"context"
	"fmt"
	"go-zero-dandan/common/fmtd"
	"runtime"
)

type Resp struct {
	ctx  context.Context
	mode string
}

const (
	levelError = "error" //真正的异常报错，需要关注
	levelInfo  = "info"  //正常业务逻辑校验的问题
)
const (
	modeDev  = "dev"
	modeTest = "test"
	modeProd = "prod"
)

// NewResd 创建统一错误返回
func NewResd(ctx context.Context, mode string) *Resp {
	if ctx == nil {
		return &Resp{
			ctx:  context.Background(),
			mode: mode,
		}
	} else {
		return &Resp{
			ctx:  ctx,
			mode: mode,
		}
	}
}

// 错误返回中的错误方法，可以多次嵌套Error，追溯整个来源
func (r *Resp) Error(err error, initErrCode ...int) error {
	errorCode := SysErr
	if len(initErrCode) > 0 {
		errorCode = initErrCode[0]
	}
	skip := 2
	danErr, ok := err.(*danError)
	if ok {
		//是自定义错误
		_, file, line, okk := runtime.Caller(skip - 1)
		if okk {
			danErr.callers = append(danErr.callers, fmt.Sprintf("%s:%d", file, line))
		}
		//如果有传，则覆盖原先的错误码
		if len(initErrCode) > 0 {
			danErr.Code = errorCode
		}
	} else {
		//不是自定义错误，创建
		danErr = newDanErr(err.Error(), errorCode)
	}
	if r.mode != "prod" {
		fmtd.WithCaller(skip).Error(err.Error())
	}
	return danErr
}
func (r *Resp) ErrorWithTemp(err error, errorCode int, temps ...string) error {
	skip := 2
	danErr, ok := err.(*danError)
	if ok {
		//是自定义错误
		_, file, line, okk := runtime.Caller(skip - 1)
		if okk {
			danErr.callers = append(danErr.callers, fmt.Sprintf("%s:%d", file, line))
		}
		//如果有传，则覆盖原先的错误码
		danErr.Code = errorCode
		danErr.SetTemps(temps)
	} else {
		//不是自定义错误，创建
		danErr = newDanErr(err.Error(), errorCode, temps...)
	}
	if r.mode != "prod" {
		fmtd.WithCaller(skip).Error(err.Error())
	}
	return danErr
}
func (r *Resp) NewError(msg string, initErrCode ...int) error {
	errorCode := SysErr
	if len(initErrCode) > 0 {
		errorCode = initErrCode[0]
	}
	skip := 2
	var danErr *danError
	if r.mode != "prod" {
		fmtd.WithCaller(skip).Error(msg)
	}
	danErr = newDanErr(msg, errorCode)
	_, file, line, ok := runtime.Caller(skip - 1)
	if ok {
		danErr.AppendCaller(fmt.Sprintf("%s:%d", file, line))
	}
	return danErr
}
func (r *Resp) NewErrorWithTemp(msg string, errorCode int, temps ...string) error {
	skip := 2
	var danErr *danError
	if r.mode != "prod" {
		fmtd.WithCaller(skip).Error(msg)
	}
	danErr = newDanErr(msg, errorCode, temps...)
	_, file, line, ok := runtime.Caller(skip - 1)
	if ok {
		danErr.AppendCaller(fmt.Sprintf("%s:%d", file, line))
	}
	return danErr
}
