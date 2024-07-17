package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type GetGroupMemberListLogic struct {
	*GetGroupMemberListLogicGen
}

func NewGetGroupMemberListLogic(ctx context.Context, svc *svc.ServiceContext) *GetGroupMemberListLogic {
	return &GetGroupMemberListLogic{
		GetGroupMemberListLogicGen: NewGetGroupMemberListLogicGen(ctx, svc),
	}
}

func (l *GetGroupMemberListLogic) GetGroupMemberList(in *types.GetGroupMemberListReq) (resp *types.GroupMemberListResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return
}
