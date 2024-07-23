package logic

import (
	"context"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
)

type SetUserSysMsgReadByIdLogic struct {
	*SetUserSysMsgReadByIdLogicGen
}

func NewSetUserSysMsgReadByIdLogic(ctx context.Context, svc *svc.ServiceContext) *SetUserSysMsgReadByIdLogic {
	return &SetUserSysMsgReadByIdLogic{
		SetUserSysMsgReadByIdLogicGen: NewSetUserSysMsgReadByIdLogicGen(ctx, svc),
	}
}

func (l *SetUserSysMsgReadByIdLogic) SetUserSysMsgReadById(in *imRpc.SetUserSysMsgReadByIdReq) (*imRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	return &imRpc.ResultResp{}, nil
}
