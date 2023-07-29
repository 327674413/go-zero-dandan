package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/pb"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
)

type GetUserByTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByTokenLogic {
	return &GetUserByTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByTokenLogic) GetUserByToken(in *pb.TokenReq) (*pb.UserMainInfo, error) {
	userInfo := &pb.UserMainInfo{}
	err := l.svcCtx.Redis.GetData(constd.RedisKeyUserToken, in.Token, userInfo)
	if err != nil {
		logx.Error(err)
		return nil, resd.RpcEncodeTempErr(resd.RedisGetUserTokenErr)
	}
	return userInfo, nil
}
