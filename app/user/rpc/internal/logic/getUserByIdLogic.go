package logic

import (
	"context"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"

	"go-zero-dandan/app/user/rpc/internal/svc"
)

type GetUserByIdLogic struct {
	*GetUserByIdLogicGen
}

func NewGetUserByIdLogic(ctx context.Context, svc *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		GetUserByIdLogicGen: NewGetUserByIdLogicGen(ctx, svc),
	}
}

func (l *GetUserByIdLogic) GetUserById(req *userRpc.IdReq) (*userRpc.UserMainInfo, error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	userInfo := &userRpc.UserMainInfo{}
	_, err := l.svc.Redis.GetData(constd.RedisKeyUserInfo, l.req.Id, userInfo)
	if err != nil {
		//有报错
		return nil, l.resd.Error(err, resd.ErrRedisGetUserToken)
	}
	//没报错，且解析后有数据
	if userInfo.Id != "" {
		return userInfo, nil
	}
	//没数据，从数据库查询
	userModel := model.NewUserMainModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	userMain, err := userModel.WhereId(l.req.Id).Find()
	if err != nil {
		return nil, l.resd.Error(err, resd.ErrMysqlSelect)
	}
	if userMain == nil {
		return nil, l.resd.Error(err, resd.ErrNotFoundUser)
	}

	return &userRpc.UserMainInfo{
		Id:        userMain.Id,
		UnionId:   userMain.UnionId,
		StateEm:   userMain.StateEm,
		Account:   userMain.Account,
		Nickname:  userMain.Nickname,
		Phone:     userMain.Phone,
		PhoneArea: userMain.PhoneArea,
		SexEm:     userMain.SexEm,
		Email:     userMain.Email,
		AvatarImg: userMain.AvatarImg,
		PlatId:    userMain.PlatId,
		Signature: userMain.Signature,
	}, nil
}
