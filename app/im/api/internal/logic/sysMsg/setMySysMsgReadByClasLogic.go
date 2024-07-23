package sysMsg

import (
	"context"
	"go-zero-dandan/app/im/rpc/types/imRpc"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type SetMySysMsgReadByClasLogic struct {
	*SetMySysMsgReadByClasLogicGen
}

func NewSetMySysMsgReadByClasLogic(ctx context.Context, svc *svc.ServiceContext) *SetMySysMsgReadByClasLogic {
	return &SetMySysMsgReadByClasLogic{
		SetMySysMsgReadByClasLogicGen: NewSetMySysMsgReadByClasLogicGen(ctx, svc),
	}
}
func (l *SetMySysMsgReadByClasLogic) SetMySysMsgReadByClas(in *types.SetMySysMsgReadByClasReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	_, err = l.svc.ImRpc.SetUserSysMsgReadByClas(l.ctx, &imRpc.SetUserSysMsgReadByClasReq{
		UserId:   &l.meta.UserId,
		ClasList: l.req.MsgClasEms,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	return &types.ResultResp{
		Result: true,
	}, nil
}
