package logic

import (
	"context"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/app/message/mq/mqClient"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendImLogic struct {
	ctx context.Context
	svc *svc.ServiceContext
	logx.Logger
}

func NewSendImLogic(ctx context.Context, svc *svc.ServiceContext) *SendImLogic {
	return &SendImLogic{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendImLogic) SendIm(in *messageRpc.SendImReq) (*messageRpc.SuccResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	err := l.svc.ImSendCli.Push(&mqClient.ImSendMsg{
		ChatType:    websocketd.ChatType(in.ChatType),
		ChannelId:   in.ChannelId,
		SendId:      in.SendUid,
		RecvId:      in.RecvUid,
		RecvIds:     in.RecvUids,
		SendAt:      in.SendAt,
		MsgType:     websocketd.MsgType(in.MsgType),
		Content:     in.Content,
		PlatId:      in.PlatId,
		MsgId:       in.MsgId,
		ContentType: websocketd.ContentType(in.ContentType),
	})
	if err != nil {
		return nil, resd.NewRpcErrWithTempCtx(l.ctx, "发送消息失败", resd.MqPushErr)
	}
	return &messageRpc.SuccResp{}, nil
}
func (l *SendImLogic) checkReqParams(in *messageRpc.SendImReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	if in.ChannelId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少channelId", resd.ReqFieldRequired1, "channelId")
	}
	if in.ChatType == 0 {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少chatType", resd.ReqFieldRequired1, "chatType")
	}
	return nil
}
