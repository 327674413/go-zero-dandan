package logic

import (
	"context"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
)

type GetUserByTokenLogic struct {
	*GetUserByTokenLogicGen
}

func NewGetUserByTokenLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserByTokenLogic {
	return &GetUserByTokenLogic{
		GetUserByTokenLogicGen: NewGetUserByTokenLogicGen(ctx, svc),
	}
}

func (l *GetUserByTokenLogic) GetUserByToken(req *userRpc.TokenReq) (*userRpc.UserMainInfo, error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	userInfo := &userRpc.UserMainInfo{}
	_, err := l.svc.Redis.GetDataCtx(l.ctx, constd.RedisKeyUserToken, l.req.Token, userInfo)
	if err != nil {
		//有报错
		return nil, l.resd.Error(err, resd.ErrRedisGetUserToken)
	}
	//没报错，且解析后有数据
	if userInfo.Id != "" {
		return userInfo, nil
	}
	//没报错，没数据，当做没登陆（redis默认有rdb持久化，快照方式页够用了，所以登陆态数据库不存了）
	err = l.resd.NewErr(resd.ErrAuthUserNotLogin)
	return nil, err

}
