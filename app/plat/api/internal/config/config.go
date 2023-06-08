package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}

var Conf Config
