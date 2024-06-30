package logic

import (
	"context"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendImAsyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendImAsyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendImAsyncLogic {
	return &SendImAsyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendImAsyncLogic) SendImAsync(in *messageRpc.SendImReq) (*messageRpc.SuccResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}

	return &messageRpc.SuccResp{}, nil
}
func (l *SendImAsyncLogic) checkReqParams(in *messageRpc.SendImReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ReqFieldRequired1, "platId")
	}
	return nil
}
