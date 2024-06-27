// Code generated by goctl. DO NOT EDIT.
// Source: plat.proto

package plat

import (
	"context"

	"go-zero-dandan/app/plat/rpc/types/platRpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdReq    = platRpc.IdReq
	PlatInfo = platRpc.PlatInfo

	Plat interface {
		GetOne(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PlatInfo, error)
	}

	defaultPlat struct {
		cli zrpc.Client
	}
)

func NewPlat(cli zrpc.Client) Plat {
	return &defaultPlat{
		cli: cli,
	}
}

func (m *defaultPlat) GetOne(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PlatInfo, error) {
	client := platRpc.NewPlatClient(m.cli.Conn())
	return client.GetOne(ctx, in, opts...)
}
