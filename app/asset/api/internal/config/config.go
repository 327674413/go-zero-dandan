package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	AssetPath struct {
		File  string
		Img   string
		Audio string
		Video string
	}
	DB struct {
		DataSource string
	}
}

var Conf Config
