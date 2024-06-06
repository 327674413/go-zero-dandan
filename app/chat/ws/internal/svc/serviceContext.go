package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/chat/mq/mqclient"
	"go-zero-dandan/app/chat/ws/internal/config"
	"go-zero-dandan/app/chat/ws/mgmodel"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/interceptor"
)

type ServiceContext struct {
	Config  config.Config
	Mode    string
	UserRpc user.User
	mgmodel.ChatLogModel
	mqclient.MsgChatTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	return &ServiceContext{
		Config:                c,
		Mode:                  c.Mode,
		UserRpc:               UserRpc,
		ChatLogModel:          mgmodel.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		MsgChatTransferClient: mqclient.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
	}
}
