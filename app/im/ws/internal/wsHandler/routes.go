// Code generated by goctl. DO NOT EDIT.
package wsHandler

import (
	"net/http"

	"go-zero-dandan/app/im/ws/internal/wsSvc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *wsSvc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/connection",
				Handler: connectionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/connection",
				Handler: connectionHandler(serverCtx),
			},
		},
	)
}