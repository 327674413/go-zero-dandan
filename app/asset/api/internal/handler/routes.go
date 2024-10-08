// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"go-zero-dandan/app/asset/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MetaMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/uploadImg",
					Handler: UploadImgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: UploadHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithTimeout(30000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MetaMiddleware, serverCtx.UserInfoMiddleware, serverCtx.UserTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/download",
					Handler: DownloadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multipartUpload/init",
					Handler: MultipartUploadInitHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multipartUpload/send",
					Handler: MultipartUploadSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multipartUpload/complete",
					Handler: MultipartUploadCompleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithTimeout(999000*time.Millisecond),
	)
}
