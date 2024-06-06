package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-dandan/app/chat/mq/internal/config"
	"go-zero-dandan/app/chat/mq/internal/handler"
	"go-zero-dandan/app/chat/mq/internal/svc"
	"go-zero-dandan/common/constd"
)

var configFile = flag.String("f", "etc/dev/mq.yaml", "config file")

func main() {

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}
	svcCtx := svc.NewServiceContext(c)
	listen := handler.NewListen(svcCtx)
	serviceGroup := service.NewServiceGroup()
	for _, s := range listen.Services() {
		serviceGroup.Add(s)
	}
	if c.Mode == constd.ModeDev {
		logx.DisableStat() //去掉定时出现的控制台打印
	}
	fmt.Printf("Starting chat mq server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
