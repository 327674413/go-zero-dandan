package account

import (
	"context"
	"go-zero-dandan/app/user/api/internal/biz"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/common/utild/copier"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
)

type LoginByAccountLogic struct {
	*LoginByAccountLogicGen
}

func NewLoginByAccountLogic(ctx context.Context, svc *svc.ServiceContext) *LoginByAccountLogic {
	return &LoginByAccountLogic{
		LoginByAccountLogicGen: NewLoginByAccountLogicGen(ctx, svc),
	}
}
func (l *LoginByAccountLogic) LoginByAccount(in *types.LoginByAccountReq) (resp *types.UserInfoResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	password, err := l.loginPass(l.req.Password)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	userMainModel := model.NewUserMainModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	userMain, err := userMainModel.Where("account=? and password=?", l.req.Account, password).Find()
	if err != nil {
		return nil, l.resd.Error(err)
	}
	resp = &types.UserInfoResp{}
	if userMain == nil {
		//未注册
		return nil, l.resd.NewErr(resd.ErrLoginAcctOrPassInvalid)
	} else {
		//已注册
		copier.Copy(&resp, userMain)
	}
	userBiz := biz.NewUserBiz(l.ctx, l.svc, l.resd, l.meta)
	token, err := userBiz.CreateLoginState(resp)
	if err != nil {
		return nil, l.resd.Error(err)
	} else {
		resp.UserToken = token
		return resp, nil
	}
}

func (l *LoginByAccountLogic) loginPass(password string) (pass string, err error) {
	if password == "" {
		return "", l.resd.NewErrWithTemp(resd.ErrReqFieldEmpty1, resd.VarPassword)
	}
	return utild.Sha256(password), nil
}
