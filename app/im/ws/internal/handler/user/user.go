package user

import (
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
)

// OnLine 获取所有在线用户，目前好像没啥用
func OnLine(svc *svc.ServiceContext) websocketd.HandlerFunc {
	return func(server *websocketd.Server, conn *websocketd.Conn, msg *websocketd.Message) {
		uids := server.GetUsers()
		u := server.GetUsers(conn)
		err := server.Send(websocketd.NewMessage(u[0], uids), conn)
		if err != nil {
			server.Error(err)
		}
	}
}
