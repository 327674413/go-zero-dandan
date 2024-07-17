package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
)

type OperateGroupMemberApplyLogic struct {
	*OperateGroupMemberApplyLogicGen
}

func NewOperateGroupMemberApplyLogic(ctx context.Context, svc *svc.ServiceContext) *OperateGroupMemberApplyLogic {
	return &OperateGroupMemberApplyLogic{
		OperateGroupMemberApplyLogicGen: NewOperateGroupMemberApplyLogicGen(ctx, svc),
	}
}

func (l *OperateGroupMemberApplyLogic) OperateGroupMemberApply(in *socialRpc.OperateGroupMemberApplyReq) (*socialRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return &socialRpc.ResultResp{}, nil
}
func (l *OperateGroupMemberApplyLogic) checkReqParams(in *socialRpc.OperateGroupMemberApplyReq) error {
	return nil
}
