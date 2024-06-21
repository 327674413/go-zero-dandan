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
var restartCh = make(chan config.Config)
var shutdownCh = make(chan struct{})
var once sync.Once

func main() {
	flag.Parse()

	go watchConfigChanges()

	var c config.Config
	// Initial load of the configuration
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
		restartCh <- newConf
		return nil
	})
	if err != nil {
		panic(err)
	}

	restartCh <- c
	select {}
}

func watchConfigChanges() {
	var wg sync.WaitGroup

	for c := range restartCh {
		wg.Add(1)
		go func(conf config.Config) {
			defer wg.Done()
			runServer(conf)
		}(c)

		// Wait for the old server to shutdown
		<-shutdownCh
	}

	wg.Wait()
}

func runServer(c config.Config) {
	server := rest.MustNewServer(c.RestConf)
	defer func() {
		server.Stop()
		shutdownCh <- struct{}{}
	}()

	once.Do(func() {
		proc.WrapUp()
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat() // Disabling periodic console logging
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
