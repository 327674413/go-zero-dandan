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
	phone := *req.Phone
	platId := l.ctx.Value("platId")
	userId := l.ctx.Value("userId")
	fmt.Println("platId:", platId, ",userId:", userId)
	userMainModel := model.NewUserMainModel()
	userMain, err := userMainModel.Alias("A").Field("id,account").
		WhereRaw("phone=?", []any{phone}).
		Find(l.ctx)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询失败")
	}
	resp = &types.UserInfoResp{}
	if err == model.ErrNotFound {
		fmt.Println("未注册用户")
		//自动注册
	} else {
		resp.Id = userMain.Id
		fmt.Println("已注册用户")
		//直接登录
	}
	return resp, nil
}
