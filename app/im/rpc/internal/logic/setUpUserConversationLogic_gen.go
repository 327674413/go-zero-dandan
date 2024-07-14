// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUpUserConversationLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		SendId   string
		RecvId   string
		ChatType int64
	}
	hasReq struct {
		SendId   bool
		RecvId   bool
		ChatType bool
	}
}

func NewSetUpUserConversationLogicGen(ctx context.Context, svc *svc.ServiceContext) *SetUpUserConversationLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &SetUpUserConversationLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(meta.Lang)),
		meta:   meta,
	}
}

func (l *SetUpUserConversationLogicGen) initReq(req *imRpc.SetUpUserConversationReq) error {

	if req.SendId != nil {
		l.req.SendId = *req.SendId
		l.hasReq.SendId = true
	} else {
		l.hasReq.SendId = false
	}

	if req.RecvId != nil {
		l.req.RecvId = *req.RecvId
		l.hasReq.RecvId = true
	} else {
		l.hasReq.RecvId = false
	}

	if req.ChatType != nil {
		l.req.ChatType = *req.ChatType
		l.hasReq.ChatType = true
	} else {
		l.hasReq.ChatType = false
	}

	return nil
}
