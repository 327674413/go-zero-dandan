package test

import (
	"bufio"
	"go-zero-dandan/cmd/nat/helper"
	"io"
	"log"
	"net"
	"sync"
	"testing"
)

const (
	ControlServerAddr = "0.0.0.0:7070"
	RequestServerAddr = "0.0.0.0:7071"
	KeepAliveStr      = "KeepAlive\n"
)

var wg sync.WaitGroup
var clientConn *net.TCPConn

// 服务端
func TestUserServer(t *testing.T) {
	wg.Add(1)
	//监听控制中心
	go ControlServer()
	//监听用户请求
	go RequestServer()
	wg.Wait()
}
func ControlServer() {
	tcpListener, err := helper.CreateListen(ControlServerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("ControlServer is listen on %s\n", ControlServerAddr)
	for {
		clientConn, err = tcpListener.AcceptTCP()
		if err != nil {
			return
		}
		go helper.KeepAlive(clientConn, KeepAliveStr)
	}
}
func RequestServer() {
	tcpListener, err := helper.CreateListen(RequestServerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("RequestServer is listen on %s\n", RequestServerAddr)
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			return
		}
		go io.Copy(clientConn, conn)
		go io.Copy(conn, clientConn)
	}
}
func TestUserClient(t *testing.T) {
	conn, err := helper.CreateConn(ControlServerAddr)
	if err != nil {
		log.Printf("[连接失败]：%s", err)
	}
	for {
		s, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("[获取数据]error：%s", err)
			continue
		}
		log.Printf("[获取数据]：%v", s)
	}
}
func TestUserRequestClient(t *testing.T) {
	conn, err := helper.CreateConn(RequestServerAddr)
	if err != nil {
		log.Printf("[连接失败]：%s", err)
	}
	_, err = conn.Write([]byte("客户单发送的\n"))
	if err != nil {
		log.Printf("[发送失败]：%s", err)
	}
}
