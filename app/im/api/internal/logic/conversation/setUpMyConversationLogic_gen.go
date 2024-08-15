// Code generated by goctl. DO NOT EDIT.
package conversation

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type SetUpMyConversationLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		RecvId   string `json:"recvId,optional"`
		ChatType int64  `json:"chatType,optional"`
	}
	hasReq struct {
		RecvId   bool
		ChatType bool
	}
}

func NewSetUpMyConversationLogicGen(ctx context.Context, svc *svc.ServiceContext) *SetUpMyConversationLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &SetUpMyConversationLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *SetUpMyConversationLogicGen) initReq(in *types.SetUpMyConversationReq) error {

	if in.RecvId != nil {
		l.req.RecvId = strings.TrimSpace(*in.RecvId)
		l.hasReq.RecvId = true
	} else {
		l.hasReq.RecvId = false
	}

	if l.hasReq.RecvId == false {
		return resd.NewErrWithTempCtx(l.ctx, "缺少参数RecvId", resd.ErrReqFieldRequired1, "RecvId")
	}

	if l.req.RecvId == "" {
		return resd.NewErrWithTempCtx(l.ctx, "RecvId不得为空", resd.ErrReqFieldEmpty1, "RecvId")
	}

	if in.ChatType != nil {
		l.req.ChatType = *in.ChatType
		l.hasReq.ChatType = true
	} else {
		l.hasReq.ChatType = false
	}

	if l.hasReq.ChatType == false {
		return resd.NewErrWithTempCtx(l.ctx, "缺少参数ChatType", resd.ErrReqFieldRequired1, "ChatType")
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}