package handler

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/plat/rpc/plat"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/ctxd"
	"net/http"
)

type UserAuth struct {
	svc *svc.ServiceContext
	logx.Logger
	UserRpc user.User
}

func NewUserAuth(svc *svc.ServiceContext) *UserAuth {
	return &UserAuth{
		svc:    svc,
		Logger: logx.WithContext(context.Background()),
	}
}

func (t *UserAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	userToken := r.Header.Get("token")
	if userToken == "" {
		//目前web端没通过url获取，放在这个协议里
		userToken = r.Header.Get("sec-websocket-protocol")
	}
	if userToken == "" {
		logx.Info("ws连接未携带token")
		return false
	}
	//todo::暂时用写死的方式判断是系统层级的token，无需用户，1101代表mq专用
	if userToken == t.svc.Config.Ws.SysToken {
		*r = *r.WithContext(context.WithValue(r.Context(), "userId", "1101"))
	} else {
		userMainInfo, err := t.svc.UserRpc.GetUserByToken(r.Context(), &user.TokenReq{Token: &userToken})
		if err != nil {
			logx.Info("未查询到用户信息", err)
			return false
		}
		platInfo, err := t.svc.PlatRpc.GetOne(r.Context(), &plat.IdReq{Id: &userMainInfo.PlatId})
		if err != nil {
			logx.Info("未查询到应用信息", err)
			return false
		}
		*r = *r.WithContext(context.WithValue(r.Context(), "userId", userMainInfo.Id))
		*r = *r.WithContext(context.WithValue(r.Context(), "platId", userMainInfo.PlatId))
		*r = *r.WithContext(context.WithValue(r.Context(), "platClasEm", platInfo.ClasEm))
	}
	return true
}
func (t *UserAuth) UserId(r *http.Request) string {
	return ctxd.UserId(r.Context())
}
