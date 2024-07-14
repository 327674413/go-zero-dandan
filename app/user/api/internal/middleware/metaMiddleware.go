package middleware

import (
	"go-zero-dandan/common/middled"
	"net/http"
)

type MetaMiddleware struct {
}

func NewMetaMiddleware() *MetaMiddleware {
	return &MetaMiddleware{}
}

func (m *MetaMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := middled.SetCtxMeta(r)
		next(w, r.WithContext(ctx))
	}
}
