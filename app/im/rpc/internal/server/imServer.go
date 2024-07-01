// Code generated by goctl. DO NOT EDIT.
// Source: im.proto

package server

import (
	"context"

	"go-zero-dandan/app/im/rpc/internal/logic"
	"go-zero-dandan/app/im/rpc/internal/svc"
	"go-zero-dandan/app/im/rpc/types/imRpc"
)

type ImServer struct {
	svcCtx *svc.ServiceContext
	imRpc.UnimplementedImServer
}

func NewImServer(svcCtx *svc.ServiceContext) *ImServer {
	return &ImServer{
		svcCtx: svcCtx,
	}
}

// 获取会话记录
func (s *ImServer) GetChatLog(ctx context.Context, in *imRpc.GetChatLogReq) (*imRpc.GetChatLogResp, error) {
	l := logic.NewGetChatLogLogic(ctx, s.svcCtx)
	return l.GetChatLog(in)
}

// 建立会话: 群聊, 私聊
func (s *ImServer) SetUpUserConversation(ctx context.Context, in *imRpc.SetUpUserConversationReq) (*imRpc.SetUpUserConversationResp, error) {
	l := logic.NewSetUpUserConversationLogic(ctx, s.svcCtx)
	return l.SetUpUserConversation(in)
}

// 获取会话
func (s *ImServer) GetConversations(ctx context.Context, in *imRpc.GetConversationsReq) (*imRpc.GetConversationsResp, error) {
	l := logic.NewGetConversationsLogic(ctx, s.svcCtx)
	return l.GetConversations(in)
}

// 更新会话
func (s *ImServer) PutConversations(ctx context.Context, in *imRpc.PutConversationsReq) (*imRpc.PutConversationsResp, error) {
	l := logic.NewPutConversationsLogic(ctx, s.svcCtx)
	return l.PutConversations(in)
}

func (s *ImServer) CreateGroupConversation(ctx context.Context, in *imRpc.CreateGroupConversationReq) (*imRpc.CreateGroupConversationResp, error) {
	l := logic.NewCreateGroupConversationLogic(ctx, s.svcCtx)
	return l.CreateGroupConversation(in)
}

// 发送系统消息
func (s *ImServer) SendSysMsg(ctx context.Context, in *imRpc.SendSysMsgReq) (*imRpc.ResultResp, error) {
	l := logic.NewSendSysMsgLogic(ctx, s.svcCtx)
	return l.SendSysMsg(in)
}
