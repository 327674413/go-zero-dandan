package websocketd

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/singleflight"
	"net/url"
	"sync"
)

// Client 给后端用的ws客户端，目前是给mq往ws发消息用的
type Client interface {
	Close() error
	Send(v any) error
	Read(v any) error
}

var g singleflight.Group

type client struct {
	*websocket.Conn
	host string
	opt  dailOption
	mu   sync.Mutex
}

func NewClient(host string, opts ...DailOptions) *client {
	opt := newDailOptions(opts...)
	c := &client{
		host: host,
		opt:  opt,
	}
	conn, err := c.dial()
	if err != nil {
		panic(err)
	}
	c.Conn = conn
	return c

}
func (t *client) dial() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: t.host, Path: t.opt.patten}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), t.opt.header)
	if err != nil {
		logx.Error("ws连接失败")
	}
	t.Conn = conn
	go t.listen()
	return conn, err
}

func (t *client) Send(v any) error {
	var err error
	if t.Conn == nil {
		logx.Debug("ws未连接，开始重连")
		_, err, _ := g.Do("ws-dial", func() (any, error) {
			return t.dial()
		})
		if err != nil {
			logx.Error("ws连接失败")
			return errors.New("ws未连接，发送失败")
		}

	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = t.WriteMessage(websocket.TextMessage, bytes)
	if err == nil {
		return nil
	}

	// 如果服务端重启断开，前面的判断就会触发，不知道这里事后是否还有必要重里连
	//if t.Conn, err = t.dial(); err != nil {
	//	return err
	//}
	//return t.WriteMessage(websocket.TextMessage, bytes)
	return nil
}
func (t *client) Read(v any) error {
	_, msg, err := t.Conn.ReadMessage()
	if err != nil {
		return err
	}
	return json.Unmarshal(msg, v)
}
func (t *client) listen() {
	for {
		_, _, err := t.Conn.ReadMessage()
		if err != nil {
			t.Conn = nil
			logx.Error("ws连接断开")
			//因为上面有重新执行，所以这里直接return结束
			return
		}
	}
}
