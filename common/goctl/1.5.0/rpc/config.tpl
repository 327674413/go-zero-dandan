package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	I18n struct {
        Default string
        Langs   []string
    }
}
