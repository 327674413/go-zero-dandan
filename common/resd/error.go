package resd

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// todo::错误返回还要重构，现在还是有点怪
// Error 返回错误同时记录日志
func Error(err error, errorCode ...int) error {
	if e, ok := err.(*FailInfo); ok {
		return e
	}
	logx.WithCallerSkip(2).Error(err)
	if len(errorCode) > 0 {
		return Fail(err.Error(), errorCode[0])
	}
	return Fail(err.Error(), SysErr)

}
func NewErr(msg string, errorCode ...int) error {
	logx.WithCallerSkip(1).Error(errors.New(msg))
	if len(errorCode) > 0 {
		return Fail(msg, errorCode[0])
	}
	return Fail(msg, SysErr)
}

// ErrCtx 返回错误同时记录带上下文的日志
func ErrCtx(ctx context.Context, err error, errorCode ...int) error {
	if e, ok := err.(*FailInfo); ok {
		return e
	}
	logx.WithCallerSkip(1).WithContext(ctx).Error(ctx, msg)
	if len(errorCode) > 0 {
		return Fail(err.Error(), errorCode[0])
	}
	return Fail(err.Error(), SysErr)
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
		return Fail(text[4:], errCode)
	} else if msgType == "tmp:" {
		var tempData []string
		err := json.Unmarshal([]byte(text[4:]), &tempData)
		if err == nil {
			return FailCode(localize, errCode, tempData)
		}

	}
	return Fail(text, errCode)
}

// FailCode 按代码返回错误信息
func FailCode(localize *i18n.Localizer, errCode int, tempData ...[]string) error {
	var tempD []string
	if len(tempData) > 0 {
		tempD = tempData[0]
	}
	langMsg := Msg(localize, errCode, tempD)
	return Fail(langMsg, errCode)
}

// ApiFail 按错误内容返回错误信息
func ApiFail(localize *i18n.Localizer, failInfo error) error {
	if err, ok := failInfo.(*FailInfo); ok {
		err.Msg = Msg(localize, err.Code, err.Temps)
		return err
	} else {
		return failInfo
	}

}
