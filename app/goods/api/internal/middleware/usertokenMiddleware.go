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
		userInfo, _ := r.Context().Value("userMainInfo").(*user.UserMainInfo)
		if userInfo != nil && userInfo.Id != "" {
			//解析用户信息成功，且有用户id
			next(w, r)
		} else {
			//解析失败或没有用户id
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			localizer, _ := r.Context().Value("lang").(*i18n.Localizer)
			if userInfo != nil && userInfo.Account == errGetUserInfoFail {
				//约定的数据查询失败类型，报错异常
				if localizer != nil {
					json.NewEncoder(w).Encode(resd.NewErrCtx(r.Context(), land.Msg(localizer, resd.SysErr)))
				} else {
					json.NewEncoder(w).Encode(resd.NewErrCtx(r.Context(), "遇到了点小问题"))
				}
			} else {
				//确实没登陆
				if localizer != nil {
					json.NewEncoder(w).Encode(resd.NewErrCtx(r.Context(), land.Msg(localizer, resd.AuthUserNotLoginErr), resd.AuthUserNotLoginErr))
				} else {
					json.NewEncoder(w).Encode(resd.NewErrCtx(r.Context(), "用户未登陆"))
				}
			}
			return
		}
	}
}
