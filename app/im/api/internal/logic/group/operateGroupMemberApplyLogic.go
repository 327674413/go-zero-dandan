package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type OperateGroupMemberApplyLogic struct {
	*OperateGroupMemberApplyLogicGen
}

func NewOperateGroupMemberApplyLogic(ctx context.Context, svc *svc.ServiceContext) *OperateGroupMemberApplyLogic {
	return &OperateGroupMemberApplyLogic{
		OperateGroupMemberApplyLogicGen: NewOperateGroupMemberApplyLogicGen(ctx, svc),
	}
}

func (l *OperateGroupMemberApplyLogic) OperateGroupMemberApply(in *types.OperateGroupMemberApplyReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return
}
