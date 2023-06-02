package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/app/user/model"
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

	userMain, err := l.svcCtx.UserMainModel.WhereId(1).Find(l.ctx)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询失败")
	}
	if err == model.ErrNotFound {
		fmt.Println("1111111")
		//自动注册
	} else {
		resp.Id = userMain.Id
		fmt.Println("222222")
		//直接登录
	}
	return resp, nil
}
