package logic

import (
	"context"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
)

type SendImChannelMsgLogic struct {
	*SendImChannelMsgLogicGen
}

func NewSendImChannelMsgLogic(ctx context.Context, svc *svc.ServiceContext) *SendImChannelMsgLogic {
	return &SendImChannelMsgLogic{
		SendImChannelMsgLogicGen: NewSendImChannelMsgLogicGen(ctx, svc),
	}
}

func (l *SendImChannelMsgLogic) SendImChannelMsg(in *messageRpc.SendImChannelMsgReq) (*messageRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}

	return &messageRpc.ResultResp{}, nil
}
