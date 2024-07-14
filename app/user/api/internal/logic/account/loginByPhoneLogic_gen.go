// Code generated by goctl. DO NOT EDIT.
package account

import (
	"context"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type LoginByPhoneLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	lang         string
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
	hasUserInfo  bool
	mustUserInfo bool
	ReqPhone     string `json:"phone"`
	ReqPhoneArea string `json:"phoneArea,optional"`
	ReqOtpCode   string `json:"otpCode"`
	ReqPortEm    int64  `json:"portEm"`
	HasReq       struct {
		Phone     bool
		PhoneArea bool
		OtpCode   bool
		PortEm    bool
	}
}

func NewLoginByPhoneLogicGen(ctx context.Context, svc *svc.ServiceContext) *LoginByPhoneLogicGen {
	lang, _ := ctx.Value("lang").(string)
	return &LoginByPhoneLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		lang:   lang,
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(lang)),
	}
}

func (l *LoginByPhoneLogicGen) initReq(req *types.LoginByPhoneReq) error {
	var err error
	if err = l.initPlat(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}

	if req.Phone != nil {
		l.ReqPhone = *req.Phone
		l.HasReq.Phone = true
	} else {
		l.HasReq.Phone = false
	}

	if req.PhoneArea != nil {
		l.ReqPhoneArea = *req.PhoneArea
		l.HasReq.PhoneArea = true
	} else {
		l.HasReq.PhoneArea = false
	}

	if req.OtpCode != nil {
		l.ReqOtpCode = *req.OtpCode
		l.HasReq.OtpCode = true
	} else {
		l.HasReq.OtpCode = false
	}

	if req.PortEm != nil {
		l.ReqPortEm = *req.PortEm
		l.HasReq.PortEm = true
	} else {
		l.HasReq.PortEm = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}

func (l *LoginByPhoneLogicGen) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.ErrUserMainInfo)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *LoginByPhoneLogicGen) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.ErrPlatClas)
	}
	platId, _ := l.ctx.Value("platId").(string)
	if platId == "" {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.ErrPlatId)
	}
	l.platId = platId
	l.platClasEm = platClasEm
	return nil
}
