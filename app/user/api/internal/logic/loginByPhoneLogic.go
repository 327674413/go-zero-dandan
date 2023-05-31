package logic

import (
	"context"
	"errors"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/common/util"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPhoneLogic {
	return &LoginByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPhoneLogic) LoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	userMain, err := l.svcCtx.UserMainModel.FindOne(l.ctx, 1)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询失败")
	}

	if err == model.ErrNotFound {
		//自动注册
	} else {
		//直接登录
		resp = &types.UserInfoResp{}
		resp.UserToken = "token"
		util.ObjToObj(userMain, resp)
	}
	return resp, nil
}
