package logic

import (
	"context"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/goods/rpc/internal/svc"
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

func (l *GetOneLogic) GetOne(in *goodsRpc.IdReq) (*goodsRpc.GoodsInfo, error) {
	// todo: add your logic here and delete this line

	return &goodsRpc.GoodsInfo{}, nil
}
