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
        resp := resd.NewResp(r.Context(), r.FormValue("lang"))
        if err.Error() == "Token is expired" {
            resd.ApiFail(w, r, resp.NewErr(resd.ErrAuthPlatExpired))
        } else {
            resd.ApiFail(w, r, resp.NewErr(resd.ErrAuthPlat))
        }
    }), rest.WithCustomCors(func(header http.Header) {
        //跨域处理
        header.Set("Access-Control-Allow-Origin", "*")
        header.Add("Access-Control-Allow-Headers", "UserToken")
        header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        header.Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
    }, nil, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
