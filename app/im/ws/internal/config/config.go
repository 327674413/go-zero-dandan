package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	UserRpc zrpc.RpcClientConf
	PlatRpc zrpc.RpcClientConf
	//RedisConf redis.RedisConf
	ListenOn string
	I18n     struct {
		Default string
		Langs   []string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mongo struct {
		Url string
		Db  string
	}
	MsgChatTransfer struct {
		Topic string
		Addrs []string
	}
	MsgReadTransfer struct {
		Topic string
		Addrs []string
	}
	Ws struct {
		SysToken string
	}
}
