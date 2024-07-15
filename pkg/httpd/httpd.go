package httpd

import "net/http"

// GetClientIP 获取客户端IP地址
func GetClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		return xff
	}
	return r.RemoteAddr
}

// GetClientKey 获取客户端标识,目前用ip 和 user-agent
func GetClientKey(r *http.Request) string {
	ip := GetClientIP(r)
	ua := r.Header.Get("User-Agent")
	return ip + ":" + ua
}
