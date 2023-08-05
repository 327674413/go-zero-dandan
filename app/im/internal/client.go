package internal

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/types/pb"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	// send buffer size
	bufSize = 256
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
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

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
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
	if existClient, ok := clientMap[user.Id]; ok {
		existClient.conn.WriteMessage(websocket.CloseMessage, []byte("您的账号在其他设备登录"))
	}

	rwLocker.Lock()
	clientMap[user.Id] = client
	rwLocker.Unlock()

	client.hub.register <- client //这里不知道啥意义

	//发送消息协程
	go client.writePump()
	//获取消息协程
	go client.readPump()
}
