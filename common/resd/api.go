package resd

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type apiErr struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result bool   `json:"result"`
}

func ApiOk(w http.ResponseWriter, r *http.Request, resp any) {
	httpx.OkJsonCtx(r.Context(), w, Succ(resp))
}
func ApiOkAsync(w http.ResponseWriter, r *http.Request, resp any) {
	httpx.OkJsonCtx(r.Context(), w, SuccAsync(resp))
}
func ApiFail(w http.ResponseWriter, r *http.Request, err error) {
	if danErr, ok := AssertErr(err); ok {
		if ok {
			danErr.Msg = I18n.NewLang(r.FormValue("lang")).Msg(danErr.Code, danErr.GetTemps()...)
		}
		httpx.OkJsonCtx(r.Context(), w, apiErr{
			Code:   danErr.Code,
			Msg:    danErr.Msg,
			Result: danErr.Result,
		})
	} else {
		httpx.OkJsonCtx(r.Context(), w, err)
	}
}
