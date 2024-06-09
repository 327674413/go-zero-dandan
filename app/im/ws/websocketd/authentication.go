package websocketd

import (
	"go-zero-dandan/pkg/strd"
	"net/http"
	"time"
)

// Authentication 定义鉴权的入口
type Authentication interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
	UserId(r *http.Request) int64
}

// authenication鉴权的一个默认实现
type authenication struct {
}

func (*authenication) Auth(w http.ResponseWriter, r *http.Request) bool {
	return true
}
func (*authenication) UserId(r *http.Request) int64 {
	query := r.URL.Query()
	if query != nil && query["userId"] != nil {
		return strd.Int64(query["userId"][0])
	}
	return time.Now().UnixMilli()
}
