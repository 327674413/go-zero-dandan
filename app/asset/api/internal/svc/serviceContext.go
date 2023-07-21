package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-dandan/app/asset/api/internal/config"
	"go-zero-dandan/app/asset/api/internal/middleware"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/storaged"
)

type ServiceContext struct {
	Config         config.Config
	LangMiddleware rest.Middleware
	SqlConn        sqlx.SqlConn
	Mode           string
	Minio          *minio.Client
	Storage        storaged.InterfaceStorage
	TxCos          *cos.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:         c,
		SqlConn:        sqlx.NewMysql(c.DB.DataSource),
		Mode:           c.Mode,
		LangMiddleware: middleware.NewLangMiddleware().Handle,
	}
	var err error
	if c.AssetMode == constd.AssetModeLocal {
		svc.Storage, err = storaged.NewStorage(&storaged.StorageConfig{
			Provider:  storaged.ProviderLocal,
			LocalPath: c.LocalPath,
		})
		if err != nil {
			panic(err)
		}
	} else if c.AssetMode == constd.AssetModeMinio {
		svc.Storage, err = storaged.NewStorage(&storaged.StorageConfig{
			Provider: storaged.ProviderMinio,
			Endpoint: c.Minio.Address,
			Key:      c.Minio.AccessKey,
			Secret:   c.Minio.SecretKey,
			Bucket:   c.Minio.Bucket,
		})
	} else if c.AssetMode == constd.AssetModeAliOss {
		svc.Storage, err = storaged.NewStorage(&storaged.StorageConfig{
			Provider: storaged.ProviderAliOss,
			Endpoint: c.AliOss.PublicBucketAddr,
			Key:      c.AliOss.AccessKeyId,
			Secret:   c.AliOss.AccessKeySecret,
			Bucket:   c.AliOss.Bucket,
		})
	} else if c.AssetMode == constd.AssetModeTxCos {
		svc.Storage, err = storaged.NewStorage(&storaged.StorageConfig{
			Provider: storaged.ProviderTxCos,
			Endpoint: c.TxCos.PublicBucketAddr,
			Key:      c.TxCos.SecretKey,
			Secret:   c.TxCos.SecretId,
			Bucket:   c.TxCos.Bucket,
		})
	} else {
		panic("不支持的文件管理类型")
	}
	return svc
}
