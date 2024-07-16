package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-dandan/app/plat/api/internal/svc"
	"go-zero-dandan/app/plat/api/internal/types"
	"go-zero-dandan/app/plat/model"
	"go-zero-dandan/common/ctxd"
	"go-zero-dandan/common/resd"
	"time"
)

type GetTokenLogic struct {
	*GetTokenLogicGen
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		GetTokenLogicGen: NewGetTokenLogicGen(ctx, svcCtx),
	}
}
func (l *GetTokenLogic) GetToken(req *types.GetTokenReq) (resp *types.GetTokenResp, err error) {
	platModel := model.NewPlatMainModel(l.ctx, l.svc.SqlConn)
	platMain, err := platModel.Where("appid = ? and secret = ?", req.Appid, req.Secret).Find()
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	resp = &types.GetTokenResp{}
	if platMain == nil {
		return nil, l.resd.NewErr(resd.ErrPlatInvalid)
	} else {
		resp.Token, err = l.getToken(l.svc.Config.Auth.AccessSecret, time.Now().Unix(), l.svc.Config.Auth.AccessExpire, platMain)
		resp.ExpireSec = l.svc.Config.Auth.AccessExpire
		//直接登录
	}
	return resp, nil
}
func (l *GetTokenLogic) getToken(secretKey string, iat int64, seconds int64, platMian *model.PlatMain) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxd.KeyPlatId] = platMian.Id
	claims[ctxd.KeyPlatClasEm] = platMian.ClasEm
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *GetTokenLogic) init(req *types.GetTokenReq) (err error) {
	if err = l.initReq(req); err != nil {
		return l.resd.Error(err)
	}
	return nil
}
