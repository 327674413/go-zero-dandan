// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package server

import (
	"context"

	"go-zero-dandan/app/goods/rpc/internal/logic"
	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/pb"
)

type GoodsServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedGoodsServer
}

func NewGoodsServer(svcCtx *svc.ServiceContext) *GoodsServer {
	return &GoodsServer{
		svcCtx: svcCtx,
	}
}

func (s *GoodsServer) GetOne(ctx context.Context, in *pb.IdReq) (*pb.GoodsInfo, error) {
	l := logic.NewGetOneLogic(ctx, s.svcCtx)
	return l.GetOne(in)
}

func (s *GoodsServer) GetPage(ctx context.Context, in *pb.GetPageReq) (*pb.GetPageResp, error) {
	l := logic.NewGetPageLogic(ctx, s.svcCtx)
	return l.GetPage(in)
}

func (s *GoodsServer) GetHotPageByCursor(ctx context.Context, in *pb.GetHotPageByCursorReq) (*pb.GetPageByCursorResp, error) {
	l := logic.NewGetHotPageByCursorLogic(ctx, s.svcCtx)
	return l.GetHotPageByCursor(in)
}
