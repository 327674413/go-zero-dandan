package main

import (
	"encoding/json"
	"go-zero-dandan/cmd/nat/conf"
	"go-zero-dandan/cmd/nat/helper"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
)

var controlConn *net.TCPConn
var userConn *net.TCPConn
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// 控制中心监听
	go controlListen()
	// 用户请求的监听
	go userRequestListen()
	// 隧道监听
	go tunnelListen()
	// 启动Web服务
	go runWeb()
	wg.Wait()
}

func controlListen() {
	tcpListener, err := helper.CreateListen(conf.ControlServerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("[控制中心] 监听中：%v\n", tcpListener.Addr().String())
	for {
		controlConn, err = tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("ControlListen AcceptTCP Error:%v\n", err)
			return
		}
		go helper.KeepAlive(controlConn)
	}
}

func userRequestListen() {
	tcpListener, err := helper.CreateListen(conf.UserRequestAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("[用户请求] 监听中：%v\n", tcpListener.Addr().String())
	for {
		userConn, err = tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("UserRequestListen AcceptTCP Error:%v\n", err)
			return
		}
		// 发送消息，告诉客户端有新的连接
		_, err := controlConn.Write([]byte(conf.NewConnection))
		if err != nil {
			log.Printf("发送失败: %v", err)
		}
	}
}

func tunnelListen() {
	tcpListener, err := helper.CreateListen(conf.TunnelServerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("[隧道] 监听中：%v\n", tcpListener.Addr().String())
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("unnelListen AcceptTCP Error:%v\n", err)
			return
		}
		// 数据转发
		go io.Copy(userConn, conn)
		go io.Copy(conn, userConn)
	}
}

func runWeb() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		b, err := json.Marshal(q)
		if err != nil {
			log.Printf("Marshal Error:%s", err)
		}
		writer.Write(b)
	})

	log.Println("本地web已启动", conf.LocalServerAddr)
	http.ListenAndServe(conf.LocalServerAddr, nil)
}
