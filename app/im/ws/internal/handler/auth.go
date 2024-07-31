package handler

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/plat/rpc/plat"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/ctxd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
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

func (t *UserAuth) Auth(w http.ResponseWriter, r *http.Request) error {
	userToken := r.Header.Get("token")
	if userToken == "" {
		//目前web端没通过url获取，放在这个协议里
		userToken = r.Header.Get("sec-websocket-protocol")
	}
	resp := resd.NewResp(r.Context(), utild.GetRequestLang(r))
	if userToken == "" {
		return resp.NewErrWithTemp(resd.ErrReqFieldEmpty1, "token")
	}
	//todo::暂时用写死的方式判断是系统层级的token，无需用户，1101代表mq专用
	if userToken == t.svc.Config.Ws.SysToken {
		*r = *r.WithContext(context.WithValue(r.Context(), "userId", "1101"))
	} else {
		userMainInfo, err := t.svc.UserRpc.GetUserByToken(r.Context(), &user.TokenReq{Token: &userToken})
		if err != nil {
			if !resd.IsUserNotLoginErr(err) {
				return resp.NewErr(resd.ErrSys)
			} else {
				return resp.NewErr(resd.ErrAuthUserNotLogin)
			}
		}
		platInfo, err := t.svc.PlatRpc.GetOne(r.Context(), &plat.IdReq{Id: &userMainInfo.PlatId})
		if err != nil {
			return resp.Error(err)
		}
		*r = *r.WithContext(context.WithValue(r.Context(), "userId", userMainInfo.Id))
		*r = *r.WithContext(context.WithValue(r.Context(), "platId", userMainInfo.PlatId))
		*r = *r.WithContext(context.WithValue(r.Context(), "platClasEm", platInfo.ClasEm))
	}
	return nil
}
func (t *UserAuth) UserId(r *http.Request) string {
	return ctxd.UserId(r.Context())
}
