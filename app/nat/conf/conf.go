package conf

const (
	KeepAliveStr      = "KeepAlive\n"
	NewConnection     = "NewConnection\n"
	LocalServerAddr   = ":8800"                //本地服务端口,应该是给客户端用的，服务端应该不需要
	ControlServerAddr = "124.222.198.229:7000" //控制器端口，不知道啥用，但双端都要用，服务端不要ip，是本地，比如:7000，客户端要写服务器ip
	TunnelServerAddr  = "124.222.198.229:7010" //隧道端口，通控制器
	UserRequestAddr   = ":7020"                //服务端要监听转发的端口，也是服务端进这个端口的东西，都会转发到本地的LocalServerAddr上

	BufSize = 1024
)
