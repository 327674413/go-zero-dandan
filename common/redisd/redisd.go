package redisd

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
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

// Set 设置值
func (t *Redisd) Set(field string, key string, str string, expireSec ...int) error {
	if len(expireSec) > 0 {
		return t.redisConn.Setex(t.prefix+":"+field+":"+key, str, expireSec[0])
	} else {
		return t.redisConn.Set(t.prefix+":"+field+":"+key, str)
	}

}

// Inc 值递增n
func (t *Redisd) Inc(field string, key string, num int, expireSec ...int) error {
	oldNumStr, err := t.Get(field, key)
	if err == nil {
		return err
	}
	oldNum, err := strconv.Atoi(oldNumStr)
	if err != nil {
		oldNum = 0
	}
	oldNum = oldNum + num
	str := fmt.Sprintf("%d", oldNum)
	if len(expireSec) > 0 {
		return t.redisConn.Setex(t.prefix+":"+field+":"+key, str, expireSec[0])
	} else {
		return t.redisConn.Set(t.prefix+":"+field+":"+key, str)
	}

}

// Dec 值递减n，会变成负数
func (t *Redisd) Dec(field string, key string, num int, expireSec ...int) error {
	oldNumStr, err := t.Get(field, key)
	if err == nil {
		return err
	}
	oldNum, err := strconv.Atoi(oldNumStr)
	if err != nil {
		oldNum = 0
	}
	oldNum = oldNum - num
	str := fmt.Sprintf("%d", oldNum)
	if len(expireSec) > 0 {
		return t.redisConn.Setex(t.prefix+":"+field+":"+key, str, expireSec[0])
	} else {
		return t.redisConn.Set(t.prefix+":"+field+":"+key, str)
	}

}

// Hset 设置哈希值
func (t *Redisd) Hset(field string, key string, data string) error {
	return t.redisConn.Hset(t.prefix+":"+field, key, data)
}

// Get 获取值
func (t *Redisd) Get(field string, key string) (string, error) {
	str, err := t.redisConn.Get(t.prefix + ":" + field + ":" + key)
	if err != nil {
		//报错返回错误信息
		return "", err
	} else if str == "" {
		//没找到数据，返回特殊错误
		return "", &NotFound{Msg: t.prefix + ":" + field + ":" + key}
	} else {
		return str, err
	}
}

// Hget 获取哈希值
func (t *Redisd) Hget(field string, key string) (string, error) {
	return t.redisConn.Hget(t.prefix+":"+field, key)
}

// SetData 将数据转成json设置
func (t *Redisd) SetData(field string, key string, data any, expireSec ...int) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return t.Set(field, key, string(str), expireSec...)
}

// GetData 获取数据并且转json
func (t *Redisd) GetData(field string, key string, targetStructPointer any) error {
	str, err := t.Get(field, key)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(str), targetStructPointer)
	return nil
}

// HsetData 转成json设置哈希值
func (t *Redisd) HsetData(field string, key string, data any) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return t.redisConn.Hset(t.prefix+":"+field, key, string(str))
}

// HgetData 获取哈希值后转json
func (t *Redisd) HgetData(field string, key string, targetStructPointer any) error {
	str, err := t.Hget(field, key)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(str), targetStructPointer)
	return nil
}
