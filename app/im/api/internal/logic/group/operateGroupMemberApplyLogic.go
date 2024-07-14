package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"go-zero-dandan/common/resd"
)

type OperateGroupMemberApplyLogic struct {
	*OperateGroupMemberApplyLogicGen
}

func NewOperateGroupMemberApplyLogic(ctx context.Context, svc *svc.ServiceContext) *OperateGroupMemberApplyLogic {
	return &OperateGroupMemberApplyLogic{
		OperateGroupMemberApplyLogicGen: NewOperateGroupMemberApplyLogicGen(ctx, svc),
	}
}

func (l *OperateGroupMemberApplyLogic) OperateGroupMemberApply(req *types.OperateGroupMemberApplyReq) (resp *types.ResultResp, err error) {
	if err = l.initReq(req); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	return
}
