package logic

import (
	"context"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"

	"go-zero-dandan/app/goods/rpc/internal/svc"
)

type GetOneLogic struct {
	*GetOneLogicGen
}

func NewGetOneLogic(ctx context.Context, svc *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		GetOneLogicGen: NewGetOneLogicGen(ctx, svc),
	}
}

func (l *GetOneLogic) GetOne(in *goodsRpc.IdReq) (*goodsRpc.GoodsInfo, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return &goodsRpc.GoodsInfo{}, nil
}
