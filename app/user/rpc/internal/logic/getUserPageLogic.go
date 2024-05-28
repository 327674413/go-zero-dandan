package logic

import (
	"context"

	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageLogic {
	return &GetUserPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPageLogic) GetUserPage(in *pb.GetUserPageReq) (*pb.GetUserPageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserPageResp{}, nil
}
