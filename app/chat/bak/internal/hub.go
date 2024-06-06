package internal

import (
	"go-zero-dandan/app/im/bak/internal/svc"
)

// Hub 维护在线客户端的集合
type Hub struct {
	// 注册的客户端。
	clients map[int64]*Client

	// 从客户端接收的入站消息，暂定为全局广播的管道
	broadcast chan []byte

	// 客户端的新连接管道
	register chan *Client

	// 客户端的断开管道
	unregister chan *Client

	// 服务中间件
	svc *svc.ServiceContext
}

func NewHub(svcCtx *svc.ServiceContext) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[int64]*Client),
		svc:        svcCtx,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: // 接收客户端的新连接请求
			h.clients[client.userId] = client // 将客户端添加到注册列表中
		case client := <-h.unregister:
			if _, ok := h.clients[client.userId]; ok {
				delete(h.clients, client.userId)
				close(client.send)
			}
		case message := <-h.broadcast:
			for userId := range h.clients {
				select {
				case h.clients[userId].send <- message:
				default:
					close(h.clients[userId].send)
					delete(h.clients, userId)
				}
			}
		}
	}
}
