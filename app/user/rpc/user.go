package main

import (
	"flag"
	"fmt"
	"go-zero-dandan/common/interceptor"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/user/rpc/internal/config"
	"go-zero-dandan/app/user/rpc/internal/server"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	var err error
	resd.Mode = c.Mode
	resd.I18n, err = resd.NewI18n(&resd.I18nConfig{
		LangPathList: c.I18n.Langs,
		DefaultLang:  c.I18n.Default,
	})
	if err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userRpc.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(interceptor.RpcServerInterceptor())
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
