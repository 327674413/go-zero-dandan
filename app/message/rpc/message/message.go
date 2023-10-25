// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package message

import (
	"context"

	"go-zero-dandan/app/message/rpc/types/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SendPhoneReq = pb.SendPhoneReq
	SuccResp     = pb.SuccResp

	Message interface {
		SendPhone(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*SuccResp, error)
		SendSMSAsync(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*SuccResp, error)
	}

	defaultMessage struct {
		cli zrpc.Client
	}
)

func NewMessage(cli zrpc.Client) Message {
	return &defaultMessage{
		cli: cli,
	}
}

func (m *defaultMessage) SendPhone(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*SuccResp, error) {
	client := pb.NewMessageClient(m.cli.Conn())
	return client.SendPhone(ctx, in, opts...)
}

func (m *defaultMessage) SendSMSAsync(ctx context.Context, in *SendPhoneReq, opts ...grpc.CallOption) (*SuccResp, error) {
	client := pb.NewMessageClient(m.cli.Conn())
	return client.SendSMSAsync(ctx, in, opts...)
}
