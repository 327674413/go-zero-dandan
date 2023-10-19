package logic

import (
	"context"

	"go-zero-dandan/app/wechat/rpc/internal/svc"
	"go-zero-dandan/app/wechat/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxpubAuthByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxpubAuthByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxpubAuthByCodeLogic {
	return &WxpubAuthByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WxpubAuthByCodeLogic) WxpubAuthByCode(in *pb.AuthByCodeReq) (*pb.AuthByCodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AuthByCodeResp{}, nil
}
