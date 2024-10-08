package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	DB        struct {
		DataSource string
	}
	I18n struct {
		Default string
		Langs   []string
	}
}
