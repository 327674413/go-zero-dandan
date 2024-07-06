// Code generated by goctl. DO NOT EDIT.
package {{.pkgName}}
import (
	{{.imports}}
	"go-zero-dandan/common/resd"
    "go-zero-dandan/common/utild"
    "go-zero-dandan/app/user/rpc/user"
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svc *svc.ServiceContext
	resd   *resd.Resp
	lang string
	userMainInfo *user.UserMainInfo
	platId     string
    platClasEm int64
    hasUserInfo bool
    mustUserInfo bool
    {{.danLogicVars}}
}

func New{{.logic}}(ctx context.Context, svc *svc.ServiceContext) *{{.logic}} {
    lang, _ := ctx.Value("lang").(string)
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc: svc,
		lang:lang,
		resd:   resd.NewResd(ctx, resd.I18n.NewLang(lang)),
	}
}

func (l *{{.logic}}) initReq({{.request}}) error {
    var err error
	if err = l.initPlat(); err != nil {
    	return resd.ErrorCtx(l.ctx,err)
    }
    {{.danInitReqFields}}
	return nil
}

func (l *{{.logic}}) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx,"未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *{{.logic}}) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
    if platClasEm == 0 {
        return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
    }
    platId,_ := l.ctx.Value("platId").(string)
    if platId == "" {
        return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
    }
    l.platId = platId
    l.platClasEm = platClasEm
    return nil
}
