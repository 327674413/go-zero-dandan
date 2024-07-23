// Code generated by goctl. DO NOT EDIT.
package sysMsg

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
)

type SetMySysMsgReadByIdLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		MsgClasEm int64   `json:"msgClasEm,optional"`
		Ids       []int64 `json:"ids,optional"`
	}
	hasReq struct {
		MsgClasEm bool
		Ids       bool
	}
}

func NewSetMySysMsgReadByIdLogicGen(ctx context.Context, svc *svc.ServiceContext) *SetMySysMsgReadByIdLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &SetMySysMsgReadByIdLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *SetMySysMsgReadByIdLogicGen) initReq(in *types.SetMySysMsgReadByIdReq) error {

	if in.MsgClasEm != nil {
		l.req.MsgClasEm = *in.MsgClasEm
		l.hasReq.MsgClasEm = true
	} else {
		l.hasReq.MsgClasEm = false
	}

	if l.hasReq.MsgClasEm == false {
		return resd.NewErrWithTempCtx(l.ctx, "缺少参数MsgClasEm", resd.ErrReqFieldRequired1, "MsgClasEm")
	}

	if in.Ids != nil {
		l.req.Ids = in.Ids
		l.hasReq.Ids = true
	} else {
		l.hasReq.Ids = false
	}

	if l.hasReq.Ids == false {
		return resd.NewErrWithTempCtx(l.ctx, "缺少参数Ids", resd.ErrReqFieldRequired1, "Ids")
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
