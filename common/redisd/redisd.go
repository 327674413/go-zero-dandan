package redisd

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Redisd struct {
	redisConn *redis.Redis
	prefix    string
}

func NewRedisd(redisConn *redis.Redis, prefix string) *Redisd {
	return &Redisd{
		redisConn: redisConn,
		prefix:    prefix,
	}
}
func (t *Redisd) Set(field string, key string, str string) error {
	return t.redisConn.Set(t.prefix+":"+field+":"+key, str)
}
func (t *Redisd) Get(field string, key string) (string, error) {
	return t.redisConn.Get(t.prefix + ":" + field + ":" + key)
}
func (t *Redisd) SetData(field string, key string, data any) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return t.Set(field, key, string(str))
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
func (t *Redisd) Hset(field string, key string, data any) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return t.Hset(field, key, string(str))
}
func (t *Redisd) HgetData(field string, key string, targetStructPointer any) error {
	str, err := t.Hget(field, key)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(str), targetStructPointer)
	return nil
}
func (t *Redisd) Hget(field string, key string) (string, error) {
	return t.redisConn.Hget(t.prefix+":"+field, key)
}
