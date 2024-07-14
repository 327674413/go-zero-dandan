package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type GetMyGroupListLogic struct {
	*GetMyGroupListLogicGen
}

func NewGetMyGroupListLogic(ctx context.Context, svc *svc.ServiceContext) *GetMyGroupListLogic {
	return &GetMyGroupListLogic{
		GetMyGroupListLogicGen: NewGetMyGroupListLogicGen(ctx, svc),
	}
}

func (l *GetMyGroupListLogic) GetMyGroupList() (resp *types.GroupListResp, err error) {

	return
}
