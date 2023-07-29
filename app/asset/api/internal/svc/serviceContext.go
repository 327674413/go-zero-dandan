package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/asset/api/internal/config"
	"go-zero-dandan/app/asset/api/internal/middleware"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/redisd"
	"go-zero-dandan/common/storaged"
)

type ServiceContext struct {
	Config              config.Config
	LangMiddleware      rest.Middleware
	UserTokenMiddleware rest.Middleware
	UserInfoMiddleware  rest.Middleware
	SqlConn             sqlx.SqlConn
	Redis               *redisd.Redisd
	UserRpc             user.User
	Mode                string
	Minio               *minio.Client
	Storage             storaged.InterfaceFactory
	TxCos               *cos.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.MustNewRedis(c.RedisConf)
	redisdConn := redisd.NewRedisd(redisConn, "asset")
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc))
	svc := &ServiceContext{
		Config:              c,
		SqlConn:             sqlx.NewMysql(c.DB.DataSource),
		Redis:               redisdConn,
		Mode:                c.Mode,
		UserRpc:             UserRpc,
		LangMiddleware:      middleware.NewLangMiddleware().Handle,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		UserInfoMiddleware:  middleware.NewUserInfoMiddleware(UserRpc).Handle,
	}
	var err error
	if c.AssetMode == constd.AssetModeLocal {
		svc.Storage, err = storaged.NewProvider(&storaged.ProviderConfig{
			Provider:  storaged.ProviderLocal,
			LocalPath: c.Local.Path,
			Bucket:    c.Local.Bucket,
			Endpoint:  c.Local.PublicBucketAddr,
		})
		if err != nil {
			panic(err)
		}
	} else if c.AssetMode == constd.AssetModeMinio {
		svc.Storage, err = storaged.NewProvider(&storaged.ProviderConfig{
			Provider: storaged.ProviderMinio,
			Endpoint: c.Minio.PublicBucketAddr,
			Key:      c.Minio.AccessKey,
			Secret:   c.Minio.SecretKey,
			Bucket:   c.Minio.Bucket,
		})
	} else if c.AssetMode == constd.AssetModeAliOss {
		svc.Storage, err = storaged.NewProvider(&storaged.ProviderConfig{
			Provider: storaged.ProviderAliOss,
			Endpoint: c.AliOss.PublicBucketAddr,
			Key:      c.AliOss.AccessKeyId,
			Secret:   c.AliOss.AccessKeySecret,
			Bucket:   c.AliOss.Bucket,
		})
	} else if c.AssetMode == constd.AssetModeTxCos {
		svc.Storage, err = storaged.NewProvider(&storaged.ProviderConfig{
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
