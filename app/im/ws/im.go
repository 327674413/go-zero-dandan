package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/ws/internal/config"
	"go-zero-dandan/app/im/ws/internal/server"
	"go-zero-dandan/app/im/ws/internal/svc"
	"go-zero-dandan/app/im/ws/internal/wsConfig"
	"go-zero-dandan/app/im/ws/internal/wsHandler"
	"go-zero-dandan/app/im/ws/internal/wsSvc"
	"go-zero-dandan/app/im/ws/types/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
)

var wsConfigFile = flag.String("w", "etc/im-ws.yaml", "ws rpcconfig file")
var rpcConfigFile = flag.String("f", "etc/im.yaml", "rpc rpcconfig file")

func ws() {
	flag.Parse()

	var c wsConfig.WsConfig
	conf.MustLoad(*wsConfigFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := wsSvc.NewServiceContext(c)
	wsHandler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting websocket server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func rpc() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*rpcConfigFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterImServer(grpcServer, server.NewImServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func main() {
	logx.DisableStat() //去掉定时出现的控制台打印
	go ws()
	logx.Info("ws 启动成功 等待1秒启动 rpc")
	time.Sleep(time.Second)
	rpc()
}
