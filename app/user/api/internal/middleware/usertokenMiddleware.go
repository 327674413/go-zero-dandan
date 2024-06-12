package middleware

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/land"
	"go-zero-dandan/common/resd"
	"net/http"
)

type UserTokenMiddleware struct {
}

func NewUserTokenMiddleware() *UserTokenMiddleware {
	return &UserTokenMiddleware{}
}

func (m *UserTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo, ok := r.Context().Value("userMainInfo").(*user.UserMainInfo)
		if !ok || userInfo.Id == "" {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			localizer := r.Context().Value("lang").(*i18n.Localizer)
			json.NewEncoder(w).Encode(resd.NewErrCtx(r.Context(), land.Msg(localizer, resd.AuthUserNotLoginErr), resd.AuthUserNotLoginErr))
			return
		}
		next(w, r)
	}
}
