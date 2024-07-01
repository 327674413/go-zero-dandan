package logic

import (
	"context"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendImChannelMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendImChannelMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendImChannelMsgLogic {
	return &SendImChannelMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendImChannelMsgLogic) SendImChannelMsg(in *messageRpc.SendImChannelMsgReq) (*messageRpc.ResultResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &messageRpc.ResultResp{}, nil
}
func (l *SendImChannelMsgLogic) checkReqParams(in *messageRpc.SendImChannelMsgReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}