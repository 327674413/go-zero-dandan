package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-dandan/app/message/mq/internal/svc"
)

type SmsSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmsSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsSendLogic {
	return &SmsSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *SmsSendLogic) Consume(key, val string) error {
	fmt.Printf("get key: %s val: %s\n", key, val)
	return nil
}
func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewSmsSendLogic(ctx, svcCtx)),
	}
}
