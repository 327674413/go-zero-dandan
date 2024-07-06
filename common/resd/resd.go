package resd

import (
	"context"
	"fmt"
	"go-zero-dandan/common/fmtd"
	"runtime"
)

var Mode string

type Resp struct {
	ctx  context.Context
	mode string
	*Transfer
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
func NewResd(ctxOrNil context.Context, langTransfer *Transfer) *Resp {
	if ctxOrNil == nil {
		return &Resp{
			ctx:      context.Background(),
			Transfer: langTransfer,
		}
	} else {
		return &Resp{
			ctx:      ctxOrNil,
			Transfer: langTransfer,
		}
	}
}

// 错误返回中的错误方法，可以多次嵌套Error，追溯整个来源
func (t *Resp) Error(err error, initErrCode ...int) error {
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
	t.fmtdError(skip, errorCode, danErr)
	return danErr
}
func (t *Resp) ErrorWithTemp(err error, errorCode int, temps ...string) error {
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
	t.fmtdError(skip, errorCode, danErr)
	return danErr
}
func (t *Resp) NewErr(initErrCode ...int) error {
	errorCode := SysErr
	if len(initErrCode) > 0 {
		errorCode = initErrCode[0]
	}
	skip := 2

	var danErr *danError
	danErr = newDanErr("", errorCode)
	_, file, line, ok := runtime.Caller(skip - 1)
	if ok {
		danErr.AppendCaller(fmt.Sprintf("%s:%d", file, line))
	}
	t.fmtdError(skip, errorCode, danErr)
	return danErr
}
func (t *Resp) NewErrWithTemp(errorCode int, temps ...string) error {
	skip := 2
	var danErr *danError

	danErr = newDanErr("", errorCode, temps...)
	_, file, line, ok := runtime.Caller(skip - 1)
	if ok {
		danErr.AppendCaller(fmt.Sprintf("%s:%d", file, line))
	}
	t.fmtdError(skip, errorCode, danErr)
	return danErr
}

// fmtdError 打印调试信息
func (t *Resp) fmtdError(skip int, errCode int, danErr *danError) {
	if t.mode != "prod" {
		if t.Transfer == nil {
			fmtd.Error("not set transfer")
			if danErr.level == levelInfo {
				fmtd.WithCaller(skip + 1).Info(t.Msg(SysErr))
			} else {
				fmtd.WithCaller(skip + 1).Error(t.Msg(SysErr))
			}

		} else {

			if danErr.level == levelInfo {
				fmtd.WithCaller(skip + 1).Info(fmt.Sprintf("%d %s", errCode, t.Msg(errCode)))
			} else {
				fmtd.WithCaller(skip + 1).Error(fmt.Sprintf("%d %s", errCode, t.Msg(errCode)))
			}
		}

	}
}
