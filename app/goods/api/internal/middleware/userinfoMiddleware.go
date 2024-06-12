package middleware

import (
	"context"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
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

const (
	errGetUserInfoFail = "查询用户信息失败"
)

func (m *UserInfoMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := &user.UserMainInfo{}
		userToken := r.Header.Get("Token")
		ctx := r.Context()
		var err error
		if userToken != "" {
			//传了token
			userInfo, err = m.UserRpc.GetUserByToken(r.Context(), &user.TokenReq{Token: userToken})
			// 存在报错
			if err != nil && !resd.IsUserNotLoginErr(err) {
				userInfo = &user.UserMainInfo{Account: errGetUserInfoFail}
			}
		}
		ctx = context.WithValue(ctx, "userMainInfo", userInfo)
		//后面断言需要注意，有可能在值为nil
		next(w, r.WithContext(ctx))
	}
}
