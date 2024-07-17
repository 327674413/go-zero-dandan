package userInfo

import (
	"context"
	"go-zero-dandan/app/user/api/internal/biz"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/utild/copier"
)

type EditMyInfoLogic struct {
	*EditMyInfoLogicGen
}

func NewEditMyInfoLogic(ctx context.Context, svc *svc.ServiceContext) *EditMyInfoLogic {
	return &EditMyInfoLogic{
		EditMyInfoLogicGen: NewEditMyInfoLogicGen(ctx, svc),
	}
}

func (l *EditMyInfoLogic) EditMyInfo(in *types.EditMyInfoReq) (resp *types.SuccessResp, err error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	userBiz := biz.NewUserBiz(l.ctx, l.svc)
	editUserInfo := &userRpc.EditUserInfoReq{}
	copier.Copy(&editUserInfo, l.req)
	editUserInfo.Id = &l.meta.UserId
	err = userBiz.EditUserInfo(editUserInfo)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	return &types.SuccessResp{Msg: ""}, nil
}
