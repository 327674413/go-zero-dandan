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
		userInfo := r.Context().Value("userInfoRpc").(*user.UserInfoRpcResp)
		if userInfo.Id == 0 {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":   resd.Auth,
				"result": false,
				"msg":    "授权失效",
			})
			return
		}
		next(w, r)
	}
}
