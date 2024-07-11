package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
)

type GetUserGroupListLogic struct {
	*GetUserGroupListLogicGen
}

func NewGetUserGroupListLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserGroupListLogic {
	return &GetUserGroupListLogic{
		GetUserGroupListLogicGen: NewGetUserGroupListLogicGen(ctx, svc),
	}
}

func (l *GetUserGroupListLogic) GetUserGroupList(in *socialRpc.GetUserGroupListReq) (*socialRpc.GroupListResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &socialRpc.GroupListResp{}, nil
}
func (l *GetUserGroupListLogic) checkReqParams(in *socialRpc.GetUserGroupListReq) error {
	return nil
}
