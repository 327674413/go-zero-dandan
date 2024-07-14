package group

import (
	"context"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type GetMyGroupApplyRecvListLogic struct {
	*GetMyGroupApplyRecvListLogicGen
}

func NewGetMyGroupApplyRecvListLogic(ctx context.Context, svc *svc.ServiceContext) *GetMyGroupApplyRecvListLogic {
	return &GetMyGroupApplyRecvListLogic{
		GetMyGroupApplyRecvListLogicGen: NewGetMyGroupApplyRecvListLogicGen(ctx, svc),
	}
}

func (l *GetMyGroupApplyRecvListLogic) GetMyGroupApplyRecvList() (resp *types.GroupApplyListResp, err error) {

	return
}
