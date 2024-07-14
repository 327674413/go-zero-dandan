package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"go-zero-dandan/common/resd"
)

type CreateGroupMemberApplyLogic struct {
	*CreateGroupMemberApplyLogicGen
}

func NewCreateGroupMemberApplyLogic(ctx context.Context, svc *svc.ServiceContext) *CreateGroupMemberApplyLogic {
	return &CreateGroupMemberApplyLogic{
		CreateGroupMemberApplyLogicGen: NewCreateGroupMemberApplyLogicGen(ctx, svc),
	}
}

func (l *CreateGroupMemberApplyLogic) CreateGroupMemberApply(req *types.CreateGroupMemberApplyReq) (resp *types.CreateGroupMemberApplyResp, err error) {
	if err = l.initReq(req); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	return
}
