package {{.packageName}}

import (
	"context"
	{{.imports}}
)

type {{.logicName}} struct {
	*{{.logicName}}Gen
}

func New{{.logicName}}(ctx context.Context,svc *svc.ServiceContext) *{{.logicName}} {
    return &{{.logicName}} {
        {{.logicName}}Gen:New{{.logicName}}Gen(ctx,svc),
    }
}
{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}req {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}}{{.response}},{{end}} error) {
	if err := l.init({{if .hasReq}}req{{end}}); err != nil {
        return nil, l.resd.Error(err)
    }

	return {{if .hasReply}}&{{.responseType}}{},{{end}} nil
}

func (l *{{.logicName}}) init({{if .hasReq}}req {{.request}}{{end}}) (err error) {
    if err = l.initReq({{if .hasReq}}req{{end}}); err != nil {
        return l.resd.Error(err)
    }
    if err = l.initUser(); err != nil {
        return l.resd.Error(err)
    }
    return nil
}
