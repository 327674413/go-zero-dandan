package push

import (
	"github.com/mitchellh/mapstructure"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/constd"
)

func Push(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		var data websocketd.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			server.Send(websocketd.NewErrMessage(err))
			return
		}
		switch data.ChatType {
		case websocketd.ChatTypeSingle:
			single(server, &data, data.RecvId)
		case websocketd.ChatTypeGroup:
			group(server, &data)
		}

	}
}

func single(server *websocketd.Server, data *websocketd.Push, recvId string) error {
	//发送
	rconn := server.GetConn(recvId)
	if rconn == nil {
		server.Info("推送用户id【", recvId, "】离线")
		//离线
		return nil
	}
	//普通聊天消息内容
	if data.MsgClas == constd.MsgClasEmChat {
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
			MsgClas:  data.MsgClas,
			SendId:   data.SendId,
			RecvId:   data.RecvId,
		}), rconn)
	} else {
		return server.Send(websocketd.NewMessage(data.SendId, &websocketd.SysMsg{
			MsgClas:    data.MsgClas,
			MsgType:    data.MsgType,
			MsgContent: data.Content,
			SendTime:   data.SendTime,
		}), rconn)
	}

}
func group(server *websocketd.Server, data *websocketd.Push) error {
	for _, id := range data.RecvIds {
		func(id string) {
			server.Schedule(func() {
				server.Info("推送群聊消息给：", id)
				single(server, data, id)
			})
		}(id)
	}
	return nil
}
