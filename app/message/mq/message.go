package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/message/mq/internal/server"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-dandan/app/message/mq/internal/config"
	"go-zero-dandan/app/message/mq/internal/svc"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	// 用于对自己用service起的服务，启动log、prometheus、trace、metricsUrl等，需要在config里增加service.ServiceConf
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	consumers := server.Consumers(c, ctx, svcCtx)
	for _, mq := range consumers {
		serviceGroup.Add(mq)
	}
	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Started %d mq service \n", len(consumers))
	serviceGroup.Start()
}
