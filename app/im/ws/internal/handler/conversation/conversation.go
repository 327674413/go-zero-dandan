package conversation

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/pkg/mapd"
	"go-zero-dandan/pkg/numd"
	"reflect"
	"strconv"
	"time"
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
		if data.ConversationId == "" {
			switch data.ChatType {
			case websocketd.SingleChatType:
				data.ConversationId = numd.CombineInt64(conn.Uid, data.RecvId)
			case websocketd.GroupChatType:
				data.ConversationId = fmt.Sprintf("%d", data.RecvId)
			}
		}
		switch data.ChatType {
		case websocketd.SingleChatType:

		case websocketd.GroupChatType:

		}
		err := svc.MsgChatTransferClient.Push(&kafkad.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			SendTime:       time.Now().UnixNano(),
			MsgType:        data.Msg.MsgType,
			Content:        data.Msg.Content,
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
