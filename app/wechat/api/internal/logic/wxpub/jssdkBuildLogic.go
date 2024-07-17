package wxpub

import (
	"context"
	"go-zero-dandan/common/wechat"

	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"

	"go-zero-dandan/common/resd"
)

type JssdkBuildLogic struct {
	*JssdkBuildLogicGen
}

func NewJssdkBuildLogic(ctx context.Context, svc *svc.ServiceContext) *JssdkBuildLogic {
	return &JssdkBuildLogic{
		JssdkBuildLogicGen: NewJssdkBuildLogicGen(ctx, svc),
	}
}

func (l *JssdkBuildLogic) JssdkBuild(in *types.JssdkBuildReq) (resp *types.JssdkBuildResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	if !l.hasReq.Url {
		return nil, l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, resd.VarUrl)
	}
	wxpubApp := wechat.NewWxpub(l.ctx, &wechat.WxpubConf{
		Appid:  "wx6ba0f04a081a54e5",
		Secret: "7e544ae275d779198af0488542ad8ba1",
		Token:  "KLJ23KLJKL231JKL312JKL31JKL",
		AESKey: "jczy1uuOSkheNb2oh7V3XvPUaVcof3AYU7fK6hyZOhU",
	}, nil)
	jssdkRes, err := wxpubApp.JssdkBuild(&wechat.JssdkBuildParams{
		Url:         l.req.Url,
		JsApiList:   l.req.JsApiList,
		OpenTagList: l.req.OpenTagList,
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	resp = &types.JssdkBuildResp{
		AppId:       jssdkRes.AppId,
		Beta:        jssdkRes.Beta,
		Debug:       jssdkRes.Debug,
		JsApiList:   jssdkRes.JsApiList,
		NonceStr:    jssdkRes.NonceStr,
		OpenTagList: jssdkRes.OpenTagList,
		Signature:   jssdkRes.Signature,
		Timestamp:   jssdkRes.Timestamp,
		Url:         jssdkRes.Url,
	}
	return resp, nil
}
