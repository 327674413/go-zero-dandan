package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	AssetMode int64
	LocalPath string
	TxCos     struct {
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
		Address   string
		AccessKey string
		SecretKey string
		Bucket    string
	}

	DB struct {
		DataSource string
	}
}

var Conf Config
