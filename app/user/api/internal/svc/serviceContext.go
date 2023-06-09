package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/message/rpc/message"
	"go-zero-dandan/app/user/api/internal/config"
	"go-zero-dandan/app/user/api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	MessageRpc     message.Message
	LangMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		MessageRpc:     message.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
		LangMiddleware: middleware.NewLangMiddleware().Handle,
	}
}
