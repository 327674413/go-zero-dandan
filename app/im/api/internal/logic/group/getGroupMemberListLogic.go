package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"

	"go-zero-dandan/common/resd"
)

type GetGroupMemberListLogic struct {
	*GetGroupMemberListLogicGen
}

func NewGetGroupMemberListLogic(ctx context.Context, svc *svc.ServiceContext) *GetGroupMemberListLogic {
	return &GetGroupMemberListLogic{
		GetGroupMemberListLogicGen: NewGetGroupMemberListLogicGen(ctx, svc),
	}
}

func (l *GetGroupMemberListLogic) GetGroupMemberList(req *types.GetGroupMemberListReq) (resp *types.GroupMemberListResp, err error) {
	if err = l.initReq(req); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	return
}
