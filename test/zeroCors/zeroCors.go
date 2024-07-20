package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func main() {
	var config = rest.RestConf{}
	config.Name = "test-api"
	config.Port = 8080
	config.Host = "0.0.0.0"
	server := rest.MustNewServer(config, rest.WithCustomCors(func(header http.Header) {
		//跨域处理
		header.Set("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Headers", "AppToken,UserToken")
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		header.Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
	}, nil, "*"))
	defer server.Stop()
	server.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodPost,
				Path:   "/test",
				Handler: func(w http.ResponseWriter, r *http.Request) {
					logx.Info("请求进来了")
					httpx.OkJsonCtx(r.Context(), w, "1")
				},
			},
		},
	)
	fmt.Printf("Starting server at %s:%d...\n", config.Host, config.Port)
	server.Start()
}
