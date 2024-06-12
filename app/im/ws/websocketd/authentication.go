package websocketd

import (
	"fmt"
	"net/http"
	"time"
)

// Authentication 定义鉴权的入口
type Authentication interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
	UserId(r *http.Request) string
}

// authenication鉴权的一个默认实现
type authenication struct {
}

func (*authenication) Auth(w http.ResponseWriter, r *http.Request) bool {
	return true
}
func (*authenication) UserId(r *http.Request) string {
	query := r.URL.Query()
	if query != nil && query["userId"] != nil {
		return query["userId"][0]
	}
	return fmt.Sprintf("%d", time.Now().UnixMilli())
}
