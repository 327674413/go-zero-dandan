package logic

import (
	"context"
	"go-zero-dandan/app/goods/model"
	"go-zero-dandan/app/goods/rpc/internal/svc"
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

func (l *GetPageLogic) GetPage(in *goodsRpc.GetPageReq) (*goodsRpc.GetPageResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	page := l.req.Page
	size := l.req.Size
	sort := l.req.Sort
	if size == 0 {
		size = defaultPageSize
	}
	goodsModel := model.NewGoodsMainModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	list, err := goodsModel.Ctx(l.ctx).Page(page, size).Order(sort).CacheSelect(l.svc.Redis)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	page, size = goodsModel.Dao().GetLastQueryPageAndSize()
	isCache := goodsModel.Dao().GetLastQueryIsCache()
	goodsList := make([]*goodsRpc.GoodsInfo, 0)
	for _, item := range list {
		goodsList = append(goodsList, &goodsRpc.GoodsInfo{
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
	return &goodsRpc.GetPageResp{
		Page:    page,
		Size:    size,
		List:    goodsList,
		IsCache: isCache,
	}, nil
}
