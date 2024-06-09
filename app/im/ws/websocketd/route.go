package websocketd

type Route struct {
	Method  string
	Handler HandlerFunc
}
type HandlerFunc func(server *Server, conn *Conn, msg *Message)
