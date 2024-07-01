package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/modelMongo"
	"go-zero-dandan/app/im/mq/mqClient"
	"go-zero-dandan/app/im/ws/internal/config"
	"go-zero-dandan/app/plat/rpc/plat"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/interceptor"
)

type ServiceContext struct {
	Config  config.Config
	Mode    string
	UserRpc user.User
	PlatRpc plat.Plat
	modelMongo.ChatLogModel
	mqClient.MsgChatTransferClient
	mqClient.MsgReadTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	PlatRpc := plat.NewPlat(zrpc.MustNewClient(c.PlatRpc, zrpc.WithUnaryClientInterceptor(interceptor.RpcClientInterceptor())))
	return &ServiceContext{
		Config:                c,
		Mode:                  c.Mode,
		UserRpc:               UserRpc,
		PlatRpc:               PlatRpc,
		ChatLogModel:          modelMongo.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		MsgChatTransferClient: mqClient.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		MsgReadTransferClient: mqClient.NewMsgReadTransferClient(c.MsgReadTransfer.Addrs, c.MsgReadTransfer.Topic),
	}
}
