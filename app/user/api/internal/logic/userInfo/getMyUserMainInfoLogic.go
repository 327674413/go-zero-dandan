package userInfo

import (
	"context"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/utild/copier"
)

type GetMyUserMainInfoLogic struct {
	*GetMyUserMainInfoLogicGen
}

func NewGetMyUserMainInfoLogic(ctx context.Context, svc *svc.ServiceContext) *GetMyUserMainInfoLogic {
	return &GetMyUserMainInfoLogic{
		GetMyUserMainInfoLogicGen: NewGetMyUserMainInfoLogicGen(ctx, svc),
	}
}
func (l *GetMyUserMainInfoLogic) GetMyUserMainInfo() (resp *types.UserMainInfo, err error) {
	if err = l.initReq(); err != nil {
		return nil, l.resd.Error(err)
	}
	user, err := l.svc.UserRpc.GetUserByToken(l.ctx, &userRpc.TokenReq{Token: &l.meta.UserToken})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	resp = &types.UserMainInfo{}
	err = copier.Copy(&resp, user)
	if err = copier.Copy(&resp, user); err != nil {
		return nil, l.resd.Error(err)
	}
	return
}
