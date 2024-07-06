// Code generated by goctl. DO NOT EDIT.
package goodsInfo

import (
	"context"

	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/app/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type GetOneLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	lang         string
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
	hasUserInfo  bool
	mustUserInfo bool
	ReqId        string `json:"id,optional"`
	HasReq       struct {
		Id bool
	}
}

func NewGetOneLogicGen(ctx context.Context, svc *svc.ServiceContext) *GetOneLogicGen {
	lang, _ := ctx.Value("lang").(string)
	return &GetOneLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		lang:   lang,
		resd:   resd.NewResd(ctx, resd.I18n.NewLang(lang)),
	}
}

func (l *GetOneLogicGen) initReq(req *types.IdReq) error {
	var err error
	if err = l.initPlat(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}

	if req.Id != nil {
		l.ReqId = *req.Id
		l.HasReq.Id = true
	} else {
		l.HasReq.Id = false
	}
	l.hasUserInfo = true

	return nil
}

func (l *GetOneLogicGen) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *GetOneLogicGen) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platId, _ := l.ctx.Value("platId").(string)
	if platId == "" {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platId
	l.platClasEm = platClasEm
	return nil
}
