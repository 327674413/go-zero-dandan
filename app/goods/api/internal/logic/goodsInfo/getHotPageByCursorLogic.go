package goodsInfo

import (
	"context"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"
)

type GetHotPageByCursorLogic struct {
	*GetHotPageByCursorLogicGen
}

func NewGetHotPageByCursorLogic(ctx context.Context, svc *svc.ServiceContext) *GetHotPageByCursorLogic {
	return &GetHotPageByCursorLogic{
		GetHotPageByCursorLogicGen: NewGetHotPageByCursorLogicGen(ctx, svc),
	}
}

func (l *GetHotPageByCursorLogic) GetHotPageByCursor(in *types.GetHotPageByCursorReq) (resp *types.GetHotPageByCursorResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	pageData, err := l.svc.GoodsRpc.GetHotPageByCursor(l.ctx, &goodsRpc.GetHotPageByCursorReq{
		Size:   &l.req.Size,
		Page:   &l.req.Page,
		Cursor: &l.req.Cursor,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	goodses := make([]*types.GoodsInfo, 0)
	for _, item := range pageData.List {
		goodses = append(goodses, &types.GoodsInfo{
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
			ViewNum:   item.ViewNum,
		})
	}
	return &types.GetHotPageByCursorResp{
		IsCache: pageData.IsCache,
		IsEnd:   pageData.IsEnd,
		Cursor:  pageData.Cursor,
		LastId:  pageData.LastId,
		Size:    pageData.Size,
		List:    goodses,
	}, nil
}
