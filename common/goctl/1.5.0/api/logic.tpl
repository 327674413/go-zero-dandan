package {{.pkgName}}

import (
	{{.imports}}
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   *i18n.Localizer
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
	// todo: add your logic here and delete this line

	{{.returnString}}
}
