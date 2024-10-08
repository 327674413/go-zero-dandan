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

type CreateGroupMemberApplyLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		PlatId       string
		GroupId      string
		ApplyMsg     string
		JoinSourceEm int64
		InviteUid    string
	}
	hasReq struct {
		PlatId       bool
		GroupId      bool
		ApplyMsg     bool
		JoinSourceEm bool
		InviteUid    bool
	}
}

func NewCreateGroupMemberApplyLogicGen(ctx context.Context, svc *svc.ServiceContext) *CreateGroupMemberApplyLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &CreateGroupMemberApplyLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *CreateGroupMemberApplyLogicGen) initReq(in *socialRpc.CreateGroupMemberApplyReq) error {

	if in.PlatId != nil {
		l.req.PlatId = *in.PlatId
		l.hasReq.PlatId = true
	} else {
		l.hasReq.PlatId = false
	}

	if in.GroupId != nil {
		l.req.GroupId = *in.GroupId
		l.hasReq.GroupId = true
	} else {
		l.hasReq.GroupId = false
	}

	if in.ApplyMsg != nil {
		l.req.ApplyMsg = *in.ApplyMsg
		l.hasReq.ApplyMsg = true
	} else {
		l.hasReq.ApplyMsg = false
	}

	if in.JoinSourceEm != nil {
		l.req.JoinSourceEm = *in.JoinSourceEm
		l.hasReq.JoinSourceEm = true
	} else {
		l.hasReq.JoinSourceEm = false
	}

	if in.InviteUid != nil {
		l.req.InviteUid = *in.InviteUid
		l.hasReq.InviteUid = true
	} else {
		l.hasReq.InviteUid = false
	}

	return nil
}
