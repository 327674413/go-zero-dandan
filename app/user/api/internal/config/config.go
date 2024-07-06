package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Conf struct {
		LoginTokenExSec int
	}
	I18n struct {
		Default string
		Langs   []string
	}
	MessageRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	RedisConf  redis.RedisConf
}

var Conf Config
