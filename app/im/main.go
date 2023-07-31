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

/*
import (

	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/im/internal"
	"net/http"

)

var (

	port    = flag.Int("port", 3333, "the port to listen")
	timeout = flag.Int64("timeout", 0, "timeout of milliseconds")
	cpu     = flag.Int64("cpu", 500, "cpu threshold")
	host    = "localhost"

)

//var configFile = flag.String("f", "etc/im-server.yaml", "the config file")

	func main() {
		flag.Parse()
		//var c config.Config
		//conf.MustLoad(*configFile, &c)
		logx.Disable()
		//ctx := svc.NewServiceContext(c)
		engine := rest.MustNewServer(rest.RestConf{
			ServiceConf: service.ServiceConf{
				Log: logx.LogConf{
					Mode: "console",
				},
			},
			Host:         host,
			Port:         *port,
			Timeout:      *timeout,
			CpuThreshold: *cpu,
		})
		defer engine.Stop()

		hub := internal.NewHub(nil)
		go hub.Run()
		fmt.Printf("websocket runing at %s:%d\n", host, *port)
		engine.Start()
		engine.AddRoute(rest.Route{
			Method: http.MethodGet,
			Path:   "/",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("请求进来了1")
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
		engine.AddRoute(rest.Route{
			Method: http.MethodGet,
			Path:   "/ws",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("请求进来了2")
				internal.ServeWs(hub, w, r)
			},
		})

}
*/
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
	ctx := svc.NewServiceContext(c)
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
	hub := internal.NewHub(ctx)
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
			internal.ServeWs(hub, w, r)
		},
	})

	engine.Start()
}
