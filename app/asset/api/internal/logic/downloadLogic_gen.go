// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"

	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type DownloadLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		Id string
	}
	hasReq struct {
		Id bool
	}
}

func NewDownloadLogicGen(ctx context.Context, svc *svc.ServiceContext) *DownloadLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &DownloadLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *DownloadLogicGen) initReq(req *types.DownloadReq) error {

	if req.Id != nil {
		l.req.Id = strings.TrimSpace(*req.Id)
		l.hasReq.Id = true
	} else {
		l.hasReq.Id = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
