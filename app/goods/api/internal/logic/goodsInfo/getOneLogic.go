package goodsInfo

import (
	"context"

	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type GetOneLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneLogic) GetOne(req *types.IdReq) (resp *types.GoodsInfo, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if req.Id == nil {
		return nil, resd.NewErrWithTempCtx(l.ctx, "缺少参数id", resd.ReqFieldRequired1, "Id")
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

func (l *GetOneLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *GetOneLogic) initPlat() (err error) {
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
