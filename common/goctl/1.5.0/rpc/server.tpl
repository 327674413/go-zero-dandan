{{.head}}

package server

import (
	{{if .notStream}}"context"{{end}}
    "go-zero-dandan/common/resd"
    "encoding/json"
    "errors"
	{{.imports}}
)

type {{.server}}Server struct {
	svcCtx *svc.ServiceContext
	{{.unimplementedServer}}
}

func New{{.server}}Server(svcCtx *svc.ServiceContext) *{{.server}}Server {
	return &{{.server}}Server{
		svcCtx: svcCtx,
	}
}

{{.funcs}}
