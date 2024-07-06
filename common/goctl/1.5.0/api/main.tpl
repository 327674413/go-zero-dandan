package main

import (
	"flag"
	"fmt"
	"net/http"
	"encoding/json"
    "github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	var err error
    resd.Mode = c.Mode
    resd.I18n, err = resd.NewI18n(&resd.I18nConfig{
        LangPathList: c.I18n.Langs,
        DefaultLang:  c.I18n.Default,
    })
    if err != nil {
        panic(err)
    }
    server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(200)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "code":   resd.AuthPlatErr,
            "result": false,
            "msg":    resd.I18n.NewLang(r.FormValue("lang")).Msg(resd.AuthPlatErr),
        })
    }), rest.WithCustomCors(nil, func(w http.ResponseWriter) {
        //跨域处理
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
        //w.Header().Set("Access-Control-Allow-Credentials", "true") //允许传输cookies
    }, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
    if (c.Mode != "prod") {
        logx.DisableStat() //去掉定时出现的控制台打印
    }
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
