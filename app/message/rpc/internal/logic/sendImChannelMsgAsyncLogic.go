package logic

import (
	"context"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
)

type SendImChannelMsgAsyncLogic struct {
	*SendImChannelMsgAsyncLogicGen
}

func NewSendImChannelMsgAsyncLogic(ctx context.Context, svc *svc.ServiceContext) *SendImChannelMsgAsyncLogic {
	return &SendImChannelMsgAsyncLogic{
		SendImChannelMsgAsyncLogicGen: NewSendImChannelMsgAsyncLogicGen(ctx, svc),
	}
}

func (l *SendImChannelMsgAsyncLogic) SendImChannelMsgAsync(req *messageRpc.SendImChannelMsgReq) (*messageRpc.ResultResp, error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}

	return &messageRpc.ResultResp{}, nil
}
