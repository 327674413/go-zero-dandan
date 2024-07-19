// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPhoneLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		Phone     string
		PhoneArea string
		TempData  []string
		TempId    string
	}
	hasReq struct {
		Phone     bool
		PhoneArea bool
		TempData  bool
		TempId    bool
	}
}

func NewSendPhoneLogicGen(ctx context.Context, svc *svc.ServiceContext) *SendPhoneLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &SendPhoneLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *SendPhoneLogicGen) initReq(req *messageRpc.SendPhoneReq) error {

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

	if req.TempData != nil {
		l.req.TempData = req.TempData
		l.hasReq.TempData = true
	} else {
		l.hasReq.TempData = false
	}

	if req.TempId != nil {
		l.req.TempId = *req.TempId
		l.hasReq.TempId = true
	} else {
		l.hasReq.TempId = false
	}

	return nil
}
