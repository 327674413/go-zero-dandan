// Code generated by goctl. DO NOT EDIT.
// Source: im.proto

package im

import (
	"context"

	"go-zero-dandan/app/im/rpc/types/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChatLog                     = pb.ChatLog
	Conversation                = pb.Conversation
	CreateGroupConversationReq  = pb.CreateGroupConversationReq
	CreateGroupConversationResp = pb.CreateGroupConversationResp
	GetChatLogReq               = pb.GetChatLogReq
	GetChatLogResp              = pb.GetChatLogResp
	GetConversationsReq         = pb.GetConversationsReq
	GetConversationsResp        = pb.GetConversationsResp
	PutConversationsReq         = pb.PutConversationsReq
	PutConversationsResp        = pb.PutConversationsResp
	SetUpUserConversationReq    = pb.SetUpUserConversationReq
	SetUpUserConversationResp   = pb.SetUpUserConversationResp

	Im interface {
		// 获取会话记录
		GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error)
		// 建立会话: 群聊, 私聊
		SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error)
		// 获取会话
		GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error)
		// 更新会话
		PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error)
		CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error)
	}

	defaultIm struct {
		cli zrpc.Client
	}
)

func NewIm(cli zrpc.Client) Im {
	return &defaultIm{
		cli: cli,
	}
}

// 获取会话记录
func (m *defaultIm) GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error) {
	client := pb.NewImClient(m.cli.Conn())
	return client.GetChatLog(ctx, in, opts...)
}

// 建立会话: 群聊, 私聊
func (m *defaultIm) SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error) {
	client := pb.NewImClient(m.cli.Conn())
	return client.SetUpUserConversation(ctx, in, opts...)
}

// 获取会话
func (m *defaultIm) GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error) {
	client := pb.NewImClient(m.cli.Conn())
	return client.GetConversations(ctx, in, opts...)
}

// 更新会话
func (m *defaultIm) PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error) {
	client := pb.NewImClient(m.cli.Conn())
	return client.PutConversations(ctx, in, opts...)
}

func (m *defaultIm) CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error) {
	client := pb.NewImClient(m.cli.Conn())
	return client.CreateGroupConversation(ctx, in, opts...)
}