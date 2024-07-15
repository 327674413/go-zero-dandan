// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegByAccountLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		PlatId    string
		Account   string
		Password  string
		Nickname  string
		Phone     string
		PhoneArea string
		SexEm     int64
		Email     string
		AvatarImg string
		IsLogin   int64
	}
	hasReq struct {
		PlatId    bool
		Account   bool
		Password  bool
		Nickname  bool
		Phone     bool
		PhoneArea bool
		SexEm     bool
		Email     bool
		AvatarImg bool
		IsLogin   bool
	}
}

func NewRegByAccountLogicGen(ctx context.Context, svc *svc.ServiceContext) *RegByAccountLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &RegByAccountLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(meta.Lang)),
		meta:   meta,
	}
}

func (l *RegByAccountLogicGen) initReq(req *userRpc.RegByAccountReq) error {

	if req.PlatId != nil {
		l.req.PlatId = *req.PlatId
		l.hasReq.PlatId = true
	} else {
		l.hasReq.PlatId = false
	}

	if req.Account != nil {
		l.req.Account = *req.Account
		l.hasReq.Account = true
	} else {
		l.hasReq.Account = false
	}

	if req.Password != nil {
		l.req.Password = *req.Password
		l.hasReq.Password = true
	} else {
		l.hasReq.Password = false
	}

	if req.Nickname != nil {
		l.req.Nickname = *req.Nickname
		l.hasReq.Nickname = true
	} else {
		l.hasReq.Nickname = false
	}

	if req.Phone != nil {
		l.req.Phone = *req.Phone
		l.hasReq.Phone = true
	} else {
		l.hasReq.Phone = false
	}

	if req.PhoneArea != nil {
		l.req.PhoneArea = *req.PhoneArea
		l.hasReq.PhoneArea = true
	} else {
		l.hasReq.PhoneArea = false
	}

	if req.SexEm != nil {
		l.req.SexEm = *req.SexEm
		l.hasReq.SexEm = true
	} else {
		l.hasReq.SexEm = false
	}

	if req.Email != nil {
		l.req.Email = *req.Email
		l.hasReq.Email = true
	} else {
		l.hasReq.Email = false
	}

	if req.AvatarImg != nil {
		l.req.AvatarImg = *req.AvatarImg
		l.hasReq.AvatarImg = true
	} else {
		l.hasReq.AvatarImg = false
	}

	if req.IsLogin != nil {
		l.req.IsLogin = *req.IsLogin
		l.hasReq.IsLogin = true
	} else {
		l.hasReq.IsLogin = false
	}

	return nil
}