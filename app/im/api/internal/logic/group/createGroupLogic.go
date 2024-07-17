package group

import (
	"context"
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
func (l *CreateGroupLogic) CreateGroup(in *types.CreateGroupReq) (resp *types.CreateGroupResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	return
}
