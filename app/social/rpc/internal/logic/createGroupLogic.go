package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
)

type CreateGroupLogic struct {
	*CreateGroupLogicGen
}

func NewCreateGroupLogic(ctx context.Context, svc *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		CreateGroupLogicGen: NewCreateGroupLogicGen(ctx, svc),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *socialRpc.CreateGroupReq) (*socialRpc.CreateGroupResp, error) {

	return &socialRpc.CreateGroupResp{}, nil
}
