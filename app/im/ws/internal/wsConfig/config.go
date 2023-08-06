package wsConfig

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type WsConfig struct {
	rest.RestConf
	UserRpc          zrpc.RpcClientConf
	RedisConf        redis.RedisConf
	SendMsgRateLimit RateLimitConfig
	WebsocketConfig  WebsocketConfig
}

var WsConf WsConfig

type RateLimitConfig struct {
	Enable  bool
	Seconds int
	Quota   int
}

type WebsocketConfig struct {
	MaxConnNum int
	TimeOut    int
	MaxMsgLen  int
}
