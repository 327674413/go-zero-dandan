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
			Method:  "user.online", //获取在线用户（暂时好像无用）
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat", //消息对话，私聊、群聊等
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "conversation.markChat", //标记消息已读
			Handler: conversation.MarkRead(svc),
		},
		{
			Method:  "push", //专门用于mq转发消息用的，比如发送消息不会直接发出去，而是先进mq
			Handler: push.Push(svc),
		},
	})
}
