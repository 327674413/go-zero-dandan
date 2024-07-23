package resd

import (
	"context"
	"fmt"
	"go-zero-dandan/common/fmtd"
)

var Mode string

type Resp struct {
	ctx  context.Context
	mode string
	*Lang
}

const (
	levelError = "error" //真正的异常报错，需要关注
	levelInfo  = "info"  //正常业务逻辑校验的问题
)

// NewResp 创建统一错误返回
func NewResp(ctxOrNil context.Context, lang string) *Resp {
	i18nLang := I18n.NewLang(lang)
	if ctxOrNil == nil {
		return &Resp{
			ctx:  context.Background(),
			Lang: i18nLang,
		}
	} else {
		return &Resp{
			ctx:  ctxOrNil,
			Lang: i18nLang,
		}
	}
}

// 错误返回中的错误方法，可以多次嵌套Error，追溯整个来源
func (t *Resp) Error(err error, initErrCode ...int) error {
	if len(initErrCode) > 0 {
		return errdWithTemp(t.ctx, t.Lang, err, initErrCode[0])
	} else {
		if danErr, ok := err.(*danError); ok {
			return errdWithTemp(t.ctx, t.Lang, err, danErr.Code, danErr.GetTemps()...)
		} else {
			return errdWithTemp(t.ctx, t.Lang, err, ErrSys, danErr.GetTemps()...)
		}
	}

}

// ErrorWithTemp 错误返回中的错误方法，重新指定错误码，注入模版
func (t *Resp) ErrorWithTemp(err error, initErrCode int, temps ...string) error {
	return errdWithTemp(t.ctx, t.Lang, err, initErrCode, temps...)
}

// NewErr 报错返回，需要关注报错原因的
func (t *Resp) NewErr(initErrCode ...int) error {
	errCode := ErrSys
	if len(initErrCode) > 0 {
		errCode = initErrCode[0]
	}
	errMsg := t.Lang.Msg(errCode)
	return newErr(t.ctx, t.Lang, errMsg, errCode)
}

// NewErrWithTemp 带模版的异常报错
func (t *Resp) NewErrWithTemp(initErrCode int, temps ...string) error {
	errCode := initErrCode
	errMsg := t.Lang.Msg(errCode, temps...)
	return newErrWithTemp(t.ctx, t.Lang, errMsg, errCode, temps...)
}

// fmtd 打印调试信息
func printErr(skip int, danErr *danError, lang *Lang) {
	if lang != nil && danErr.Msg == "" {
		fmtd.WithCaller(skip + 1).Error(fmt.Sprintf("[%d]%s", danErr.Code, lang.Msg(danErr.Code, danErr.GetTemps()...)))
	} else {
		fmtd.WithCaller(skip + 1).Error(fmt.Sprintf("[%d]%s", danErr.Code, danErr.Msg))
	}

}
