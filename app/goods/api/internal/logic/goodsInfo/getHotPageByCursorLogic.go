package goodsInfo

import (
	"context"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type GetHotPageByCursorLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
}

func NewGetHotPageByCursorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotPageByCursorLogic {
	return &GetHotPageByCursorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotPageByCursorLogic) GetHotPageByCursor(req *types.GetHotPageByCursorReq) (resp *types.GetHotPageByCursorResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	pageData, err := l.svcCtx.GoodsRpc.GetHotPageByCursor(l.ctx, &goodsRpc.GetHotPageByCursorReq{
		Size:   req.Size,
		PlatId: l.platId,
		Page:   req.Page,
		Cursor: req.Cursor,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
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

func (l *GetHotPageByCursorLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *GetHotPageByCursorLogic) initPlat() (err error) {
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
