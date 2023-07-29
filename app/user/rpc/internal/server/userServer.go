// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero-dandan/app/user/rpc/internal/logic"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUserByToken(ctx context.Context, in *pb.TokenReq) (*pb.UserMainInfo, error) {
	l := logic.NewGetUserByTokenLogic(ctx, s.svcCtx)
	return l.GetUserByToken(in)
}

func (s *UserServer) EditUserInfo(ctx context.Context, in *pb.EditUserInfoReq) (*pb.SuccResp, error) {
	l := logic.NewEditUserInfoLogic(ctx, s.svcCtx)
	return l.EditUserInfo(in)
}
