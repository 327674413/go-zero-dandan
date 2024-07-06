package middleware

import (
	"encoding/json"
	"go-zero-dandan/app/user/rpc/user"
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
			resp := resd.NewResd(r.Context(), resd.I18n.NewLang(r.FormValue("lang")))
			if userInfo != nil && userInfo.Account == errGetUserInfoFail {
				//约定的数据查询失败类型，报错异常
				json.NewEncoder(w).Encode(resp.NewErr(resd.SysErr))
			} else {
				//确实没登陆
				json.NewEncoder(w).Encode(resp.NewErr(resd.AuthUserNotLoginErr))
			}
			return
		}
	}
}
