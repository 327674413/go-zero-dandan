package group

import (
	"context"
	"go-zero-dandan/common/resd"

	"go-zero-dandan/app/im/api/internal/svc"
	"go-zero-dandan/app/im/api/internal/types"
)

type CreateGroupLogic struct {
	*CreateGroupLogicGen
}

func NewCreateGroupLogic(ctx context.Context, svc *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		CreateGroupLogicGen: NewCreateGroupLogicGen(ctx, svc),
	}
}
func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupReq) (resp *types.CreateGroupResp, err error) {
	if err = l.initReq(req); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return
}
