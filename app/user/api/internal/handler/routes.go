// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	account "go-zero-dandan/app/user/api/internal/handler/account"
	userInfo "go-zero-dandan/app/user/api/internal/handler/userInfo"
	"go-zero-dandan/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MetaMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/regByPhone",
					Handler: account.RegByPhoneHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/loginByPhone",
					Handler: account.LoginByPhoneHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/getPhoneVerifyCode",
					Handler: account.GetPhoneVerifyCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/loginByWxappCode",
					Handler: account.LoginByWxappCodeHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MetaMiddleware, serverCtx.UserInfoMiddleware, serverCtx.UserTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/editMyInfo",
					Handler: userInfo.EditMyInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user/v1"),
	)
}
