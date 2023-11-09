package goodsInfo

import (
	"context"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"
	"go-zero-dandan/app/goods/rpc/types/pb"
	"go-zero-dandan/common/utild/copier"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type GetHotPageLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       int64
	platClasEm   int64
}

func NewGetHotPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotPageLogic {
	return &GetHotPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotPageLogic) GetHotPage(req *types.GetHotPageReq) (resp *types.GetPageResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	list, err := l.svcCtx.GoodsRpc.GetHotPage(l.ctx, &pb.GetHotPageReq{
		Page:   req.Page,
		Size:   req.Size,
		PlatId: l.platId,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	resp = &types.GetPageResp{}
	err = copier.Copy(&resp, list)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return
}

func (l *GetHotPageLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *GetHotPageLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
