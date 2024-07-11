// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateGroupMemberApplyLogicGen struct {
	ctx    context.Context
	svc    *svc.ServiceContext
	resd   *resd.Resp
	lang   string
	userId string
	platId string
	logx.Logger
	ReqApplyId        string
	ReqGroupId        string
	ReqOperateUid     string
	ReqOperateStateEm int64
	ReqPlatId         string
	ReqOperateMsg     string
	HasReq            struct {
		ApplyId        bool
		GroupId        bool
		OperateUid     bool
		OperateStateEm bool
		PlatId         bool
		OperateMsg     bool
	}
}

func NewOperateGroupMemberApplyLogicGen(ctx context.Context, svc *svc.ServiceContext) *OperateGroupMemberApplyLogicGen {
	lang, _ := ctx.Value("lang").(string)
	return &OperateGroupMemberApplyLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		lang:   lang,
		resd:   resd.NewResd(ctx, resd.I18n.NewLang(lang)),
	}
}

func (l *OperateGroupMemberApplyLogicGen) initReq(req *socialRpc.OperateGroupMemberApplyReq) error {
	var err error
	if err = l.initPlat(); err != nil {
		return l.resd.Error(err)
	}

	if req.ApplyId != nil {
		l.ReqApplyId = *req.ApplyId
		l.HasReq.ApplyId = true
	} else {
		l.HasReq.ApplyId = false
	}

	if req.GroupId != nil {
		l.ReqGroupId = *req.GroupId
		l.HasReq.GroupId = true
	} else {
		l.HasReq.GroupId = false
	}

	if req.OperateUid != nil {
		l.ReqOperateUid = *req.OperateUid
		l.HasReq.OperateUid = true
	} else {
		l.HasReq.OperateUid = false
	}

	if req.OperateStateEm != nil {
		l.ReqOperateStateEm = *req.OperateStateEm
		l.HasReq.OperateStateEm = true
	} else {
		l.HasReq.OperateStateEm = false
	}

	if req.PlatId != nil {
		l.ReqPlatId = *req.PlatId
		l.HasReq.PlatId = true
	} else {
		l.HasReq.PlatId = false
	}

	if req.OperateMsg != nil {
		l.ReqOperateMsg = *req.OperateMsg
		l.HasReq.OperateMsg = true
	} else {
		l.HasReq.OperateMsg = false
	}

	return nil
}

func (l *OperateGroupMemberApplyLogicGen) initUser() (err error) {
	userId, _ := l.ctx.Value("userId").(string)
	l.userId = userId
	return nil
}

func (l *OperateGroupMemberApplyLogicGen) initPlat() (err error) {
	platId, _ := l.ctx.Value("platId").(string)
	l.platId = platId
	return nil
}
