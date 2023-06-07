package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/api/global"
	"go-zero-dandan/app/user/api/internal/config"
	"go-zero-dandan/app/user/api/internal/handler"
	"go-zero-dandan/app/user/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	global.Config = c //自定义应用内全局的配置参数，在其他文件里使用
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		// 将错误对象转换为 JSON 格式，并写入响应
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":   401,
			"result": false,
			"msg":    err.Error(),
		})
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
