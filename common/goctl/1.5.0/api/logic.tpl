package {{.pkgName}}
import (
    "go-zero-dandan/common/resd"
	{{.imports}}
)

type {{.logic}} struct {
	*{{.logic}}Gen
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		{{.logic}}Gen: New{{.logic}}Gen(ctx,svcCtx),
	}
}
func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
    if err = l.init(req);err != nil{
         return nil,resd.ErrorCtx(l.ctx,err)
    }

	{{.returnString}}
}
func (l *{{.logic}}) init({{.request}}) (err error) {
    if err = l.initReq(req);err != nil{
         return resd.ErrorCtx(l.ctx,err)
    }
    if err = l.initUser(); err != nil {
        return resd.ErrorCtx(l.ctx,err)
    }
    return nil
}