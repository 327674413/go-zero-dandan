package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	AssetMode int64
	AssetPath struct {
		File  string
		Img   string
		Audio string
		Video string
	}
	Minio struct {
		Address   string
		AccessKey string
		SecretKey string
	}
	DB struct {
		DataSource string
	}
}

var Conf Config
