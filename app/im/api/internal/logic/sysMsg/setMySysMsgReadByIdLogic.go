package sysMsg

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type SetMySysMsgReadByIdLogic struct {
	*SetMySysMsgReadByIdLogicGen
}

func NewSetMySysMsgReadByIdLogic(ctx context.Context, svc *svc.ServiceContext) *SetMySysMsgReadByIdLogic {
	return &SetMySysMsgReadByIdLogic{
		SetMySysMsgReadByIdLogicGen: NewSetMySysMsgReadByIdLogicGen(ctx, svc),
	}
}
func (l *SetMySysMsgReadByIdLogic) SetMySysMsgReadById(in *types.SetMySysMsgReadByIdReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	return
}
