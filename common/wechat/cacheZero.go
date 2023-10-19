package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-dandan/common/redisd"
	"go-zero-dandan/common/resd"
	"time"
)

var (
	ErrCacheMiss    = errors.New("cache: miss")
	ErrCASConflict  = errors.New("cache: compare-and-swap conflict")
	ErrNoStats      = errors.New("cache: no statistics available")
	ErrNotStored    = errors.New("cache: item not stored")
	ErrServerError  = errors.New("cache: server error")
	ErrInvalidValue = errors.New("cache: invalid value")
)

type PowerWechatCache struct {
	Redis *redisd.Redisd
}

func NewPowerWechatCache(redisd *redisd.Redisd) *PowerWechatCache {

	return &PowerWechatCache{
		Redis: redisd,
	}
}
func (t *PowerWechatCache) Get(key string, ptrValue interface{}) (returnValue interface{}, err error) {
	b, err := t.Redis.Get("power", key)
	if err == redis.Nil {
		return nil, ErrCacheMiss
	}
	err = json.Unmarshal([]byte(b), &ptrValue)
	returnValue = ptrValue
	return returnValue, resd.Error(err)
}

func (t *PowerWechatCache) Set(key string, value interface{}, expires time.Duration) error {
	mValue, err := json.Marshal(value)
	if err != nil {
		return resd.Error(err)
	}
	if expires > 0 {
		return t.Redis.SetEx("power", key, string(mValue), int(expires))
	} else {
		return t.Redis.Set("power", key, string(mValue))
	}

}

func (t *PowerWechatCache) Has(key string) bool {
	value, err := t.Get(key, nil)
	if value != nil && err == nil {
		return true
	}
	return false
}

type transDecodeStr struct {
	Str string `json:"data"`
}
type transEncode struct {
	Data interface{} `json:"data"`
}

func valueToStr(value interface{}) (string, error) {
	trans := &transEncode{
		Data: value,
	}
	d, err := json.Marshal(trans)
	if err != nil {
		return "", resd.Error(err)
	}
	transRes := &transDecodeStr{}
	err = json.Unmarshal(d, &transRes)
	if err != nil {
		return "", resd.Error(err)
	}
	return transRes.Str, nil
}
func (t *PowerWechatCache) AddNX(key string, value interface{}, ttl time.Duration) bool {
	str, err := valueToStr(value)
	if err != nil {
		logx.Error(err)
		return false
	}
	err = t.Redis.SetNx("power", key, str, int(ttl))

	if err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func (t *PowerWechatCache) Add(key string, value interface{}, ttl time.Duration) (err error) {
	var obj interface{}
	obj, err = t.Get(key, obj)
	if err == ErrCacheMiss {
		return t.SetEx(key, value, ttl)
	} else {
		return resd.NewErr("this value has been actually added to the cache")
	}
}
func (t *PowerWechatCache) SetEx(key string, value interface{}, expires time.Duration) (err error) {
	mValue, err := json.Marshal(value)
	if err != nil {
		return resd.Error(err)
	}
	return t.Redis.SetEx("power", key, string(mValue), int(expires))

}
func (t *PowerWechatCache) Remember(key string, ttl time.Duration, callback func() (interface{}, error)) (obj interface{}, err error) {
	var value interface{}
	value, err = t.Get(key, value)
	if err != nil && err != ErrCacheMiss {
		return nil, resd.Error(err)

	} else if value != nil {
		return value, resd.Error(err)
	}

	value, err = callback()
	if err != nil {
		return nil, resd.Error(err)
	}

	result := t.Put(key, value, ttl)
	if !result {
		err = errors.New(fmt.Sprintf("remember cache put err, ttl:%d", ttl))
	}
	// ErrCacheMiss and query value from source
	return value, resd.Error(err)
}

func (t *PowerWechatCache) Put(key interface{}, value interface{}, ttl time.Duration) bool {

	err := t.SetEx(key.(string), value, ttl)
	if err != nil {
		logx.Error(err)
		return false
	}
	return true

}
