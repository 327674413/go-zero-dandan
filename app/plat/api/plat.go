package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/plat/api/global"
	"go-zero-dandan/app/plat/api/internal/config"
	"go-zero-dandan/app/plat/api/internal/handler"
	"go-zero-dandan/app/plat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/plat-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	global.Config = c //自定义应用内全局的配置参数，在其他文件里使用
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
