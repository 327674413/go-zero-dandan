package goodsInfo

import (
	"context"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
	"go-zero-dandan/app/goods/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type GetPageLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
}

func NewGetPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageLogic {
	return &GetPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPageLogic) GetPage(req *types.GetPageReq) (resp *types.GetPageResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	list, err := l.svcCtx.GoodsRpc.GetPage(l.ctx, &pb.GetPageReq{
		Page:   req.Page,
		Size:   req.Size,
		Sort:   req.Sort,
		PlatId: l.platId,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
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

func (l *GetPageLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *GetPageLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platId, _ := l.ctx.Value("platId").(string)
	if platId == "" {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platId
	l.platClasEm = platClasEm
	return nil
}
