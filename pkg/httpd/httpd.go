package httpd

import (
	"net/http"
	"strings"
)

// GetClientIP 获取客户端IP地址
func GetClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	ip := ""
	if xff != "" {
		// 处理可能存在的多个 IP 地址，取第一个
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			ip = strings.TrimSpace(ips[0])
		}

	}
	if ip == "" {
		// r.RemoteAddr 可能包含 IP 地址和端口，去掉端口
		ips := strings.Split(r.RemoteAddr, ":")
		if len(ips) > 0 {
			ip = ips[0]
		}
	}
	// postman这种本地请求，ip是[::1]:端口，误截取
	if ip == "" || len(ip) < 4 {
		return "0.0.0.0"
	}
	return ip
}

// GetClientKey 获取客户端标识,目前用ip 和 user-agent
func GetClientKey(r *http.Request) string {
	ip := GetClientIP(r)
	ua := r.Header.Get("User-Agent")
	return ip + ":" + ua
}
