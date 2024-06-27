package logic

import (
	"context"
	"go-zero-dandan/app/user/rpc/types/userRpc"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/internal/svc"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *userRpc.IdReq) (*userRpc.UserMainInfo, error) {
	// todo: add your logic here and delete this line

	return &userRpc.UserMainInfo{}, nil
}
