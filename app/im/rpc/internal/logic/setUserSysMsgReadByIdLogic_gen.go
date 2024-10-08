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

type SetUserSysMsgReadByIdLogicGen struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	resd *resd.Resp
	meta *typed.ReqMeta
	logx.Logger
	req struct {
		UserId    string   `json:"userId,optional" check:"required"`
		MsgIdList []string `json:"msgIdList,optional" check:"required"`
	}
	hasReq struct {
		UserId    bool
		MsgIdList bool
	}
}

func NewSetUserSysMsgReadByIdLogicGen(ctx context.Context, svc *svc.ServiceContext) *SetUserSysMsgReadByIdLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &SetUserSysMsgReadByIdLogicGen{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *SetUserSysMsgReadByIdLogicGen) initReq(in *imRpc.SetUserSysMsgReadByIdReq) error {

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

	if in.MsgIdList != nil {
		l.req.MsgIdList = in.MsgIdList
		l.hasReq.MsgIdList = true
	} else {
		l.hasReq.MsgIdList = false
	}

	if l.hasReq.MsgIdList == false {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, "MsgIdList")
	}

	if len(l.req.MsgIdList) == 0 {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, "MsgIdList")
	}

	return nil
}
