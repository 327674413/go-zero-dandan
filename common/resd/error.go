package resd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/constd"
)

// danError自定义错误类型，兼容rpc错误
type danError struct {
	Result     bool     `json:"result"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
	Temps      []string `json:"temps"`
	level      string   `json:"-"` //用来区分异常还是业务合法的，但想全走日志，当作完整的链路返回，就暂时不需要了
	callers    []string `json:"-"`
	callerSkip int      `json:"-"`
}

func JsonToErr(ctxOrNil context.Context, jsonStr string) error {
	var ctx context.Context
	if ctxOrNil == nil {
		ctx = context.Background()
	} else {
		ctx = ctxOrNil
	}
	if jsonStr[:1] != "{" {
		return NewErrCtx(ctx, "RPC服务错误："+jsonStr, ErrRpcService)
	}
	err := danError{}
	jErr := json.Unmarshal([]byte(jsonStr), &err)
	if jErr != nil {
		return ErrorCtx(ctx, jErr, ErrJsonDecode)
	}
	return &err
}

// error实现
func (t *danError) Error() string {
	return fmt.Sprintf("%s", t.Msg)
}

// SetTemps 设置模i18n的模版
func (t *danError) SetTemps(temps []string) *danError {
	t.Temps = temps
	return t
}

// GetTemps 获取当前设置了的模版
func (t *danError) GetTemps() []string {
	return t.Temps
}

func (t *danError) AppendCaller(caller string) *danError {
	t.callers = append(t.callers, caller)
	return t
}

// 创建错误
func newDanErr(msg string, errCode int, tempStr ...string) *danError {
	res := &danError{
		Result:  false,
		Code:    errCode,
		Msg:     msg,
		level:   levelError,
		callers: make([]string, 0),
	}
	if len(tempStr) > 0 {
		res.SetTemps(tempStr)
	}
	return res
}

func errdWithTemp(ctxOrNil context.Context, langOrNil *Lang, err error, initErrCode int, temps ...string) *danError {
	skip := 3
	if err == nil {
		return newDanErr("err is nil", initErrCode)
	}
	danErr, ok := err.(*danError)
	if ok {
		//如果有传，则覆盖原先的错误码
		danErr.Code = initErrCode
		danErr.SetTemps(temps)
	} else {
		//不是自定义错误，创建
		danErr = newDanErr(err.Error(), initErrCode, temps...)
	}
	if Mode != constd.ModePro {
		printErr(skip, danErr, langOrNil)
	}
	if Mode != constd.ModeDev {
		if ctxOrNil != nil {
			logx.WithCallerSkip(skip).Error(err)
		} else {
			logx.WithCallerSkip(skip).WithContext(ctxOrNil).Error(err)
		}
	}
	return danErr
}
func newErr(ctxOrNil context.Context, langOrNil *Lang, errMsg string, initErrCode ...int) error {
	errorCode := ErrSys
	if len(initErrCode) > 0 {
		errorCode = initErrCode[0]
	}
	skip := 3
	var danErr *danError
	danErr = newDanErr(errMsg, errorCode)

	if Mode != constd.ModePro {
		printErr(skip, danErr, langOrNil)
	}
	if Mode != constd.ModeDev {
		if ctxOrNil != nil {
			logx.WithCallerSkip(skip).Error(danErr)
		} else {
			logx.WithCallerSkip(skip).WithContext(ctxOrNil).Error(danErr)
		}
	}
	return danErr
}
func newErrWithTemp(ctxOrNil context.Context, langOrNil *Lang, errMsg string, initErrCode int, temps ...string) error {
	errorCode := initErrCode
	skip := 3
	var danErr *danError
	danErr = newDanErr(errMsg, errorCode, temps...)
	if Mode != constd.ModePro {
		printErr(skip, danErr, langOrNil)
	}
	if Mode != constd.ModeDev {
		if ctxOrNil != nil {
			logx.WithCallerSkip(skip).Error(danErr)
		} else {
			logx.WithCallerSkip(skip).WithContext(ctxOrNil).Error(danErr)
		}
	}
	return danErr
}

// Error 对error进行封装返回
func Error(err error, newErrCode ...int) error {
	if len(newErrCode) > 0 {
		return errdWithTemp(nil, nil, err, newErrCode[0])
	} else {
		if danErr, ok := err.(*danError); ok {
			return errdWithTemp(nil, nil, err, danErr.Code)
		} else {
			return errdWithTemp(nil, nil, err, ErrSys)
		}
	}
}

// ErrorCtx 对error进行封装返回,带上下文
func ErrorCtx(ctx context.Context, err error, newErrCode ...int) error {
	if len(newErrCode) > 0 {
		return errdWithTemp(ctx, nil, err, newErrCode[0])
	} else {
		if danErr, ok := err.(*danError); ok {
			return errdWithTemp(ctx, nil, err, danErr.Code)
		} else {
			return errdWithTemp(ctx, nil, err, ErrSys)
		}
	}
}

// ErrorWithTemp 对error进行封装返回，带模版
func ErrorWithTemp(err error, initErrCode int, temps ...string) error {
	return errdWithTemp(nil, nil, err, initErrCode, temps...)
}

// ErrorWithTempCtx 对error进行封装返回,带上下文，带模版
func ErrorWithTempCtx(ctx context.Context, err error, initErrCode int, temps ...string) error {
	return errdWithTemp(ctx, nil, err, initErrCode, temps...)
}

// NewErr 创建新的error
func NewErr(errMsgOrEmpty string, initErrCode ...int) error {
	return newErr(nil, nil, errMsgOrEmpty, initErrCode...)
}

// NewErrCtx 创建新的error，带上下文
func NewErrCtx(ctx context.Context, errMsgOrEmpty string, initErrCode ...int) error {
	return newErr(ctx, nil, errMsgOrEmpty, initErrCode...)
}

// NewErrWithTemp 创建新的error，带模版,errorCode用resd.xxxxx，temps直接用语言包里的变量
func NewErrWithTemp(errMsgOrEmpty string, errorCode int, temps ...string) error {
	return newErrWithTemp(nil, nil, errMsgOrEmpty, errorCode, temps...)
}

// NewErrWithTempCtx 创建新的error，带模版,errorCode用resd.xxxxx，temps直接用语言包里的变量，带上下文
func NewErrWithTempCtx(ctx context.Context, errMsgOrEmpty string, errorCode int, temps ...string) error {
	return newErrWithTemp(ctx, nil, errMsgOrEmpty, errorCode, temps...)
}

func AssertErr(failErr error) (*danError, bool) {
	if err, ok := failErr.(*danError); ok {
		return err, ok
	} else {
		return nil, false
	}
}
func IsUserNotLoginErr(failErr error) bool {
	if err, ok := failErr.(*danError); ok {
		if err.Code == ErrAuthUserNotLogin {
			return true
		}
	}
	return false
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
