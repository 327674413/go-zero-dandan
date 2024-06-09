package websocketd

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
)

type Client interface {
	Close() error
	Send(v any) error
	Read(v any) error
}

type client struct {
	*websocket.Conn
	host string
	opt  dailOption
}

func NewClient(host string, opts ...DailOptions) *client {
	opt := newDailOptions(opts...)
	c := &client{
		host: host,
		opt:  opt,
	}
	conn, err := c.dail()
	if err != nil {
		panic(err)
	}
	c.Conn = conn
	return c

}
func (t *client) dail() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: t.host, Path: t.opt.patten}
	logx.Info("connecting to ", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), t.opt.header)
	return conn, err
}
func (t *client) Send(v any) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = t.WriteMessage(websocket.TextMessage, bytes)
	if err == nil {
		return nil
	}
	//如果发失败，重连发送
	conn, err := t.dail()
	if err != nil {
		return err
	}
	t.Conn = conn
	return t.WriteMessage(websocket.TextMessage, bytes)
}
func (t *client) Read(v any) error {
	_, msg, err := t.Conn.ReadMessage()
	if err != nil {
		return err
	}
	return json.Unmarshal(msg, v)
}
