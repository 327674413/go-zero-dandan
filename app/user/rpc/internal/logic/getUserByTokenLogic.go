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
	//return nil, resd.NewRpcErrCtx(l.ctx, "1111", resd.AuthUserNotLoginErr)
	err := l.svcCtx.Redis.GetData(constd.RedisKeyUserToken, in.Token, userInfo)
	if err != nil {
		//有报错
		return nil, resd.RpcErrEncode(resd.ErrorCtx(l.ctx, err, resd.RedisGetUserTokenErr))
	}
	//没报错，且解析后有数据
	if userInfo.Id != "" {
		return userInfo, nil
	}
	//没报错，没数据，当做没登陆（redis默认有rdb持久化，快照方式页够用了，所以登陆态数据库不存了）
	return nil, resd.NewRpcErrCtx(l.ctx, "未登陆", resd.AuthUserNotLoginErr)

}
