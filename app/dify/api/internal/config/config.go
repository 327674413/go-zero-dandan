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
	I18n struct {
		Default string
		Langs   []string
	}
	RedisConf redis.RedisConf
	UserRpc   zrpc.RpcClientConf
	Auth      struct {
		AccessSecret string
		AccessExpire int64
	}
	ReqRateLimitByIpAgent struct {
		Seconds   int
		Quota     int
		KeyPrefix string
	}
	Dify struct {
		Url    string
		AppKey string
	}
}
