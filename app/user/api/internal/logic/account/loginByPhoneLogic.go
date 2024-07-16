package account

import (
	"context"
	"go-zero-dandan/app/user/api/internal/biz"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/utild/copier"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"go-zero-dandan/common/utild"
)

type LoginByPhoneLogic struct {
	*LoginByPhoneLogicGen
}

func NewLoginByPhoneLogic(ctx context.Context, svc *svc.ServiceContext) *LoginByPhoneLogic {
	return &LoginByPhoneLogic{
		LoginByPhoneLogicGen: NewLoginByPhoneLogicGen(ctx, svc),
	}
}

func (l *LoginByPhoneLogic) LoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	if err := l.initReq(req); err != nil {
		return nil, l.resd.Error(err)
	}
	loginByPhoneStrage := map[int64]func(*types.LoginByPhoneReq) (*types.UserInfoResp, error){
		constd.PlatClasEmMall: l.mallLoginByPhone,
	}
	if strateFunc, ok := loginByPhoneStrage[l.meta.PlatClasEm]; ok {
		return strateFunc(req)
	} else {
		return l.defaultLoginByPhone(req)
	}
}

// defaultLoginByPhone 默认手机号登录
func (l *LoginByPhoneLogic) defaultLoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	userBiz := biz.NewUserBiz(l.ctx, l.svc)
	phone := *req.Phone
	otpCode := *req.OtpCode
	var phoneArea string
	if req.PhoneArea != nil {
		phoneArea = *req.PhoneArea
	} else {
		phoneArea = constd.PhoneAreaEmChina
	}
	if l.svc.Config.Mode != "dev" || otpCode != "5210" {
		if err = userBiz.CheckPhoneVerifyCode(phone, phoneArea, otpCode); err != nil {
			return nil, l.resd.Error(err)
		}
	}

	userMainModel := model.NewUserMainModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	userMain, err := userMainModel.Where("phone=?", phone).Find()
	if err != nil {
		return nil, l.resd.Error(err)
	}
	resp = &types.UserInfoResp{}
	if userMain == nil {
		//未注册
		regInfo := biz.UserRegInfo{
			Id:        utild.MakeId(),
			Phone:     phone,
			PhoneArea: phoneArea,
		}
		resp, err = userBiz.RegByPhone(&regInfo)
	} else {
		//已注册
		copier.Copy(&resp, userMain)
	}
	token, err := userBiz.CreateLoginState(resp)
	if err != nil {
		return nil, err
	} else {
		resp.UserToken = token
		return resp, nil
	}

}

// mallLoginByPhone 商城应用登录
func (l *LoginByPhoneLogic) mallLoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	resp, err = l.defaultLoginByPhone(req)
	if err != nil {
		return resp, err
	}
	resp.PlatId = l.meta.PlatId
	return resp, nil
}
