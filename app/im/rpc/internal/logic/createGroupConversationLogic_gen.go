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

type CreateGroupConversationLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		GroupId  string
		CreateId string
	}
	hasReq struct {
		GroupId  bool
		CreateId bool
	}
}

func NewCreateGroupConversationLogicGen(ctx context.Context, svc *svc.ServiceContext) *CreateGroupConversationLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &CreateGroupConversationLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *CreateGroupConversationLogicGen) initReq(in *imRpc.CreateGroupConversationReq) error {

	if in.GroupId != nil {
		l.req.GroupId = *in.GroupId
		l.hasReq.GroupId = true
	} else {
		l.hasReq.GroupId = false
	}

	if in.CreateId != nil {
		l.req.CreateId = *in.CreateId
		l.hasReq.CreateId = true
	} else {
		l.hasReq.CreateId = false
	}

	return nil
}
