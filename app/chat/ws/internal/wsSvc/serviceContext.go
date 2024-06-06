package wsSvc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/chat/ws/internal/wsConfig"
	"go-zero-dandan/app/user/rpc/user"
)

type ServiceContext struct {
	Config  wsConfig.WsConfig
	SqlConn sqlx.SqlConn
	UserRpc user.User
	//Redis   *redisd.Redisd
	Redis    *redis.Client
	Mode     string
	Producer *kq.Pusher
}

func NewServiceContext(c wsConfig.WsConfig) *ServiceContext {
	//redisConn := redis.MustNewRedis(c.RedisConf)
	//redisdConn := redisd.NewRedisd(redisConn, "user")
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc))
	producer := kq.NewPusher([]string{
		"127.0.0.1:19092",
		"127.0.0.1:19093",
		"127.0.0.1:19094",
	}, "test")

	Redis := redis.NewClient(&redis.Options{
		Addr:         c.RedisConf.Host,
		Password:     c.RedisConf.Pass,
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis start fail")
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		//SqlConn:             sqlx.NewMysql(c.DB.DataSource),
		UserRpc: UserRpc,
		//Redis:   redisdConn,
		Redis:    Redis,
		Mode:     c.Mode,
		Producer: producer,
	}
}
