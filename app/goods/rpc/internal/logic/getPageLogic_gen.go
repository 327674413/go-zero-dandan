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

type GetPageLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		Page      int64
		Size      int64
		Sort      string
		TotalFlag int64
	}
	hasReq struct {
		Page      bool
		Size      bool
		Sort      bool
		TotalFlag bool
	}
}

func NewGetPageLogicGen(ctx context.Context, svc *svc.ServiceContext) *GetPageLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &GetPageLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(meta.Lang)),
		meta:   meta,
	}
}

func (l *GetPageLogicGen) initReq(req *goodsRpc.GetPageReq) error {

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

	if req.Sort != nil {
		l.req.Sort = *req.Sort
		l.hasReq.Sort = true
	} else {
		l.hasReq.Sort = false
	}

	if req.TotalFlag != nil {
		l.req.TotalFlag = *req.TotalFlag
		l.hasReq.TotalFlag = true
	} else {
		l.hasReq.TotalFlag = false
	}

	return nil
}