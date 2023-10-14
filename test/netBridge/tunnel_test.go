package test

import (
	"fmt"
	"net"
	"strconv"
	"testing"
)

const (
	serverAddr = "0.0.0.0:22300"
	tunnelAddr = "0.0.0.0:22200"
)

// server
func TestServer(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
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
		b := make([]byte, 0)
		//读数据
		for {
			var buf [bufSize]byte
			n, err := tcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			b = append(b, buf[:n]...)
			if n < bufSize {
				break
			}
		}
		//写数据
		i, err := strconv.Atoi(string(b))
		if err != nil {
			t.Fatal(err)
		}
		i = i + 2
		tcpConn.Write([]byte(strconv.Itoa(i)))
	}
}

// 创建链接
func TestClient(t *testing.T) {
	tunnelTcp, err := net.ResolveTCPAddr("tcp", tunnelAddr)
	if err != nil {
		t.Fatal(err)
	}
	tcpConn, err := net.DialTCP("tcp", nil, tunnelTcp)
	if err != nil {
		t.Fatal(err)
	}
	//写入数据
	tcpConn.Write([]byte("1200"))
	//读数据
	b := make([]byte, 0)
	for {
		var buf [bufSize]byte
		n, err := tcpConn.Read(buf[:])
		if err != nil {
			t.Fatal(err)
		}
		b = append(b, buf[:n]...)
		if n < bufSize {
			break
		}
	}
	fmt.Println(string(b))
}

// tunnel
func TestTunnel(t *testing.T) {
	tunnelTcp, err := net.ResolveTCPAddr("tcp", tunnelAddr)
	if err != nil {
		t.Fatal(err)
	}
	tunnelListen, err := net.ListenTCP("tcp", tunnelTcp)
	if err != nil {
		t.Fatal(err)
	}
	for {
		clientTcpConn, err := tunnelListen.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}
		//获取用户过来的数据
		b := make([]byte, 0)
		for {
			var buf [bufSize]byte
			n, err := clientTcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			b = append(b, buf[:n]...)
			if n < bufSize {
				break
			}
		}
		//与服务端创建链接
		serverTcp, err := net.ResolveTCPAddr("net", serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		serverTcpConn, err := net.DialTCP("tcp", nil, serverTcp)
		serverTcpConn.Write(b)

		//获取服务端响应过来的数据
		b2 := make([]byte, 0)
		for {
			var buf [bufSize]byte
			n, err := serverTcpConn.Read(buf[:])
			if err != nil {
				t.Fatal(err)
			}
			b2 = append(b, buf[:n]...)
			if n < bufSize {
				break
			}
		}

		clientTcpConn.Write(b2)
	}
}
