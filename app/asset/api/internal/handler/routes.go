// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-dandan/app/asset/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LangMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/uploadImg",
					Handler: UploadImgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/uploadFile",
					Handler: UploadFileHandler(serverCtx),
				},
			}...,
		),
	)
}