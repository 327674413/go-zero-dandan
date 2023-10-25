package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-dandan/app/message/mq/internal/config"
	"go-zero-dandan/app/message/mq/internal/logic"
	"go-zero-dandan/app/message/mq/internal/svc"
)

var configFile = flag.String("f", "etc/message-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	mqServices := logic.Consumers(ctx, svcCtx)
	for _, mq := range mqServices {
		serviceGroup.Add(mq)
	}
	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Started %d mq service \n", len(mqServices))
	serviceGroup.Start()
}
