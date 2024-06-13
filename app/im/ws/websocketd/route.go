package websocketd

// Route 会注册到server中的routes里，在执行写入消息方法的时候，根据这个里的method，执行对应的HandlerFunc处理方法
type Route struct {
	Method  string
	Handler HandlerFunc
}

type HandlerFunc func(server *Server, conn *Conn, msg *Message)
