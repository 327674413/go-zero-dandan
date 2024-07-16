package account

import (
	"context"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
)

type LoginByWxappCodeLogic struct {
	*LoginByWxappCodeLogicGen
}

func NewLoginByWxappCodeLogic(ctx context.Context, svc *svc.ServiceContext) *LoginByWxappCodeLogic {
	return &LoginByWxappCodeLogic{
		LoginByWxappCodeLogicGen: NewLoginByWxappCodeLogicGen(ctx, svc),
	}
}

func (l *LoginByWxappCodeLogic) LoginByWxappCode(req *types.LoginByWxappCodeReq) (resp *types.LoginByWxappCodeResp, err error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	resp = &types.LoginByWxappCodeResp{
		UserInfo:      &types.UserInfoResp{},
		WxappUserInfo: &types.WxappUserInfoResp{},
	}
	return resp, nil
}
