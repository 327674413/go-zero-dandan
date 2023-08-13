package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/asset/api/internal/config"
	"go-zero-dandan/app/asset/api/internal/handler"
	"go-zero-dandan/app/asset/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/asset-api-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		// 将错误对象转换为 JSON 格式，并写入响应
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":   401,
			"result": false,
			"msg":    err.Error(),
		})
	}), rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		//跨域处理
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Headers", "*") // 这个好像可以不要，已经哟过token作为header解决
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
