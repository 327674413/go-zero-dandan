package svc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-dandan/app/im/bak/internal/config"
	"go-zero-dandan/app/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn
	UserRpc user.User
	//Redis   *redisd.Redisd
	Redis *redis.Client
	Mode  string
}

func NewServiceContext(c config.Config) *ServiceContext {
	//redisConn := redis.MustNewRedis(c.RedisConf)
	//redisdConn := redisd.NewRedisd(redisConn, "user")
	UserRpc := user.NewUser(zrpc.MustNewClient(c.UserRpc))

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
		Redis: Redis,
		Mode:  c.Mode,
	}
}
