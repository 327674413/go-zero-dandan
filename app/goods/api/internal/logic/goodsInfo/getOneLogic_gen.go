// Code generated by goctl. DO NOT EDIT.
package goodsInfo

import (
	"context"

	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type GetOneLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		Id string `json:"id,optional"`
	}
	hasReq struct {
		Id bool
	}
}

func NewGetOneLogicGen(ctx context.Context, svc *svc.ServiceContext) *GetOneLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &GetOneLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *GetOneLogicGen) initReq(req *types.IdReq) error {

	if req.Id != nil {
		l.req.Id = strings.TrimSpace(*req.Id)
		l.hasReq.Id = true
	} else {
		l.hasReq.Id = false
	}

	if l.hasReq.Id == false {
		return resd.NewErrWithTempCtx(l.ctx, "缺少参数Id", resd.ErrReqFieldRequired1, "Id")
	}

	if l.req.Id == "" {
		return resd.NewErrWithTempCtx(l.ctx, "Id不得为空", resd.ErrReqFieldEmpty1, "Id")
	}
	l.hasUserInfo = true

	return nil
}
