// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"
	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotPageByCursorLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		Page   int64
		Size   int64
		Cursor int64
		LastId string
	}
	hasReq struct {
		Page   bool
		Size   bool
		Cursor bool
		LastId bool
	}
}

func NewGetHotPageByCursorLogicGen(ctx context.Context, svc *svc.ServiceContext) *GetHotPageByCursorLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &GetHotPageByCursorLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(meta.Lang)),
		meta:   meta,
	}
}

func (l *GetHotPageByCursorLogicGen) initReq(req *goodsRpc.GetHotPageByCursorReq) error {

	if req.Page != nil {
		l.req.Page = *req.Page
		l.hasReq.Page = true
	} else {
		l.hasReq.Page = false
	}

	if req.Size != nil {
		l.req.Size = *req.Size
		l.hasReq.Size = true
	} else {
		l.hasReq.Size = false
	}

	if req.Cursor != nil {
		l.req.Cursor = *req.Cursor
		l.hasReq.Cursor = true
	} else {
		l.hasReq.Cursor = false
	}

	if req.LastId != nil {
		l.req.LastId = *req.LastId
		l.hasReq.LastId = true
	} else {
		l.hasReq.LastId = false
	}

	return nil
}