// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByTokenLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		Token string
	}
	hasReq struct {
		Token bool
	}
}

func NewGetUserByTokenLogicGen(ctx context.Context, svc *svc.ServiceContext) *GetUserByTokenLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &GetUserByTokenLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *GetUserByTokenLogicGen) initReq(in *userRpc.TokenReq) error {

	if in.Token != nil {
		l.req.Token = *in.Token
		l.hasReq.Token = true
	} else {
		l.hasReq.Token = false
	}

	return nil
}
