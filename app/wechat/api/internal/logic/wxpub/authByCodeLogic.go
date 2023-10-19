package wxpub

import (
	"context"
	"fmt"
	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"
	"go-zero-dandan/common/wechat"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type AuthByCodeLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       int64
	platClasEm   int64
}

func NewAuthByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthByCodeLogic {
	return &AuthByCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthByCodeLogic) AuthByCode(req *types.AuthByCodeReq) (resp *types.AuthByCodeResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	if err != nil {
		return nil, err
	}
	fmt.Println("进来的是新的接口")
	wxpub := wechat.NewWxpub(l.ctx, &wechat.WxpubConf{
		Appid:  "wx6ba0f04a081a54e5",               // 小程序、公众号或者企业微信的appid
		Secret: "7e544ae275d779198af0488542ad8ba1", // 商户号 appID
		Token:  "KLJ23KLJKL231JKL312JKL31JKL",
		AESKey: "jczy1uuOSkheNb2oh7V3XvPUaVcof3AYU7fK6hyZOhU",
	}, l.svcCtx.Redis)
	resp = &types.AuthByCodeResp{}
	err = wxpub.AuthByCode(req.Code, &resp)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	ip, err := wxpub.Client.Base.GetCallbackIP(l.ctx)
	fmt.Println("获取了ip", ip)
	l.svcCtx.Redis.Set("test", "测试", "11")

	return resp, nil
}

func (l *AuthByCodeLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *AuthByCodeLogic) initPlat() (err error) {
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
