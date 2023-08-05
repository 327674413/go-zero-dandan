package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/im/internal"
	"go-zero-dandan/app/im/internal/config"
	"go-zero-dandan/app/im/internal/svc"
	"net/http"
)

var (
	port    = flag.Int("port", 3333, "the port to listen")
	timeout = flag.Int64("timeout", 0, "timeout of milliseconds")
	cpu     = flag.Int64("cpu", 500, "cpu threshold")
	host    = "localhost"
)
var configFile = flag.String("f", "etc/im-server.yaml", "the config file")

func main() {
	flag.Parse()

	//logx.Disable()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Host:         "localhost",
		Port:         *port,
		Timeout:      *timeout,
		CpuThreshold: *cpu,
	})
	defer engine.Stop()
	hub := internal.NewHub(svcCtx)
	go hub.Run()
	fmt.Printf("websocket runing at %s:%d\n", host, *port)
	/*engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			if r.Method != "GET" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			http.ServeFile(w, r, "home.html")
		},
	})
	*/
	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			//如果不想让对方连接，可以在这里拒绝，但是websocket的onerror无法收到任何消息体，没法携带错误提示信息
			//解决方案只能先让对方连接成功，然后发消息给他告诉他失败
			internal.ServeWs(hub, w, r)
		},
	})

	engine.Start()
}
