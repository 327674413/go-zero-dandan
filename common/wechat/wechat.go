package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"go-zero-dandan/common/redisd"
	"go-zero-dandan/common/resd"
)

type Wxpub struct {
	Appid      string
	Secret     string
	Token      string
	AESKey     string
	jssdkBuild JssdkBuildResult
	Redis      *PowerWechatCache
	ctx        context.Context
	Client     *officialAccount.OfficialAccount
}

func NewWxpub(ctx context.Context, conf *WxpubConf, redis *redisd.Redisd) *Wxpub {
	wxpub := &Wxpub{
		Appid:  conf.Appid,
		Secret: conf.Secret,
		Token:  conf.Token,
		AESKey: conf.AESKey,
	}
	if redis != nil {
		wxpub.Redis = NewPowerWechatCache(redis)
	}
	if ctx != nil {
		wxpub.ctx = ctx
	} else {
		wxpub.ctx = context.Background()
	}
	return wxpub
}
func (t *Wxpub) AuthByCode(code string, target any) error {
	fmt.Println("AuthByCode进来了")
	conf := &officialAccount.UserConfig{
		AppID:  "wx6ba0f04a081a54e5",               // 小程序、公众号或者企业微信的appid
		Secret: "7e544ae275d779198af0488542ad8ba1", // 商户号 appID

		Token:  "KLJ23KLJKL231JKL312JKL31JKL",
		AESKey: "jczy1uuOSkheNb2oh7V3XvPUaVcof3AYU7fK6hyZOhU",
		// ResponseType: os.Getenv("response_type"),
		/*Log: officialAccount.Log{
			Level: "debug",
			File:  "./wechat.log",
		},*/
		HttpDebug: true,
		Debug:     false,
	}
	if t.Redis != nil {
		fmt.Println("redis接管执行到了")
		conf.Cache = t.Redis
	}
	officialAccountApp, err := officialAccount.NewOfficialAccount(conf)
	if err != nil {
		return resd.ErrorCtx(t.ctx, err)
	}
	t.Client = officialAccountApp
	tokenResponse, err := officialAccountApp.OAuth.TokenFromCode(code)
	if err != nil {
		return resd.ErrorCtx(t.ctx, err)
	}
	j, err := json.Marshal(tokenResponse)
	if err != nil {
		return resd.ErrorCtx(t.ctx, err)
	}
	return json.Unmarshal(j, target)
}
func (t *Wxpub) JssdkBuild(params *JssdkBuildParams) (res *JssdkBuildResult, err error) {
	conf := &officialAccount.UserConfig{
		AppID:  "wx6ba0f04a081a54e5",               // 小程序、公众号或者企业微信的appid
		Secret: "7e544ae275d779198af0488542ad8ba1", // 商户号 appID

		Token:  "KLJ23KLJKL231JKL312JKL31JKL",
		AESKey: "jczy1uuOSkheNb2oh7V3XvPUaVcof3AYU7fK6hyZOhU",
		// ResponseType: os.Getenv("response_type"),
		/*Log: officialAccount.Log{
			Level: "debug",
			File:  "./wechat.log",
		},*/
		HttpDebug: true,
		Debug:     false,
	}
	if t.Redis != nil {
		fmt.Println("redis接管执行到了")
		conf.Cache = t.Redis
	}
	OfficialAccountApp, err := officialAccount.NewOfficialAccount(conf)
	if err != nil {
		return nil, err
	}
	debug := false
	beta := false
	result, err := OfficialAccountApp.JSSDK.BuildConfig(context.Background(), params.JsApiList, debug, beta, params.OpenTagList, params.Url)
	if err != nil {
		return nil, err
	}
	res = &JssdkBuildResult{}
	h, ok := result.(*power.HashMap)
	if ok {
		/*
			res.AppId, _ = (*h)["appId"].(string)
			res.Beta, _ = (*h)["beta"].(bool)
			res.Url, _ = (*h)["url"].(string)
			res.JsApiList, _ = (*h)["jsApiList"].([]string)
			res.OpenTagList, _ = (*h)["openTagList"].([]string)
			res.Signature, _ = (*h)["signature"].(string)
			res.NonceStr, _ = (*h)["nonceStr"].(string)
			res.Timestamp, _ = (*h)["timestamp"].(int64)
		*/
		jsonStr, err := json.Marshal(h)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonStr, &res)
		if err != nil {
			return nil, err
		}

	} else {
		fmt.Println("未定位到类型")
	}
	fmt.Println("获取数据成功2")
	fmt.Println(res)
	return res, nil
}
