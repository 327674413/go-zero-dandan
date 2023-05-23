package logic

import (
	"context"
	"errors"
	util "go-zero-dandan/common/util"
	"go-zero-dandan/user/model"

	"go-zero-dandan/user/api/internal/svc"
	"go-zero-dandan/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLoginLogic {
	return &AccountLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLoginLogic) AccountLogin(req *types.AccountLoginReq) (resp *types.AccountLoginResp, err error) {
	user := &model.UserDetail{}
	password := util.Sha256(req.Password)
	err = l.svcCtx.DB.Where("account = ? AND password = ?", req.Account, password).First(user).Error
	if err != nil {
		logx.Error("[DB ERROR]:", err)
		err = errors.New("用户名或密码不正确")
		return
	}
	resp = &types.AccountLoginResp{
		Token: "",
	}
	return
}
