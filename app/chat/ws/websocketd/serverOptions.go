package websocketd

import "time"

type ServerOptions func(opt *serverOption)
type serverOption struct {
	Authentication
	ack               AckType       //ack模式
	ackTimeout        time.Duration //ack等待超时时间
	patten            string
	maxConnectionIdle time.Duration
}

func newServerOptions(opts ...ServerOptions) serverOption {
	o := serverOption{
		Authentication:    new(authenication),
		patten:            "/ws",
		maxConnectionIdle: defaultMaxConnectionIdle,
		ackTimeout:        defaultAckTimeout,
	}
	//执行传入的options类型的函数
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
func WithServerAuthentication(auth Authentication) ServerOptions {
	return func(opt *serverOption) {
		opt.Authentication = auth
	}
}
func WithServerPatten(patten string) ServerOptions {
	return func(opt *serverOption) {
		opt.patten = patten
	}
}
func WithServerAck(ack AckType) ServerOptions {
	return func(opt *serverOption) {
		opt.ack = ack
	}
}
func WithServerMaxConnectionIdle(maxConnectionIdle time.Duration) ServerOptions {
	return func(opt *serverOption) {
		if maxConnectionIdle > 0 {
			opt.maxConnectionIdle = maxConnectionIdle
		}
	}
}
