// Code generated by goctl. DO NOT EDIT.
package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type CreateGroupMemberApplyLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		GroupId   string `json:"groupId,optional"`
		GroupCode string `json:"groupCode,optional"`
		ApplyMsg  string `json:"applyMsg,optional"`
		SourceEm  int64  `json:"sourceEm,optional"`
	}
	hasReq struct {
		GroupId   bool
		GroupCode bool
		ApplyMsg  bool
		SourceEm  bool
	}
}

func NewCreateGroupMemberApplyLogicGen(ctx context.Context, svc *svc.ServiceContext) *CreateGroupMemberApplyLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &CreateGroupMemberApplyLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *CreateGroupMemberApplyLogicGen) initReq(req *types.CreateGroupMemberApplyReq) error {

	if req.GroupId != nil {
		l.req.GroupId = strings.TrimSpace(*req.GroupId)
		l.hasReq.GroupId = true
	} else {
		l.hasReq.GroupId = false
	}

	if req.GroupCode != nil {
		l.req.GroupCode = strings.TrimSpace(*req.GroupCode)
		l.hasReq.GroupCode = true
	} else {
		l.hasReq.GroupCode = false
	}

	if req.ApplyMsg != nil {
		l.req.ApplyMsg = strings.TrimSpace(*req.ApplyMsg)
		l.hasReq.ApplyMsg = true
	} else {
		l.hasReq.ApplyMsg = false
	}

	if req.SourceEm != nil {
		l.req.SourceEm = *req.SourceEm
		l.hasReq.SourceEm = true
	} else {
		l.hasReq.SourceEm = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
