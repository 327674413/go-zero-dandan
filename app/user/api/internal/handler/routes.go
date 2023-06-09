// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-dandan/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LangMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/LoginByPhone",
					Handler: LoginByPhoneHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/getPhoneVerifyCode",
					Handler: getPhoneVerifyCodeHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
