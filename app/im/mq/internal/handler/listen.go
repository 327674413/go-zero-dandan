package handler

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-dandan/app/im/mq/internal/handler/msgTransfer"
	"go-zero-dandan/app/im/mq/internal/svc"
)

type Listen struct {
	svc *svc.ServiceContext
}

func NewListen(svc *svc.ServiceContext) *Listen {
	return &Listen{
		svc: svc,
	}
}
func (t *Listen) Services() []service.Service {
	return []service.Service{
		//这里可以加载多个消费者
		kq.MustNewQueue(t.svc.Config.MsgChatTransfer, msgTransfer.NewMsgChatTransfer(t.svc)),
		kq.MustNewQueue(t.svc.Config.MsgReadTransfer, msgTransfer.NewMsgReadTransfer(t.svc)),
	}
}
