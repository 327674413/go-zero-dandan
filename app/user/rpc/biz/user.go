package biz

import (
	"context"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type UserBiz struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewUserBiz(ctx context.Context, svcCtx *svc.ServiceContext) *UserBiz {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	biz := &UserBiz{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
	biz.initPlat()
	return biz
}

func (t *UserBiz) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(t.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(t.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(t.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(t.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	t.platId = platClasId
	t.platClasEm = platClasEm
	return nil
}
