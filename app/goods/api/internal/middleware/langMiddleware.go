package middleware

import (
	"context"
	"net/http"
)

type LangMiddleware struct {
}

func NewLangMiddleware() *LangMiddleware {
	return &LangMiddleware{}
}

func (m *LangMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "lang", lang)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
