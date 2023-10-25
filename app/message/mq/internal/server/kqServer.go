package server

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-dandan/app/message/mq/internal/config"
	"go-zero-dandan/app/message/mq/internal/logic"
	"go-zero-dandan/app/message/mq/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqConsumerConf, logic.NewSmsSendLogic(ctx, svcCtx)),
		//.....
	}

}
