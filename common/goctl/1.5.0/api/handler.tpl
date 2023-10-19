package {{.PkgName}}

import (
	"net/http"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/land"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, resd.Error(err))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        if err != nil {
            if danErr, ok := resd.AssertErr(err); ok {
                localizer, ok := r.Context().Value("lang").(*i18n.Localizer)
                if ok {
                    danErr.Msg = land.Msg(localizer, danErr.Code, danErr.GetTemps())
                }
                httpx.OkJsonCtx(r.Context(), w, danErr)
            } else {
                httpx.OkJsonCtx(r.Context(), w, err)
            }
		} else {
			{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resd.Succ(resp)){{else}}httpx.Ok(w){{end}}
		}
	}
}
