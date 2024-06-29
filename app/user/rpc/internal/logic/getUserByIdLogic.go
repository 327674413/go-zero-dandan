package logic

import (
	"context"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/internal/svc"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *userRpc.IdReq) (*userRpc.UserMainInfo, error) {
	userInfo := &userRpc.UserMainInfo{}
	_, err := l.svcCtx.Redis.GetData(constd.RedisKeyUserInfo, in.Id, userInfo)
	if err != nil {
		//有报错
		return nil, resd.RpcErrEncode(resd.ErrorCtx(l.ctx, err, resd.RedisGetUserTokenErr))
	}
	//没报错，且解析后有数据
	if userInfo.Id != "" {
		return userInfo, nil
	}
	//没数据，从数据库查询
	userModel := model.NewUserMainModel(l.svcCtx.SqlConn, in.PlatId)
	userMain, err := userModel.WhereId(in.Id).Find()
	if err != nil {
		return nil, resd.RpcErrEncode(resd.ErrorCtx(l.ctx, err, resd.MysqlSelectErr))
	}
	if userMain == nil {
		return nil, resd.RpcErrEncode(resd.NewErrCtx(l.ctx, "不存在该用户", resd.NotFoundUser))
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
