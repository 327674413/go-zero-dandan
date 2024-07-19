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

type OperateGroupMemberApplyLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		ApplyId        string
		GroupId        string
		OperateUid     string
		OperateStateEm int64
		PlatId         string
		OperateMsg     string
	}
	hasReq struct {
		ApplyId        bool
		GroupId        bool
		OperateUid     bool
		OperateStateEm bool
		PlatId         bool
		OperateMsg     bool
	}
}

func NewOperateGroupMemberApplyLogicGen(ctx context.Context, svc *svc.ServiceContext) *OperateGroupMemberApplyLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &OperateGroupMemberApplyLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *OperateGroupMemberApplyLogicGen) initReq(req *socialRpc.OperateGroupMemberApplyReq) error {

	if req.ApplyId != nil {
		l.req.ApplyId = *req.ApplyId
		l.hasReq.ApplyId = true
	} else {
		l.hasReq.ApplyId = false
	}

	if req.GroupId != nil {
		l.req.GroupId = *req.GroupId
		l.hasReq.GroupId = true
	} else {
		l.hasReq.GroupId = false
	}

	if req.OperateUid != nil {
		l.req.OperateUid = *req.OperateUid
		l.hasReq.OperateUid = true
	} else {
		l.hasReq.OperateUid = false
	}

	if req.OperateStateEm != nil {
		l.req.OperateStateEm = *req.OperateStateEm
		l.hasReq.OperateStateEm = true
	} else {
		l.hasReq.OperateStateEm = false
	}

	if req.PlatId != nil {
		l.req.PlatId = *req.PlatId
		l.hasReq.PlatId = true
	} else {
		l.hasReq.PlatId = false
	}

	if req.OperateMsg != nil {
		l.req.OperateMsg = *req.OperateMsg
		l.hasReq.OperateMsg = true
	} else {
		l.hasReq.OperateMsg = false
	}

	return nil
}
