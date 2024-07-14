package wxpub

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/zeromicro/go-zero/core/logc"
	"io"
	"net/http"

	"go-zero-dandan/app/wechat/api/internal/svc"
	"go-zero-dandan/app/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type ServiceLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId       int64
	platClasEm   int64
}

func NewServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServiceLogic {
	return &ServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceLogic) Service(req *types.WxpubReq, r *http.Request) (resp []byte, err error) {
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
		logc.Error(l.ctx, err)
		return nil, err
	}
	text, _ := io.ReadAll(rs.Body)
	fmt.Println(string(text))
	return text, nil
}

func (l *ServiceLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.ErrUserMainInfo)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *ServiceLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.ErrPlatClas)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.ErrPlatId)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
