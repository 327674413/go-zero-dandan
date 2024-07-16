package goodsInfo

import (
	"context"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"
)

type GetPageLogic struct {
	*GetPageLogicGen
}

func NewGetPageLogic(ctx context.Context, svc *svc.ServiceContext) *GetPageLogic {
	return &GetPageLogic{
		GetPageLogicGen: NewGetPageLogicGen(ctx, svc),
	}
}

func (l *GetPageLogic) GetPage(req *types.GetPageReq) (resp *types.GetPageResp, err error) {
	if err = l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	list, err := l.svc.GoodsRpc.GetPage(l.ctx, &goodsRpc.GetPageReq{
		Page: &l.req.Page,
		Size: &l.req.Size,
		Sort: &l.req.Sort,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	goodsList := make([]*types.GoodsInfo, 0)
	for _, item := range list.List {
		goodsList = append(goodsList, &types.GoodsInfo{
			Id:        item.Id,
			Name:      item.Name,
			Spec:      item.Spec,
			Cover:     item.Cover,
			SellPrice: item.SellPrice,
			StoreQty:  item.StoreQty,
			State:     item.State,
			IsSpecial: item.IsSpecial,
			UnitId:    item.UnitId,
			UnitName:  item.UnitName,
			PlatId:    item.PlatId,
		})
	}
	resp = &types.GetPageResp{
		Page:    list.Page,
		Size:    list.Size,
		IsCache: list.IsCache,
		List:    goodsList,
	}
	return resp, nil
}
