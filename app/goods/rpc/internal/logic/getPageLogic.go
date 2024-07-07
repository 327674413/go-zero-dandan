package logic

import (
	"context"
	"go-zero-dandan/app/goods/model"
	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"
	"go-zero-dandan/common/resd"

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

func (l *GetPageLogic) GetPage(in *goodsRpc.GetPageReq) (*goodsRpc.GetPageResp, error) {
	page := in.Page
	size := in.Size
	sort := in.Sort
	if size == 0 {
		size = defaultPageSize
	}
	goodsModel := model.NewGoodsMainModel(l.ctx, l.svcCtx.SqlConn, in.PlatId)
	list, err := goodsModel.Ctx(l.ctx).Page(page, size).Order(sort).CacheSelect(l.svcCtx.Redis)
	if err != nil {
		return nil, resd.RpcErrEncode(resd.ErrorCtx(l.ctx, err))
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
