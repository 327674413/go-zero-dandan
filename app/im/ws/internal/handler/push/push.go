package push

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/pkg/mapd"
)

func Push(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		var data websocketd.Push
		if err := mapd.AnyToStruct(msg.Data, &data); err != nil {
			server.Send(websocketd.NewErrMessage(err))
			return
		}
		switch data.ChatType {
		case websocketd.SingleChatType:
			single(server, &data, data.RecvId)
		case websocketd.GroupChatType:
			group(server, &data)
		}

	}
}
func single(server *websocketd.Server, data *websocketd.Push, recvId string) error {
	//发送
	rconn := server.GetConn(recvId)
	if rconn == nil {
		logx.Info("推送目标对象", recvId, "离线")
		//离线
		return nil
	}

	return server.Send(websocketd.NewMessage(data.SendId, &websocketd.Chat{
		ConversationId: data.ConversationId,
		Msg: websocketd.Msg{
			Content:     data.Content,
			MsgType:     data.MsgType,
			MsgId:       data.MsgId,
			ReadRecords: data.ReadRecords,
		},
		ChatType: data.ChatType,
		SendTime: data.SendTime,
	}), rconn)
}
func group(server *websocketd.Server, data *websocketd.Push) error {
	for _, id := range data.RecvIds {
		func(id string) {
			server.Schedule(func() {
				logx.Info("推送群聊消息给：", id)
				single(server, data, id)
			})
		}(id)
	}
	return nil
}
