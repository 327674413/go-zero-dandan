package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	I18n struct {
		Default string
		Langs   []string
	}
	UserRpc zrpc.RpcClientConf
	ImRpc   zrpc.RpcClientConf
	DB      struct {
		DataSource string
	}
	RedisConf redis.RedisConf
}
