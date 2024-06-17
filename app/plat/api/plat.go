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
	err := configServer.NewConfigServer(*configFile, configServer.NewSail(&configServer.Config{
		ETCDEndpoints:  "127.0.0.1:2379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "plat",
		Configs:        "plat-api.yaml",
		ConfigFilePath: "./etc/conf",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		var c config.Config
		configServer.LoadFromJsonBytes(bytes, &c)

		proc.WrapUp()
		wg.Add(1)
		go func(c config.Config) {
			defer wg.Done()

			Run(c)
		}(c)
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
