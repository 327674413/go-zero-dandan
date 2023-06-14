package redisd

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Redisd struct {
	redisConn *redis.Redis
	prefix    string
}
type NotFound struct {
	Msg string
}

func (t *NotFound) Error() string {
	return fmt.Sprintf("redis key: %s not found", t.Msg)
}
func NewRedisd(redisConn *redis.Redis, prefix string) *Redisd {
	return &Redisd{
		redisConn: redisConn,
		prefix:    prefix,
	}
}
func (t *Redisd) Set(field string, key string, str string, expireSec ...int) error {
	if len(expireSec) > 0 {
		return t.redisConn.Setex(t.prefix+":"+field+":"+key, str, expireSec[0])
	} else {
		return t.redisConn.Set(t.prefix+":"+field+":"+key, str)
	}

}
func (t *Redisd) Hset(field string, key string, data string) error {
	return t.redisConn.Hset(t.prefix+":"+field, key, data)
}
func (t *Redisd) Get(field string, key string) (string, error) {
	str, err := t.redisConn.Get(t.prefix + ":" + field + ":" + key)
	if err != nil {
		return "", err
	} else if str == "" {
		return "", &NotFound{Msg: t.prefix + ":" + field + ":" + key}
	} else {
		return str, err
	}
}
func (t *Redisd) Hget(field string, key string) (string, error) {
	return t.redisConn.Hget(t.prefix+":"+field, key)
}
func (t *Redisd) SetData(field string, key string, data any, expireSec ...int) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return t.Set(field, key, string(str), expireSec...)
}
func (t *Redisd) GetData(field string, key string, targetStructPointer any) error {
	str, err := t.Get(field, key)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(str), targetStructPointer)
	return nil
}
func (t *Redisd) HsetData(field string, key string, data any) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return t.redisConn.Hset(t.prefix+":"+field, key, string(str))
}

func (t *Redisd) HgetData(field string, key string, targetStructPointer any) error {
	str, err := t.Hget(field, key)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(str), targetStructPointer)
	return nil
}
