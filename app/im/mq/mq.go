package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"go-zero-dandan/app/im/mq/internal/config"
	"go-zero-dandan/app/im/mq/internal/handler"
	"go-zero-dandan/app/im/mq/internal/svc"
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
	//参照zero方式启动mq消费者
	listen := handler.NewListen(svcCtx)
	serviceGroup := service.NewServiceGroup()
	for _, s := range listen.Services() {
		//循环handler的listenh中配置的消费者，进行加载
		serviceGroup.Add(s)
	}
	fmt.Printf("Starting chat mq server at %s...\n", c.ListenOn)
	//应该是gozero的kq包封装好的，这样就会自动开启加载的消费者了
	serviceGroup.Start()
}
