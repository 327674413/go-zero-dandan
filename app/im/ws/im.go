package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/ws/internal/config"
	"go-zero-dandan/app/im/ws/internal/handler"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/websocketd"
	"time"
)

//var rpcConfigFile = flag.String("f", "etc/im-dev.yaml", "rpc rpcconfig file")

var configFile = flag.String("f", "etc/dev/im.yaml", "config file")

func ws() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}
	svcCtx := svc.NewServiceContext(c)
	//todo::参照api和rpc的方式封装进去
	server := websocketd.NewServer(
		c.ListenOn,
		websocketd.WithServerAuthentication(handler.NewUserAuth(svcCtx)),
		websocketd.WithServerMaxConnectionIdle(10*time.Second), // 超时自动断开时间
		websocketd.WithServerAck(websocketd.NoAck),
	)

	defer server.Stop()

	handler.RegisterHandlers(server, svcCtx)
	fmt.Printf("Starting websocket server at %s...\n", c.ListenOn)
	server.Start()
}

func rpc() {
	//flag.Parse()
	//
	//var c config.Config
	//conf.MustLoad(*rpcConfigFile, &c)
	//ctx := svc.NewServiceContext(c)
	//
	//s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
	//	pb.RegisterImServer(grpcServer, server.NewImServer(ctx))
	//
	//	if c.Mode == service.DevMode || c.Mode == service.TestMode {
	//		reflection.Register(grpcServer)
	//	}
	//})
	//defer s.Stop()
	//
	//fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	//s.Start()
}

func main() {
	logx.DisableStat() //去掉定时出现的控制台打印
	ws()
}
