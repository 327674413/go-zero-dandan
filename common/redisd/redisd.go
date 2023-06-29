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

// Del 删除
func (t *Redisd) Del(field string, keys ...string) (int, error) {
	for i, v := range keys {
		keys[i] = t.prefix + ":" + field + ":" + v
	}
	return t.redisConn.Del(keys...)
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
func (t *Redisd) Hset(field string, key string, data string, expireSec ...int) error {
	err := t.redisConn.Hset(t.prefix+":"+field, key, data)
	if err != nil {
		return err
	}
	if len(expireSec) > 0 {
		return t.SetExSec(field, key, expireSec[0])
	}
	return nil

}

// Get 获取值
func (t *Redisd) Get(field string, key string) (string, error) {
	str, err := t.redisConn.Get(t.prefix + ":" + field + ":" + key)
	if err != nil && err != redis.Nil {
		//报错返回错误信息
		return "", err
	} else if str == "" {
		//没找到数据，按空返回
		return "", nil //&NotFound{Msg: t.prefix + ":" + field + ":" + key}
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
	if str == "" {
		return &NotFound{Msg: t.prefix + ":" + field + ":" + key}
	}
	json.Unmarshal([]byte(str), targetStructPointer)
	return nil
}

/*

//目前感觉这个方法很奇怪，暂时先不考虑
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

*/

// SetExSec 设置新的过期时间,传入多少秒后过期
func (t *Redisd) SetExSec(field string, key string, expireSec int) error {
	res, err := t.redisConn.SetnxEx(t.prefix+":"+field, key, expireSec)
	//目前理解res的布尔可以直接判断，不需要看err
	if res == true {
		return nil
	} else {
		return err
	}
}
