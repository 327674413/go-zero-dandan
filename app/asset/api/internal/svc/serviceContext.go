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
			Bucket:   "public",
		})
	} else if c.AssetMode == constd.AssetModeTxCos {
		svc.Storage, err = storaged.NewStorage(&storaged.StorageConfig{
			Provider: storaged.ProviderTxCos,
			Endpoint: c.TxCos.PublicBucketAddr,
			Key:      c.TxCos.SecretKey,
			Secret:   c.TxCos.SecretId,
			Bucket:   "public",
		})
		/*
			// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
				// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
				/*u, _ := url.Parse(c.TxCos.PublicBucketAddr)
				b := &cos.BaseURL{BucketURL: u}
				txCosClient := cos.NewClient(b, &http.Client{
					Transport: &cos.AuthorizationTransport{
						// 通过环境变量获取密钥
						// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
						SecretID: os.Getenv(c.TxCos.SecretId), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
						// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
						SecretKey: os.Getenv(c.TxCos.SecretKey), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
					},
				})
		*/
	}
	return svc
}
