package account

import (
	"context"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type LoginByWxappCodeLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       int64
	platClasEm   int64
}

func NewLoginByWxappCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByWxappCodeLogic {
	return &LoginByWxappCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByWxappCodeLogic) LoginByWxappCode(req *types.LoginByWxappCodeReq) (resp *types.LoginByWxappCodeResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	resp = &types.LoginByWxappCodeResp{
		UserInfo:      &types.UserInfoResp{},
		WxappUserInfo: &types.WxappUserInfoResp{},
	}
	return resp, nil
}

func (l *LoginByWxappCodeLogic) initPlat() (err error) {
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
