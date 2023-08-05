package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/app/plat/api/internal/svc"
	"go-zero-dandan/app/plat/api/internal/types"
	"go-zero-dandan/app/plat/model"
	"go-zero-dandan/common/resd"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   *i18n.Localizer
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenLogic) GetToken(req *types.GetTokenReq) (resp *types.GetTokenResp, err error) {
	platModel := model.NewPlatMainModel(l.svcCtx.SqlConn)
	platMain, err := platModel.Ctx(l.ctx).Where("appid = ? and secret = ?", req.Appid, req.Secret).Find()
	if err != nil && err != model.ErrNotFound {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	resp = &types.GetTokenResp{}
	if err == model.ErrNotFound {
		return nil, resd.NewErrCtx(l.ctx, "无效plat", resd.PlatInvalid)
	} else {
		resp.Token, err = l.getToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, platMain)
		resp.ExpireSec = l.svcCtx.Config.Auth.AccessExpire
		//直接登录
	}
	return resp, nil
}

func (l *GetTokenLogic) getToken(secretKey string, iat int64, seconds int64, platMian *model.PlatMain) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["platId"] = platMian.Id
	claims["platClasEm"] = platMian.ClasEm
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
