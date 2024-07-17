package wxpub

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"io"
	"net/http"

	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"
)

type ServiceLogic struct {
	*ServiceLogicGen
}

func NewServiceLogic(ctx context.Context, svc *svc.ServiceContext) *ServiceLogic {
	return &ServiceLogic{
		ServiceLogicGen: NewServiceLogicGen(ctx, svc),
	}
}

func (l *ServiceLogic) Service(in *types.WxpubReq, r *http.Request) (resp []byte, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	officialAccountApp, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:  "wx6ba0f04a081a54e5",               // 小程序、公众号或者企业微信的appid
		Secret: "7e544ae275d779198af0488542ad8ba1", // 商户号 appID

		Token:  "KLJ23KLJKL231JKL312JKL31JKL",
		AESKey: "jczy1uuOSkheNb2oh7V3XvPUaVcof3AYU7fK6hyZOhU",
		// ResponseType: os.Getenv("response_type"),
		/*Log: officialAccount.Log{
			Level: "debug",
			File:  "./wechat.log",
		},*/
		//Cache:     cache,
		HttpDebug: true,
		Debug:     false,
	})
	rs, err := officialAccountApp.Server.VerifyURL(r)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	text, _ := io.ReadAll(rs.Body)
	fmt.Println(string(text))
	return text, nil
}
