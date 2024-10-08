package helper

import (
	"go-zero-dandan/cmd/nat/conf"
	"log"
	"net"
	"time"
)

// CreateListen 监听
func CreateListen(serverAddr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	return tcpListen, err
}

// CreateConn 创建连接
func CreateConn(serverAddr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	return conn, err
}

// KeepAlive 设置连接保活
func KeepAlive(conn *net.TCPConn) {
	for {
		_, err := conn.Write([]byte(conf.KeepAliveStr))
		if err != nil {
			log.Printf("[KeepAlive] Error %s", err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}

// GetDataFromConnection 获取Connection中的数据
func GetDataFromConnection(bufSize int, conn *net.TCPConn) ([]byte, error) {
	b := make([]byte, 0)
	for {
		buf := make([]byte, bufSize)
		n, err := conn.Read(buf)
		if err != nil {
			return nil, err
		}
		b = append(b, buf[:n]...)
		if n < bufSize {
			break
		}
	}
	return b, nil
}
