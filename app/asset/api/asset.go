package main

import (
	"flag"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/config"
	"go-zero-dandan/app/asset/api/internal/handler"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/common/resd"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/asset-api-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	var err error
	conf.MustLoad(*configFile, &c)
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
		resd.ApiFail(w, r, resp.NewErr(resd.ErrAuthPlat))
	}), rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		//跨域处理
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
