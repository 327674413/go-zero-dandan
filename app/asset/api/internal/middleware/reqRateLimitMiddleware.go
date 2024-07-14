package middleware

import (
	"github.com/zeromicro/go-zero/core/limit"
	"go-zero-dandan/common/middled"
	"go-zero-dandan/common/resd"
	"net/http"
)

type ReqRateLimitMiddleware struct {
	limiter *limit.PeriodLimit
}

func NewReqRateLimitMiddleware(limiter *limit.PeriodLimit) *ReqRateLimitMiddleware {
	return &ReqRateLimitMiddleware{
		limiter: limiter,
	}
}

func (m *ReqRateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := middled.ReqRateLimit(r, m.limiter)
		if err == nil {
			next(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			resp := resd.NewResp(r.Context(), resd.I18n.NewLang(r.FormValue("lang")))
			resd.ApiFail(w, r, resp.NewErr(resd.ErrReqRateLimit))
		}

	}
}
