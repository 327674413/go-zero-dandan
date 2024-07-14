package middleware

import (
	"go-zero-dandan/common/ctxd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"net/http"
)

type UserTokenMiddleware struct {
}

func NewUserTokenMiddleware() *UserTokenMiddleware {
	return &UserTokenMiddleware{}
}

func (m *UserTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqMeta, _ := r.Context().Value(ctxd.KeyReqMeta).(*typed.ReqMeta)
		if reqMeta != nil && reqMeta.UserId != "" {
			//解析用户信息成功，且有用户id
			next(w, r)
		} else {
			//解析失败或没有用户id
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			resp := resd.NewResp(r.Context(), resd.I18n.NewLang(r.FormValue("lang")))
			if reqMeta != nil && reqMeta.UserErr != "" {
				resd.ApiFail(w, r, resp.NewErr(resd.ErrSys))
			} else {
				resd.ApiFail(w, r, resp.NewErr(resd.ErrAuthUserNotLogin))
			}
			return
		}
	}
}
