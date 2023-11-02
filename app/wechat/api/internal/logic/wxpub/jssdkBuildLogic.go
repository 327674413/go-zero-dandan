package wxpub

import (
	"context"
	"go-zero-dandan/common/wechat"

	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type JssdkBuildLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       int64
	platClasEm   int64
}

func NewJssdkBuildLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JssdkBuildLogic {
	return &JssdkBuildLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JssdkBuildLogic) JssdkBuild(req *types.JssdkBuildReq) (resp *types.JssdkBuildResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if req.Url == nil {
		return nil, resd.NewErrWithTempCtx(l.ctx, "缺少url", resd.ReqFieldRequired1, "url")
	}
	wxpubApp := wechat.NewWxpub(l.ctx, &wechat.WxpubConf{
		Appid:  "wx6ba0f04a081a54e5",
		Secret: "7e544ae275d779198af0488542ad8ba1",
		Token:  "KLJ23KLJKL231JKL312JKL31JKL",
		AESKey: "jczy1uuOSkheNb2oh7V3XvPUaVcof3AYU7fK6hyZOhU",
	}, nil)
	jssdkRes, err := wxpubApp.JssdkBuild(&wechat.JssdkBuildParams{
		Url:         *req.Url,
		JsApiList:   req.JsApiList,
		OpenTagList: req.OpenTagList,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
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

func (l *JssdkBuildLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *JssdkBuildLogic) initPlat() (err error) {
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
