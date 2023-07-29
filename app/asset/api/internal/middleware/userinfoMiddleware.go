package middleware

import (
	"context"
	"fmt"
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
		userInfo := &user.UserMainInfo{}
		userToken := r.Header.Get("Token")
		var ctx context.Context
		fmt.Println(userToken)
		if userToken == "" {
			ctx = context.WithValue(r.Context(), "userMainInfo", userInfo)
		} else {
			userInfoRpc, err := m.UserRpc.GetUserByToken(r.Context(), &user.TokenReq{Token: userToken})
			if err == nil {
				ctx = context.WithValue(r.Context(), "userMainInfo", userInfoRpc)
			} else {
				ctx = context.WithValue(r.Context(), "userMainInfo", userInfo)
			}

		}
		next(w, r.WithContext(ctx))
	}
}
