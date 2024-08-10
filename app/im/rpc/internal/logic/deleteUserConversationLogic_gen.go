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

type DeleteUserConversationLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		UserId         string `json:"userId,optional" check:"required"`
		ConversationId string `json:"conversationId,optional" check:"required"`
	}
	hasReq struct {
		UserId         bool
		ConversationId bool
	}
}

func NewDeleteUserConversationLogicGen(ctx context.Context, svc *svc.ServiceContext) *DeleteUserConversationLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &DeleteUserConversationLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *DeleteUserConversationLogicGen) initReq(in *imRpc.DeleteUserConversationReq) error {

	if in.UserId != nil {
		l.req.UserId = *in.UserId
		l.hasReq.UserId = true
	} else {
		l.hasReq.UserId = false
	}

	if l.hasReq.UserId == false {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, "UserId")
	}

	if l.req.UserId == "" {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "UserId")
	}

	if in.ConversationId != nil {
		l.req.ConversationId = *in.ConversationId
		l.hasReq.ConversationId = true
	} else {
		l.hasReq.ConversationId = false
	}

	if l.hasReq.ConversationId == false {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, "ConversationId")
	}

	if l.req.ConversationId == "" {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "ConversationId")
	}

	return nil
}
