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

type GetUserPageLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		Page      int64
		Size      int64
		PlatId    string
		NeedTotal int64
		Match     map[string]*userRpc.MatchField
	}
	hasReq struct {
		Page      bool
		Size      bool
		PlatId    bool
		NeedTotal bool
		Match     bool
	}
}

func NewGetUserPageLogicGen(ctx context.Context, svc *svc.ServiceContext) *GetUserPageLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &GetUserPageLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *GetUserPageLogicGen) initReq(in *userRpc.GetUserPageReq) error {

	if in.Page != nil {
		l.req.Page = *in.Page
		l.hasReq.Page = true
	} else {
		l.hasReq.Page = false
	}

	if in.Size != nil {
		l.req.Size = *in.Size
		l.hasReq.Size = true
	} else {
		l.hasReq.Size = false
	}

	if in.PlatId != nil {
		l.req.PlatId = *in.PlatId
		l.hasReq.PlatId = true
	} else {
		l.hasReq.PlatId = false
	}

	if in.NeedTotal != nil {
		l.req.NeedTotal = *in.NeedTotal
		l.hasReq.NeedTotal = true
	} else {
		l.hasReq.NeedTotal = false
	}

	if in.Match != nil {
		l.req.Match = in.Match
		l.hasReq.Match = true
	} else {
		l.hasReq.Match = false
	}

	return nil
}
