package middleware

import (
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/middled"
	"net/http"
)

type UserInfoMiddleware struct {
	UserRpc user.User
}

func NewUserInfoMiddleware(userRpc user.User) *UserInfoMiddleware {
	return &UserInfoMiddleware{
		UserRpc: userRpc,
	}
}

func (m *UserInfoMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := middled.SetCtxUser(r, m.UserRpc)
		next(w, r.WithContext(ctx))
	}
}
