package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
)

type CreateGroupMemberApplyLogic struct {
	*CreateGroupMemberApplyLogicGen
}

func NewCreateGroupMemberApplyLogic(ctx context.Context, svc *svc.ServiceContext) *CreateGroupMemberApplyLogic {
	return &CreateGroupMemberApplyLogic{
		CreateGroupMemberApplyLogicGen: NewCreateGroupMemberApplyLogicGen(ctx, svc),
	}
}

func (l *CreateGroupMemberApplyLogic) CreateGroupMemberApply(in *socialRpc.CreateGroupMemberApplyReq) (*socialRpc.CreateGroupMemberApplyResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	return &socialRpc.CreateGroupMemberApplyResp{}, nil
}
