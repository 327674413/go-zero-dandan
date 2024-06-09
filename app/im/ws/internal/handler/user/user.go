package user

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
)

func OnLine(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		uids := server.GetUsers()
		u := server.GetUsers(conn)
		err := server.Send(websocketd.NewMessage(fmt.Sprintf("%d", u[0]), uids), conn)
		logx.Info("err", err)
	}
}
