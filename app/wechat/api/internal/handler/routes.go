// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	wxpub "go-zero-dandan/app/wechat/api/internal/handler/wxpub"
	"go-zero-dandan/app/wechat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LangMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/service",
					Handler: wxpub.ServiceHandler(serverCtx),
				},
			}...,
		),
		rest.WithTimeout(30000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LangMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/jssdkBuild",
					Handler: wxpub.JssdkBuildHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authByCode",
					Handler: wxpub.AuthByCodeHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithTimeout(30000*time.Millisecond),
	)
}