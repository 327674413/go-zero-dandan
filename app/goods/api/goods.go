package main

import (
	"flag"
	"fmt"
	"go-zero-dandan/app/goods/api/internal/config"
	"go-zero-dandan/app/goods/api/internal/handler"
	"go-zero-dandan/app/goods/api/internal/svc"
	"go-zero-dandan/common/resd"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/goods-api.yaml", "the config file")

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
		resp := resd.NewResp(r.Context(), resd.I18n.NewLang(r.FormValue("lang")))
		resd.ApiFail(w, r, resp.NewErr(resd.ErrAuthPlat))
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
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
