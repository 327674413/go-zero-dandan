package push

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/chat/ws/internal/svc"
	"go-zero-dandan/app/chat/ws/websocketd"
	"go-zero-dandan/pkg/mapd"
	"time"
)

func Push(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		var data websocketd.Push
		logx.Info("11111")
		if err := mapd.AnyToStruct(msg.Data, &data); err != nil {
			server.Send(websocketd.NewErrMessage(err))
			return
		}
		logx.Info(msg.Data)
		//发送
		rconn := server.GetConn(data.RecvId)
		if rconn == nil {
			logx.Info("目标对象离线")
			//离线
			return
		}

		server.Send(websocketd.NewMessage(fmt.Sprintf("%d", data.SendId), &websocketd.Chat{
			ConversationCode: data.ConversationCode,
			SendId:           data.SendId,
			RecvId:           data.RecvId,
			Msg:              websocketd.Msg{Content: data.Content},
			ChatType:         data.ChatType,
			SendTime:         time.Now().UnixNano(),
		}), rconn)
	}
}
