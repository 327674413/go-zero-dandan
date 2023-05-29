package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/util"
	"go-zero-dandan/user/api/internal/svc"
	"go-zero-dandan/user/api/internal/types"
	"go-zero-dandan/user/model"
	"gorm.io/gorm"
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
	account := *req.Account
	user := &model.User{}
	password := util.Sha256(*req.Password)
	err = l.svcCtx.DB.Where("account = ? AND password = ?", account, password).First(user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logx.Error("[DB ERROR]:", err)
			err = errors.New("遇到了点小问题")
			return
		} else {
			err = errors.New("用户名或密码不正确")
			return
		}

	}
	resp = &types.AccountLoginResp{
		Token: "",
	}
	return
}
