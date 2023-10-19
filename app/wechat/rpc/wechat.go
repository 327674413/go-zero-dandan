package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/wechat/global"
	"go-zero-dandan/app/wechat/rpc/internal/config"
	"go-zero-dandan/app/wechat/rpc/internal/server"
	"go-zero-dandan/app/wechat/rpc/internal/svc"
	"go-zero-dandan/app/wechat/rpc/types/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/wechat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	global.Config = c //自定义应用内全局的配置参数，在其他文件里使用
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.DisableStat() //去掉定时出现的控制台打印
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
