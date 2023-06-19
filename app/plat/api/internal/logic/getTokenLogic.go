package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/app/plat/api/internal/svc"
	"go-zero-dandan/app/plat/api/internal/types"
	"go-zero-dandan/app/plat/model"
	"go-zero-dandan/common/respd"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenLogic) GetToken(req *types.GetTokenReq) (resp *types.GetTokenResp, err error) {
	localizer := l.ctx.Value("lang").(*i18n.Localizer)
	platModel := model.NewPlatMainModel()
	platMain, err := platModel.WhereRaw("appid = ? and secret = ?", []any{req.Appid, req.Secret}).Find(l.ctx)
	if err != nil && err != model.ErrNotFound {
		return nil, respd.FailCode(localizer, respd.MysqlErr)
	}
	resp = &types.GetTokenResp{}
	if err == model.ErrNotFound {
		return nil, respd.FailCode(localizer, respd.PlatInvalid)
	} else {
		resp.Token, err = l.getToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, platMain.Id)
		resp.ExpireSec = l.svcCtx.Config.Auth.AccessExpire
		//直接登录
	}
	return resp, nil
}

func (l *GetTokenLogic) getToken(secretKey string, iat, seconds, platId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["platId"] = platId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
