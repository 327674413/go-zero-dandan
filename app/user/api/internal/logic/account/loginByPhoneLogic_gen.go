// Code generated by goctl. DO NOT EDIT.
package account

import (
	"context"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type LoginByPhoneLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		Phone     string `json:"phone"`
		PhoneArea string `json:"phoneArea,optional"`
		OtpCode   string `json:"otpCode"`
		PortEm    int64  `json:"portEm"`
	}
	hasReq struct {
		Phone     bool
		PhoneArea bool
		OtpCode   bool
		PortEm    bool
	}
}

func NewLoginByPhoneLogicGen(ctx context.Context, svc *svc.ServiceContext) *LoginByPhoneLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &LoginByPhoneLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *LoginByPhoneLogicGen) initReq(in *types.LoginByPhoneReq) error {

	if in.Phone != nil {
		l.req.Phone = strings.TrimSpace(*in.Phone)
		l.hasReq.Phone = true
	} else {
		l.hasReq.Phone = false
	}

	if in.PhoneArea != nil {
		l.req.PhoneArea = strings.TrimSpace(*in.PhoneArea)
		l.hasReq.PhoneArea = true
	} else {
		l.hasReq.PhoneArea = false
	}

	if in.OtpCode != nil {
		l.req.OtpCode = strings.TrimSpace(*in.OtpCode)
		l.hasReq.OtpCode = true
	} else {
		l.hasReq.OtpCode = false
	}

	if in.PortEm != nil {
		l.req.PortEm = *in.PortEm
		l.hasReq.PortEm = true
	} else {
		l.hasReq.PortEm = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
