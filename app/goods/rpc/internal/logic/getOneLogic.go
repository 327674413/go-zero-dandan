package logic

import (
	"context"

	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneLogic) GetOne(in *pb.IdReq) (*pb.GoodsInfo, error) {
	// todo: add your logic here and delete this line

	return &pb.GoodsInfo{}, nil
}
