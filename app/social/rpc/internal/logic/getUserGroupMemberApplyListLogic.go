package logic

import (
	"context"
	"go-zero-dandan/app/social/rpc/internal/svc"
	"go-zero-dandan/app/social/rpc/types/socialRpc"
)

type GetUserGroupMemberApplyListLogic struct {
	*GetUserGroupMemberApplyListLogicGen
}

func NewGetUserGroupMemberApplyListLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserGroupMemberApplyListLogic {
	return &GetUserGroupMemberApplyListLogic{
		GetUserGroupMemberApplyListLogicGen: NewGetUserGroupMemberApplyListLogicGen(ctx, svc),
	}
}

func (l *GetUserGroupMemberApplyListLogic) GetUserGroupMemberApplyList(in *socialRpc.GetUserGroupMemberApplyListReq) (*socialRpc.GroupMemberApplyListResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return &socialRpc.GroupMemberApplyListResp{}, nil
}
func (l *GetUserGroupMemberApplyListLogic) checkReqParams(in *socialRpc.GetUserGroupMemberApplyListReq) error {
	return nil
}
