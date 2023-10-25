package test

import (
	"net"
	"testing"
)

const (
	addr    = "0.0.0.0:7777"
	bufSize = 10
)

// 监听
func TestTcpListen(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}
		//读数据
		for {
			var buf [bufSize]byte
			n, err := tcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(buf[:n]))
			if n < bufSize {
				break
			}
		}
		//写数据
		tcpConn.Write([]byte("hellow 你好"))
	}
}

// 创建链接
func TestCreateTcp(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Fatal(err)
	}
	//写入数据
	tcpConn.Write([]byte("client=====>发出 "))
	//读数据
	for {
		var buf [bufSize]byte
		n, err := tcpConn.Read(buf[:])
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(buf[:n]))
		if n < bufSize {
			break
		}
	}
}
