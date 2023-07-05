package middleware

import (
	"context"
	"go-zero-dandan/common/land"
	"net/http"
)

type LangMiddleware struct {
}

func NewLangMiddleware() *LangMiddleware {
	return &LangMiddleware{}
}

func (m *LangMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := r.FormValue("lang")
		lang := land.Set(l)
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "lang", lang)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
