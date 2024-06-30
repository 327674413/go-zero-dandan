// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package server

import (
	"context"

	"go-zero-dandan/app/message/rpc/internal/logic"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
)

type MessageServer struct {
	svcCtx *svc.ServiceContext
	messageRpc.UnimplementedMessageServer
}

func NewMessageServer(svcCtx *svc.ServiceContext) *MessageServer {
	return &MessageServer{
		svcCtx: svcCtx,
	}
}

func (s *MessageServer) SendPhone(ctx context.Context, in *messageRpc.SendPhoneReq) (*messageRpc.SuccResp, error) {
	l := logic.NewSendPhoneLogic(ctx, s.svcCtx)
	return l.SendPhone(in)
}

func (s *MessageServer) SendPhoneAsync(ctx context.Context, in *messageRpc.SendPhoneReq) (*messageRpc.SuccResp, error) {
	l := logic.NewSendPhoneAsyncLogic(ctx, s.svcCtx)
	return l.SendPhoneAsync(in)
}

func (s *MessageServer) SendIm(ctx context.Context, in *messageRpc.SendImReq) (*messageRpc.SuccResp, error) {
	l := logic.NewSendImLogic(ctx, s.svcCtx)
	return l.SendIm(in)
}

func (s *MessageServer) SendImAsync(ctx context.Context, in *messageRpc.SendImReq) (*messageRpc.SuccResp, error) {
	l := logic.NewSendImAsyncLogic(ctx, s.svcCtx)
	return l.SendImAsync(in)
}
