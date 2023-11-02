package main

import (
	"flag"
	"fmt"
    "github.com/zeromicro/go-zero/core/logx"
	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

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
    	}), rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
    logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
