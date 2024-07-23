package {{.pkgName}}
import (
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
    if err = l.initReq(in);err != nil{
         return nil,l.resd.Error(err)
    }
	{{.returnString}}
}