package logic

import (
	"context"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
)

type SendSysMsgLogic struct {
	*SendSysMsgLogicGen
}

func NewSendSysMsgLogic(ctx context.Context, svc *svc.ServiceContext) *SendSysMsgLogic {
	return &SendSysMsgLogic{
		SendSysMsgLogicGen: NewSendSysMsgLogicGen(ctx, svc),
	}
}

// SendSysMsg 发送系统消息
func (l *SendSysMsgLogic) SendSysMsg(in *imRpc.SendSysMsgReq) (*imRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	err := l.svc.SysToUserTransferClient.Push(&kafkad.SysToUserMsg{
		MsgClas:    websocketd.MsgClas(l.req.MsgClasEm),
		UserId:     l.req.UserId,
		SendTime:   l.req.SendTime,
		MsgType:    websocketd.MsgType(l.req.MsgTypeEm),
		MsgContent: l.req.MsgContent,
	})
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrMqPush)
	}
	return &imRpc.ResultResp{Code: constd.ResultTasking}, nil
}
