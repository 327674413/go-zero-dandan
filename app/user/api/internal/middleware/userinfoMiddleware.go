package middleware

import (
	"context"
	"go-zero-dandan/app/user/rpc/user"
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
		userInfo := &user.UserInfoRpcResp{}
		userToken := r.Header.Get("User-Token")
		var ctx context.Context
		if userToken == "" {
			ctx = context.WithValue(r.Context(), "userInfoRpc", userInfo)
		} else {
			userInfoRpc, err := m.UserRpc.GetUserByToken(r.Context(), &user.TokenReq{Token: userToken})
			if err == nil {
				ctx = context.WithValue(r.Context(), "userInfoRpc", userInfoRpc)
			} else {
				ctx = context.WithValue(r.Context(), "userInfoRpc", userInfo)
			}

		}
		next(w, r.WithContext(ctx))
	}
}
