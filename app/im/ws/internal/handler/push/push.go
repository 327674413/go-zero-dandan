package push

import (
	"github.com/mitchellh/mapstructure"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/fmtd"
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
	fmtd.Json(data)
	rconn := server.GetConn(recvId)
	if rconn == nil {
		server.Info("推送用户id【", recvId, "】离线")
		//离线
		return nil
	}
	//普通聊天消息内容
	switch data.MsgClas {
	//普通聊天
	case constd.MsgClasEmChat:
		return server.Send(websocketd.NewMessage(data.SendId, &websocketd.Chat{
			ConversationId: data.ConversationId,
			MsgContent:     data.MsgContent,
			MsgType:        data.MsgType,
			Id:             data.Id,
			MsgState:       data.MsgState,
			ChatType:       data.ChatType,
			SendTime:       data.SendTime,
			SendAtMs:       data.SendAtMs,
			MsgClas:        data.MsgClas,
			SendId:         data.SendId,
			RecvId:         data.RecvId,
		}), rconn)
	//消息已读
	case constd.MsgClasEmMarkRead:
		return server.Send(websocketd.NewMessage(data.SendId, &websocketd.Chat{
			ConversationId: data.ConversationId,
			MsgContent:     data.MsgContent,
			MsgType:        data.MsgType,
			Id:             data.Id,
			ChatType:       data.ChatType,
			SendTime:       data.SendTime,
			SendAtMs:       data.SendAtMs,
			MsgClas:        data.MsgClas,
			SendId:         data.SendId,
			RecvId:         data.RecvId,
			ReadRecords:    data.ReadRecords,
		}), rconn)
	default:
		return server.Send(websocketd.NewMessage(data.SendId, &websocketd.SysMsg{
			MsgClas:    data.MsgClas,
			MsgType:    data.MsgType,
			MsgContent: data.MsgContent,
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
