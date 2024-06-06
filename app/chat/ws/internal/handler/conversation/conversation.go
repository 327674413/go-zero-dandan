package conversation

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/chat/mq/kafkad"
	"go-zero-dandan/app/chat/ws/internal/svc"
	"go-zero-dandan/app/chat/ws/websocketd"
	"go-zero-dandan/pkg/mapd"
	"go-zero-dandan/pkg/numd"
	"reflect"
	"strconv"
)

func stringToInt64HookFunc(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String && t.Kind() == reflect.Int64 {
		return strconv.ParseInt(data.(string), 10, 64)
	}
	return data, nil
}

func Chat(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		var data websocketd.Chat
		if err := mapd.AnyToStruct(msg.Data, &data); err != nil {
			server.Send(websocketd.NewErrMessage(fmt.Errorf("解析消息失败：%v", err)), conn)
			return
		}
		if data.ConversationCode == "" {
			switch data.ChatType {
			case websocketd.SingleChatType:
				data.ConversationCode = numd.CombineInt64(conn.Uid, data.RecvId)
			case websocketd.GroupChatType:
				data.ConversationCode = fmt.Sprintf("%d", data.RecvId)
			}
		}
		logx.Info("Chat触发：", data)
		switch data.ChatType {
		case websocketd.SingleChatType:
			err := svc.Push(&kafkad.MsgChatTransfer{
				ConversationCode: data.ConversationCode,
				ChatType:         data.ChatType,
				SendId:           data.SendId,
				RecvId:           data.RecvId,
				SendTime:         data.SendTime,
				MsgType:          data.MsgType,
				Content:          data.Content,
			})
			if err != nil {
				server.Send(websocketd.NewErrMessage(err), conn)
				return
			}

		}
	}
}
