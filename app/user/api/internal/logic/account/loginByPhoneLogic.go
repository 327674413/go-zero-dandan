package account

import (
	"context"
	"go-zero-dandan/app/user/api/internal/biz"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/common/constd"

	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type LoginByPhoneLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPhoneLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &LoginByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

func (l *LoginByPhoneLogic) LoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, err
	}
	loginByPhoneStrage := map[int64]func(*types.LoginByPhoneReq) (*types.UserInfoResp, error){
		constd.PlatClasEmMall: l.mallLoginByPhone,
	}
	if strateFunc, ok := loginByPhoneStrage[l.platClasEm]; ok {
		return strateFunc(req)
	} else {
		return l.defaultLoginByPhone(req)
	}
}

// defaultLoginByPhone 默认手机号登录
func (l *LoginByPhoneLogic) defaultLoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	userBiz := biz.NewUserBiz(l.ctx, l.svcCtx)
	phone := *req.Phone
	otpCode := *req.OtpCode
	var phoneArea string
	if req.PhoneArea != nil {
		phoneArea = *req.PhoneArea
	} else {
		phoneArea = constd.PhoneAreaEmChina
	}
	if err = userBiz.CheckPhoneVerifyCode(phone, phoneArea, otpCode); err != nil {
		return nil, err
	}
	userMainModel := model.NewUserMainModel(l.svcCtx.SqlConn, l.platId)
	userMain, err := userMainModel.WhereRaw("phone=?", []any{phone}).Find(l.ctx)
	if err != nil {
		return nil, resd.FailCode(l.lang, resd.MysqlErr)
	}
	resp = &types.UserInfoResp{}
	if userMain.Id == 0 {
		//未注册
		regInfo := biz.UserRegInfo{
			Id:        utild.MakeId(),
			Phone:     phone,
			PhoneArea: phoneArea,
		}
		resp, err = userBiz.RegByPhone(&regInfo)
	} else {
		//已注册
		utild.Copy(&resp, userMain)
	}
	token, err := userBiz.CreateLoginState(userMain)
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
	resp.PlatInfo = map[string]string{
		"test": "aaa",
	}
	return resp, nil
}

func (l *LoginByPhoneLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.FailCode(l.lang, resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.FailCode(l.lang, resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
