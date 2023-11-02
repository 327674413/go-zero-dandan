package logic

import (
	"context"

	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPageWithTotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPageWithTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageWithTotalLogic {
	return &GetPageWithTotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPageWithTotalLogic) GetPageWithTotal(in *pb.GetPageReq) (*pb.GetPageWithTotalResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPageWithTotalResp{}, nil
}
