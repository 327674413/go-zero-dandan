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

type SetUserSysMsgReadByClasLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		UserId   string  `json:"userId,optional" check:"required"`
		ClasList []int64 `json:"clasList,optional" check:"required"`
	}
	hasReq struct {
		UserId   bool
		ClasList bool
	}
}

func NewSetUserSysMsgReadByClasLogicGen(ctx context.Context, svc *svc.ServiceContext) *SetUserSysMsgReadByClasLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &SetUserSysMsgReadByClasLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *SetUserSysMsgReadByClasLogicGen) initReq(in *imRpc.SetUserSysMsgReadByClasReq) error {

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

	if in.ClasList != nil {
		l.req.ClasList = in.ClasList
		l.hasReq.ClasList = true
	} else {
		l.hasReq.ClasList = false
	}

	if l.hasReq.ClasList == false {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, "ClasList")
	}

	if len(l.req.ClasList) == 0 {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "ClasList")
	}

	return nil
}
