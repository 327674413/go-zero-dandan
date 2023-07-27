package {{.pkgName}}

import (
	{{.imports}}
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/common/resd"
    "go-zero-dandan/common/utild"
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   *i18n.Localizer
	platId     int64
    platClasEm int64
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
    localizer := ctx.Value("lang").(*i18n.Localizer)
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	if err = l.initPlat(); err != nil {
    	return l.apiFail(err)
    }

	{{.returnString}}
}

func (l *{{.logic}}) apiFail(err error) {{.responseType}} {
	return nil, resd.ApiFail(l.lang, resd.ErrorCtx(l.ctx, err))
}

func (l *{{.logic}}) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
    if platClasEm == 0 {
        return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
    }
    platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
    if platClasId == 0 {
        return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
    }
    l.platId = platClasId
    l.platClasEm = platClasEm
    return nil
}
