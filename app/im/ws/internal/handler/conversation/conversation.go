package conversation

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/pkg/mapd"
)

// Chat 聊天核心方法，包括私聊、群聊等
func Chat(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		var data websocketd.Chat
		// 将读取的消息结构体，转化为chat聊天消息的结构体
		if err := mapd.AnyToStruct(msg.Data, &data); err != nil {
			server.Send(websocketd.NewErrMessage(fmt.Errorf("解析消息失败：%v", err)), conn)
			return
		}
		// 当会话id不存在时进行处理
		if data.ConversationId == "" {
			switch data.ChatType {
			case websocketd.ChatTypeSingle: //私聊，将两个用户的id组装成会话id
				data.ConversationId = utild.CombineId(conn.Uid, data.RecvId)
			case websocketd.ChatTypeGroup: //群聊，recivId为群id,作为会话id
				data.ConversationId = data.RecvId
			}
		}
		// 根据不同消息类型，做不同的二次处理
		switch data.ChatType {
		case websocketd.ChatTypeSingle:

		case websocketd.ChatTypeGroup:

		}
		logx.Info("推送消息到kafka")
		// 将消息转化为mq消息的格式，并发送的mq
		err := svc.MsgChatTransferClient.Push(&kafkad.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			SendTime:       utild.Date("Y-m-d H:i:s"),
			MsgType:        data.Msg.MsgType,
			Content:        data.Msg.Content,
			MsgId:          data.MsgId,
			TempId:         data.TempId,
		})
		if err != nil {
			server.Send(websocketd.NewErrMessage(err), conn)
			return
		}
	}
}
func MarkRead(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		// todo: 已读未读处理
		var data websocketd.MarkRead
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			server.Send(websocketd.NewErrMessage(err), conn)
			return
		}

		err := svc.MsgReadTransferClient.Push(&kafkad.MsgMarkRead{
			ChatType:       data.ChatType,
			ConversationId: data.ConversationId,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MsgIds:         data.MsgIds,
		})

		if err != nil {
			server.Send(websocketd.NewErrMessage(err), conn)
			return
		}
	}
}
