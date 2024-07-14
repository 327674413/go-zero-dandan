package wxpub

import (
	"context"
	"fmt"
	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"
	"go-zero-dandan/common/wechat"

	"go-zero-dandan/common/resd"
)

type AuthByCodeLogic struct {
	*AuthByCodeLogicGen
}

func NewAuthByCodeLogic(ctx context.Context, svc *svc.ServiceContext) *AuthByCodeLogic {
	return &AuthByCodeLogic{
		AuthByCodeLogicGen: NewAuthByCodeLogicGen(ctx, svc),
	}
}

func (l *AuthByCodeLogic) AuthByCode(req *types.AuthByCodeReq) (resp *types.AuthByCodeResp, err error) {
	if err = l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
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
	}, l.svc.Redis)
	resp = &types.AuthByCodeResp{}
	err = wxpub.AuthByCode(*req.Code, &resp)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	ip, err := wxpub.Client.Base.GetCallbackIP(l.ctx)
	fmt.Println("获取了ip", ip)
	l.svc.Redis.Set("test", "测试", "11")

	return resp, nil
}
