package logic

import (
	"context"

	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgLogic) SendMsg(in *pb.SendMsgReq) (*pb.Success, error) {
	// todo: add your logic here and delete this line

	return &pb.Success{}, nil
}
