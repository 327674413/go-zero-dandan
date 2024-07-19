// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupListLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		UserId string
		PlatId string
	}
	hasReq struct {
		UserId bool
		PlatId bool
	}
}

func NewGetUserGroupListLogicGen(ctx context.Context, svc *svc.ServiceContext) *GetUserGroupListLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &GetUserGroupListLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *GetUserGroupListLogicGen) initReq(req *socialRpc.GetUserGroupListReq) error {

	if req.UserId != nil {
		l.req.UserId = *req.UserId
		l.hasReq.UserId = true
	} else {
		l.hasReq.UserId = false
	}

	if req.PlatId != nil {
		l.req.PlatId = *req.PlatId
		l.hasReq.PlatId = true
	} else {
		l.hasReq.PlatId = false
	}

	return nil
}
