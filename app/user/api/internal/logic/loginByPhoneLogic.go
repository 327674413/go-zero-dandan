package logic

import (
	"context"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/api/internal/svc"
	"go-zero-dandan/app/user/api/internal/types"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type LoginByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   *i18n.Localizer
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
	platClasEm := utild.AnyToInt(l.ctx.Value("clasEm"))
	loginByPhoneStrage := map[int]func(*types.LoginByPhoneReq) (*types.UserInfoResp, error){
		constd.PlatClasEmMall: l.mallLoginByPhone,
	}
	if strateFunc, ok := loginByPhoneStrage[platClasEm]; ok {
		return strateFunc(req)
	} else {
		return l.defaultLoginByPhone(req)
	}

}
func (l *LoginByPhoneLogic) defaultLoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	phone := *req.Phone
	platId := utild.AnyToInt64(l.ctx.Value("platId"))
	userMainModel := model.NewUserMainModel(l.svcCtx.SqlConn, platId)
	userMain, err := userMainModel.WhereRaw("phone=?", []any{phone}).Find(l.ctx)
	if err != nil {
		return nil, resd.FailCode(l.lang, resd.MysqlErr)
	}
	if userMain.Id == 0 {
		//未注册
		return nil, nil
	} else {
		//已注册
		resp = &types.UserInfoResp{}
		utild.Copy(&resp, userMain)
		return resp, nil
	}
	return nil, nil
}
func (l *LoginByPhoneLogic) mallLoginByPhone(req *types.LoginByPhoneReq) (resp *types.UserInfoResp, err error) {
	//商城应用
	resp, err = l.defaultLoginByPhone(req)
	if err != nil {
		return resp, err
	}
	resp.PlatInfo = map[string]string{
		"test": "aaa",
	}
	return resp, nil
}
