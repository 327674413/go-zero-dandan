package logic

import (
	"context"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
)

type BindUnionUserLogic struct {
	*BindUnionUserLogicGen
}

func NewBindUnionUserLogic(ctx context.Context, svc *svc.ServiceContext) *BindUnionUserLogic {
	return &BindUnionUserLogic{
		BindUnionUserLogicGen: NewBindUnionUserLogicGen(ctx, svc),
	}
}

func (l *BindUnionUserLogic) BindUnionUser(req *userRpc.BindUnionUserReq) (*userRpc.BindUnionUserResp, error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}

	return &userRpc.BindUnionUserResp{}, nil
}
