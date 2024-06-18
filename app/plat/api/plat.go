package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/plat/api/internal/config"
	"go-zero-dandan/app/plat/api/internal/handler"
	"go-zero-dandan/app/plat/api/internal/svc"
	"go-zero-dandan/common/configServer"
	"sync"
)

var configFile = flag.String("f", "etc/plat-api-dev.yaml", "the config file")
var wg sync.WaitGroup

func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)
	//Run(c)

	err := configServer.NewConfigServer(*configFile, configServer.NewSail(&configServer.Config{
		ETCDEndpoints:  "127.0.0.1:2379",
		ProjectKey:     "2-public",
		Namespace:      "plat",
		Configs:        "plat-api.yaml",
		ConfigFilePath: "./etc/conf",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		logx.Info("配置发生变化，重新加载")
		var newConf config.Config
		configServer.LoadFromJsonBytes(bytes, &newConf)
		proc.WrapUp()
		wg.Add(1)
		go func(newConf config.Config) {
			defer wg.Done()
			Run(newConf)
		}(newConf)
		return nil
	})
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func(c config.Config) {
		defer wg.Done()
		Run(c)
	}(c)
	wg.Wait()
}
func Run(c config.Config) {
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
