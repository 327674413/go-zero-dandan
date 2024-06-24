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

//func (m *UserInfoMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		userInfo := &user.UserMainInfo{}
//		userToken := r.Header.Get("Token")
//		var ctx context.Context
//		if userToken == "" {
//			//未传token，没登陆
//			ctx = context.WithValue(r.Context(), "userMainInfo", userInfo)
//		} else {
//			//传了token
//			userInfoFind, err := m.UserRpc.GetUserByToken(r.Context(), &user.TokenReq{Token: userToken})
//			if err == nil {
//				//查询正常,不管userInfoFind是有数据还是nil，都往下丢，后面用的时候断言来判断
//				ctx = context.WithValue(r.Context(), "userMainInfo", userInfoFind)
//			} else {
//				//查询报错，并且不是未登陆的错误类型，则视为查询报错，约定往Account属性里加标识
//				if !resd.IsUserNotLoginErr(err) {
//					userInfo.Account = errGetUserInfoFail
//				}
//				ctx = context.WithValue(r.Context(), "userMainInfo", userInfo)
//				//未登陆就不用处理，直接往下
//			}
//
//		}
//		next(w, r.WithContext(ctx))
//	}
//}
