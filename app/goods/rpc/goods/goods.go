// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package goods

import (
	"context"

	"go-zero-dandan/app/goods/rpc/types/goodsRpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmptyReq              = goodsRpc.EmptyReq
	GetHotPageByCursorReq = goodsRpc.GetHotPageByCursorReq
	GetPageByCursorResp   = goodsRpc.GetPageByCursorResp
	GetPageReq            = goodsRpc.GetPageReq
	GetPageResp           = goodsRpc.GetPageResp
	GoodsInfo             = goodsRpc.GoodsInfo
	IdReq                 = goodsRpc.IdReq
	SuccResp              = goodsRpc.SuccResp

	Goods interface {
		GetOne(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*GoodsInfo, error)
		GetPage(ctx context.Context, in *GetPageReq, opts ...grpc.CallOption) (*GetPageResp, error)
		GetHotPageByCursor(ctx context.Context, in *GetHotPageByCursorReq, opts ...grpc.CallOption) (*GetPageByCursorResp, error)
	}

	defaultGoods struct {
		cli zrpc.Client
	}
)

func NewGoods(cli zrpc.Client) Goods {
	return &defaultGoods{
		cli: cli,
	}
}

func (m *defaultGoods) GetOne(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*GoodsInfo, error) {
	client := goodsRpc.NewGoodsClient(m.cli.Conn())
	return client.GetOne(ctx, in, opts...)
}

func (m *defaultGoods) GetPage(ctx context.Context, in *GetPageReq, opts ...grpc.CallOption) (*GetPageResp, error) {
	client := goodsRpc.NewGoodsClient(m.cli.Conn())
	return client.GetPage(ctx, in, opts...)
}

func (m *defaultGoods) GetHotPageByCursor(ctx context.Context, in *GetHotPageByCursorReq, opts ...grpc.CallOption) (*GetPageByCursorResp, error) {
	client := goodsRpc.NewGoodsClient(m.cli.Conn())
	return client.GetHotPageByCursor(ctx, in, opts...)
}
