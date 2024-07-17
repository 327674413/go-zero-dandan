package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type CreateGroupMemberApplyLogic struct {
	*CreateGroupMemberApplyLogicGen
}

func NewCreateGroupMemberApplyLogic(ctx context.Context, svc *svc.ServiceContext) *CreateGroupMemberApplyLogic {
	return &CreateGroupMemberApplyLogic{
		CreateGroupMemberApplyLogicGen: NewCreateGroupMemberApplyLogicGen(ctx, svc),
	}
}

func (l *CreateGroupMemberApplyLogic) CreateGroupMemberApply(in *types.CreateGroupMemberApplyReq) (resp *types.CreateGroupMemberApplyResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return
}
