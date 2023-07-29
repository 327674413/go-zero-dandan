package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	AssetMode int64
	Local     struct {
		Path             string
		Bucket           string
		PublicBucketAddr string
	}
	TxCos struct {
		SecretKey        string
		SecretId         string
		PublicBucketAddr string
		Bucket           string
	}
	AliOss struct {
		AccessKeySecret  string
		AccessKeyId      string
		PublicBucketAddr string
		Bucket           string
	}
	Minio struct {
		PublicBucketAddr string
		AccessKey        string
		SecretKey        string
		Bucket           string
	}
	UserRpc   zrpc.RpcClientConf
	RedisConf redis.RedisConf
	DB        struct {
		DataSource string
	}
}

var Conf Config
