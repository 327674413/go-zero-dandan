package {{.pkgName}}
import (
    "go-zero-dandan/common/resd"
	{{.imports}}
)

type {{.logic}} struct {
	*{{.logic}}Gen
}

func New{{.logic}}(ctx context.Context, svc *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		{{.logic}}Gen: New{{.logic}}Gen(ctx,svc),
	}
}
func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
    if err = l.init(req);err != nil{
         return l.resd.Error(err)
    }

	{{.returnString}}
}