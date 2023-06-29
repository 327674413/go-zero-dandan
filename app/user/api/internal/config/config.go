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
	MessageRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	RedisConf  redis.RedisConf
}

var Conf Config
