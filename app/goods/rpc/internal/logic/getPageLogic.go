package logic

import (
	"context"

	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageLogic {
	return &GetPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPageLogic) GetPage(in *pb.GetPageReq) (*pb.GetPageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPageResp{}, nil
}
