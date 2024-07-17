package goodsInfo

import (
	"context"

	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
)

type GetOneLogic struct {
	*GetOneLogicGen
}

func NewGetOneLogic(ctx context.Context, svc *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		GetOneLogicGen: NewGetOneLogicGen(ctx, svc),
	}
}

func (l *GetOneLogic) GetOne(in *types.IdReq) (resp *types.GoodsInfo, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	resp = &types.GoodsInfo{
		Id:        "",
		Name:      "",
		Spec:      "",
		Cover:     "",
		SellPrice: 0,
		StoreQty:  0,
		State:     0,
		IsSpecial: 0,
		UnitId:    "",
		UnitName:  "",
		PlatId:    "",
	}
	return resp, nil
}
