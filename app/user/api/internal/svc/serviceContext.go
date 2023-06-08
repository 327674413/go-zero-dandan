package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/app/user/api/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	MessageRpc message.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MessageRpc: message.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
	}
}
