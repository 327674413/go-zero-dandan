package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mongo struct {
		Url string
		Db  string
	}
	MsgSysTransfer struct {
		Topic string
		Addrs []string
	}
}
