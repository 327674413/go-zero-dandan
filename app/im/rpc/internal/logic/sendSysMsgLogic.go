package logic

import (
	"context"
	"go-zero-dandan/app/im/mq/kafkad"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
	"go-zero-dandan/app/im/ws/websocketd"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSysMsgLogic struct {
	ctx context.Context
	svc *svc.ServiceContext
	logx.Logger
}

func NewSendSysMsgLogic(ctx context.Context, svc *svc.ServiceContext) *SendSysMsgLogic {
	return &SendSysMsgLogic{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
	}
}

// SendSysMsg 发送系统消息
func (l *SendSysMsgLogic) SendSysMsg(in *imRpc.SendSysMsgReq) (*imRpc.ResultResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	err := l.svc.SysToUserTransferClient.Push(&kafkad.SysToUserMsg{
		MsgClas:    websocketd.MsgClas(in.MsgClasEm),
		UserId:     in.UserId,
		SendTime:   in.SendTime,
		MsgType:    websocketd.MsgType(in.MsgTypeEm),
		MsgContent: in.MsgContent,
	})
	if err != nil {
		return nil, l.resd.NewRpcErrWithTempCtx(l.ctx, "发送消息失败", resd.MqPushErr)
	}
	return &imRpc.ResultResp{Code: constd.ResultTasking}, nil
}
func (l *SendSysMsgLogic) checkReqParams(in *imRpc.SendSysMsgReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ErrReqFieldRequired1, "platId")
	}
	return nil
}
