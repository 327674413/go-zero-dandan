// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupMemberApplyLogicGen struct {
	ctx    context.Context
	svc    *svc.ServiceContext
	resd   *resd.Resp
	lang   string
	userId string
	platId string
	logx.Logger
	ReqPlatId       string
	ReqGroupId      string
	ReqApplyMsg     string
	ReqJoinSourceEm int64
	ReqInviteUid    string
	HasReq          struct {
		PlatId       bool
		GroupId      bool
		ApplyMsg     bool
		JoinSourceEm bool
		InviteUid    bool
	}
}

func NewCreateGroupMemberApplyLogicGen(ctx context.Context, svc *svc.ServiceContext) *CreateGroupMemberApplyLogicGen {
	lang, _ := ctx.Value("lang").(string)
	return &CreateGroupMemberApplyLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		lang:   lang,
		resd:   resd.NewResd(ctx, resd.I18n.NewLang(lang)),
	}
}

func (l *CreateGroupMemberApplyLogicGen) initReq(req *socialRpc.CreateGroupMemberApplyReq) error {
	var err error
	if err = l.initPlat(); err != nil {
		return l.resd.Error(err)
	}

	if req.PlatId != nil {
		l.ReqPlatId = *req.PlatId
		l.HasReq.PlatId = true
	} else {
		l.HasReq.PlatId = false
	}

	if req.GroupId != nil {
		l.ReqGroupId = *req.GroupId
		l.HasReq.GroupId = true
	} else {
		l.HasReq.GroupId = false
	}

	if req.ApplyMsg != nil {
		l.ReqApplyMsg = *req.ApplyMsg
		l.HasReq.ApplyMsg = true
	} else {
		l.HasReq.ApplyMsg = false
	}

	if req.JoinSourceEm != nil {
		l.ReqJoinSourceEm = *req.JoinSourceEm
		l.HasReq.JoinSourceEm = true
	} else {
		l.HasReq.JoinSourceEm = false
	}

	if req.InviteUid != nil {
		l.ReqInviteUid = *req.InviteUid
		l.HasReq.InviteUid = true
	} else {
		l.HasReq.InviteUid = false
	}

	return nil
}

func (l *CreateGroupMemberApplyLogicGen) initUser() (err error) {
	userId, _ := l.ctx.Value("userId").(string)
	l.userId = userId
	return nil
}

func (l *CreateGroupMemberApplyLogicGen) initPlat() (err error) {
	platId, _ := l.ctx.Value("platId").(string)
	l.platId = platId
	return nil
}