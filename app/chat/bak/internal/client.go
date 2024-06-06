package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/im/bak/internal/types"
	"go-zero-dandan/app/user/rpc/types/pb"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// 写入消息到对端的超时等待时间，如果在该时间内无法写入消息，连接可能会被认为已失效。
	writeWait = 10 * time.Second

	// 心跳等待时间，如果在该时间内没有收到 pong 消息，连接可能会被认为已失效。
	pongWait = 60 * time.Second

	// 发送 ping 消息给对端的时间间隔。它应该小于 pongWait 的值，以确保在超时之前发送新的 ping 消息。
	pingPeriod = (pongWait * 9) / 10

	// 允许的最大消息大小
	maxMessageSize = 512

	// 发送缓冲区的大小
	bufSize = 256
)

var (
	newline = []byte{'\n'} // 换行符的字节表示
	space   = []byte{' '}  // 空格的字节表示
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有的源，不加这个的话，自带的例子不让连接上
		return true
	},
}

// 管理连接关系
var clientMap = make(map[int64]*Client, 0)

// 读写锁
var rwLocker sync.RWMutex

// Client 是 WebSocket 连接 和 Hub 之间的 连接客户端。
type Client struct {
	hub *Hub

	// socket连接
	conn *websocket.Conn
	// 绑定的用户id
	userId int64
	// 缓冲的出站消息管道
	send chan []byte
}

// readPump 从 WebSocket 连接将消息传输到 Hub。
// 应用程序在每个连接上的 goroutine 中运行 readPump。应用程序通过在此 goroutine 中执行所有读取操作，确保每个连接最多只有一个读取器。
func (c *Client) readPump() {
	defer func() {
		// 在函数结束时将客户端从 hub.unregister 通道中注销
		c.hub.unregister <- c
		// 关闭客户端连接
		c.conn.Close()
	}()
	// 设置客户端连接的最大读取消息大小
	c.conn.SetReadLimit(maxMessageSize)
	// 设置客户端连接的读取截止时间为当前时间加上 pongWait
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// 设置客户端连接的 Pong 处理函数
	// 在接收到 Pong 时，更新读取截止时间为当前时间加上 pongWait
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		// 从客户端连接读取消息
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			// 如果读取消息时发生意外关闭或异常关闭的错误，记录错误日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// 去除消息中的空白字符和换行符
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// 将处理后的消息通过 hub.broadcast 通道广播给所有客户端
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
// 每个连接都会启动一个运行 writePump 的 goroutine。应用程序通过在此 goroutine 中执行所有写操作，确保每个连接最多只有一个写入器。
func (c *Client) writePump() {
	// 创建一个定时器，心跳监测
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		// 在函数结束时停止定时器
		ticker.Stop()
		// 关闭与客户端的连接
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send: //从消息管道获取要发送的消息
			c.conn.SetWriteDeadline(time.Now().Add(writeWait)) // 设置写入超时时间
			if !ok {
				// 如果通道已关闭，表示连接已断开
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// 获取下一个用于写入的 Writer
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 将排队的聊天消息添加到当前的 WebSocket 消息中
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C: // 定时器触发，发送 Ping 消息
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// 发送 Ping 消息
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Publish 发送redis消息
func (c *Client) Publish(ctx context.Context, channel string, msg string) error {
	fmt.Println("Publish……")
	err := c.hub.svc.Redis.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 订阅redis消息
func (c *Client) Subscribe(ctx context.Context, channel string) (string, error) {
	sub := c.hub.svc.Redis.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe……", msg.Payload)
	return msg.Payload, err
}

// SendMsg 发送消息给指定用户的方法
func (c *Client) SendMsg(toUserId int64, message *types.Message) error {
	msgByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	//todo::这里单机没问题，多机的话client不一定在这个hub上？？？？
	if client, ok := c.hub.clients[toUserId]; ok {
		client.conn.WriteMessage(websocket.TextMessage, msgByte)
	} else {
		//todo::当真的没找到，没上线，则是离线消息处理？
	}
	return nil
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	//升级socket协议
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Error(err)
		return
	}
	//创建ctx，后面redis等操作都得用
	ctx := context.Background()
	//创建连接端
	client := &Client{
		hub:  hub,                        //不知道啥用，感觉集成服务的
		conn: conn,                       //socket连接
		send: make(chan []byte, bufSize), //发送数据管道
	}
	//除url外，还可以通过Subprotocol: []string{"chat"}来接收
	//前端用 let sock = new WebSocket('ws://example.com', ['chat', $token])
	token := r.URL.Query().Get("token")
	user, err := hub.svc.UserRpc.GetUserByToken(ctx, &pb.TokenReq{
		Token: token,
	})
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, []byte("登录验证失败")) //CloseMessage这个会进onerror
		//conn.Close() //这个直接断开连接，前端触发onclose
		return
	}
	if existClient, ok := hub.clients[user.Id]; ok {
		existClient.conn.WriteMessage(websocket.CloseMessage, []byte("您的账号在其他设备登录"))
	}

	client.userId = user.Id
	client.hub.register <- client //这里不知道啥意义

	//发送消息协程
	go client.writePump()
	//获取消息协程
	go client.readPump()
	//通知登录成功
	hub.broadcast <- []byte("登录成功饿了")
	//conn.WriteMessage(websocket.TextMessage, []byte("登录成功"))
}
