package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/plat/rpc/internal/config"
	"go-zero-dandan/app/plat/rpc/internal/server"
	"go-zero-dandan/app/plat/rpc/internal/svc"
	"go-zero-dandan/app/plat/rpc/types/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/plat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//err := configServer.NewConfigServer(*configFile, configServer.NewSail(&configServer.Config{
	//	ETCDEndpoints:  "127.0.0.1:2379",
	//	ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
	//	Namespace:      "plat",
	//	Configs:        "plat-api.yaml",
	//	ConfigFilePath: "./etc/conf",
	//	LogLevel:       "DEBUG",
	//})).MustLoad(&c)
	//if err != nil {
	//	panic(err)
	//}
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPlatServer(grpcServer, server.NewPlatServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
