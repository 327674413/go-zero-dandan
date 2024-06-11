package handler

import (
	"go-zero-dandan/app/im/ws/internal/handler/conversation"
	"go-zero-dandan/app/im/ws/internal/handler/push"
	"go-zero-dandan/app/im/ws/internal/handler/user"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
)

func RegisterHandlers(server *websocketd.Server, svc *svc.ServiceContext) {
	server.AddRoutes([]websocketd.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "conversation.markChat",
			Handler: conversation.MarkRead(svc),
		},
		{
			Method:  "push",
			Handler: push.Push(svc),
		},
	})
}
