// Code generated by goctl. DO NOT EDIT.
// Source: wechat.proto

package user

import (
	"context"

	"go-zero-dandan/app/wechat/rpc/types/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthByCodeReq  = pb.AuthByCodeReq
	AuthByCodeResp = pb.AuthByCodeResp

	User interface {
		WxpubAuthByCode(ctx context.Context, in *AuthByCodeReq, opts ...grpc.CallOption) (*AuthByCodeResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) WxpubAuthByCode(ctx context.Context, in *AuthByCodeReq, opts ...grpc.CallOption) (*AuthByCodeResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.WxpubAuthByCode(ctx, in, opts...)
}
