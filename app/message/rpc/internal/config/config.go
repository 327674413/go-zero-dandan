package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Db struct {
		DataSource string
	}
	RedisConf   redis.RedisConf
	ImRpc       zrpc.RpcClientConf
	KqSmsPusher struct {
		Addrs []string
		Topic string
	}
}
