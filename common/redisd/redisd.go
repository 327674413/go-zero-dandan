package redisd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	redisx "github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-dandan/common/resd"
	"strconv"
)

type Redisd struct {
	redisConn *redisx.Redis
	prefix    string
}

// 如果自定义未找到数据的错误，和err判断的时候还要断言，会比较麻烦，现在是按照没找到数据也不报错的方式走

func NewRedisd(redisConn *redisx.Redis, prefix string) *Redisd {
	return &Redisd{
		redisConn: redisConn,
		prefix:    prefix,
	}
}

// Set 设置值
func (t *Redisd) Set(field string, key string, str string) (danErr error) {
	err := t.redisConn.Set(t.FieldKey(field, key), str)
	if err != nil {
		return resd.Error(err)
	}
	return nil
}

// SetCtx 设置值,带上下文
func (t *Redisd) SetCtx(ctx context.Context, field string, key string, str string) (danErr error) {
	err := t.redisConn.SetCtx(ctx, t.FieldKey(field, key), str)
	if err != nil {
		return resd.ErrorCtx(ctx, err)
	}
	return nil
}

// SetNx 不存在则设置值
func (t *Redisd) SetNx(field string, key string, str string, expireSec ...int) (danErr error) {
	var res bool
	var err error
	if len(expireSec) > 0 {
		res, err = t.redisConn.SetnxEx(t.FieldKey(field, key), str, expireSec[0])
	} else {
		res, err = t.redisConn.Setnx(t.FieldKey(field, key), str)
	}
	if err != nil {
		return resd.Error(err, resd.ErrRedisSet)
	} else if !res {
		return resd.NewErr("setnx fail", resd.ErrRedisSet)
	}
	return resd.Error(err, resd.ErrRedisSet)
}

// SetNxCtx 不存在则设置值,带上下文
func (t *Redisd) SetNxCtx(ctx context.Context, field string, key string, str string, expireSec ...int) (danErr error) {
	var res bool
	var err error
	if len(expireSec) > 0 {
		res, err = t.redisConn.SetnxExCtx(ctx, t.FieldKey(field, key), str, expireSec[0])
	} else {
		res, err = t.redisConn.SetnxCtx(ctx, t.FieldKey(field, key), str)
	}
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrRedisSet)
	} else if !res {
		return resd.NewErrCtx(ctx, "setnx fail", resd.ErrRedisSet)
	}
	return resd.ErrorCtx(ctx, err, resd.ErrRedisSet)
}

// Del 删除
func (t *Redisd) Del(field string, keys ...string) (effectNum int, danErr error) {
	for i, v := range keys {
		keys[i] = t.prefix + ":" + field + ":" + v
	}
	effectNum, err := t.redisConn.Del(keys...)
	if err != nil {
		return effectNum, resd.Error(err)
	} else {
		return effectNum, nil
	}
}

// DelCtx 删除，带上下文
func (t *Redisd) DelCtx(ctx context.Context, field string, keys ...string) (effectNum int, danErr error) {
	for i, v := range keys {
		keys[i] = t.prefix + ":" + field + ":" + v
	}
	effectNum, err := t.redisConn.DelCtx(ctx, keys...)
	if err != nil {
		return effectNum, resd.Error(err)
	} else {
		return effectNum, nil
	}
}

// Inc 值递增n
func (t *Redisd) Inc(field string, key string, num int, expireSec ...int) (danErr error) {
	oldNumStr, err := t.Get(field, key)
	if err == nil {
		return resd.Error(err)
	}
	oldNum, err := strconv.Atoi(oldNumStr)
	if err != nil {
		oldNum = 0
	}
	oldNum = oldNum + num
	str := fmt.Sprintf("%d", oldNum)
	if len(expireSec) > 0 {
		return t.redisConn.Setex(t.FieldKey(field, key), str, expireSec[0])
	} else {
		return t.redisConn.Set(t.FieldKey(field, key), str)
	}

}

// IncCtx 值递增n，带上下文
func (t *Redisd) IncCtx(ctx context.Context, field string, key string, num int, expireSec ...int) (danErr error) {
	oldNumStr, err := t.GetCtx(ctx, field, key)
	if err == nil {
		return resd.ErrorCtx(ctx, err, resd.ErrRedisGet)
	}
	oldNum, err := strconv.Atoi(oldNumStr)
	if err != nil {
		oldNum = 0
	}
	oldNum = oldNum + num
	str := fmt.Sprintf("%d", oldNum)
	if len(expireSec) > 0 {
		return t.redisConn.SetexCtx(ctx, t.FieldKey(field, key), str, expireSec[0])
	} else {
		return t.redisConn.SetCtx(ctx, t.FieldKey(field, key), str)
	}

}

// Dec 值递减n，会变成负数
func (t *Redisd) Dec(field string, key string, num int, expireSec ...int) (danErr error) {
	oldNumStr, err := t.Get(field, key)
	if err == nil {
		return resd.Error(err, resd.ErrRedisGet)
	}
	oldNum, err := strconv.Atoi(oldNumStr)
	if err != nil {
		oldNum = 0
	}
	oldNum = oldNum - num
	str := fmt.Sprintf("%d", oldNum)
	if len(expireSec) > 0 {
		return t.redisConn.Setex(t.FieldKey(field, key), str, expireSec[0])
	} else {
		return t.redisConn.Set(t.FieldKey(field, key), str)
	}

}

// DecCtx 值递减n，会变成负数，带上下文
func (t *Redisd) DecCtx(ctx context.Context, field string, key string, num int, expireSec ...int) (danErr error) {
	oldNumStr, err := t.GetCtx(ctx, field, key)
	if err == nil {
		return resd.ErrorCtx(ctx, err, resd.ErrRedisGet)
	}
	oldNum, err := strconv.Atoi(oldNumStr)
	if err != nil {
		oldNum = 0
	}
	oldNum = oldNum - num
	str := fmt.Sprintf("%d", oldNum)
	if len(expireSec) > 0 {
		return t.redisConn.SetexCtx(ctx, t.FieldKey(field, key), str, expireSec[0])
	} else {
		return t.redisConn.SetCtx(ctx, t.FieldKey(field, key), str)
	}

}

// Hset 设置哈希值
func (t *Redisd) Hset(field string, key string, data string) (danErr error) {
	err := t.redisConn.Hset(t.prefix+":"+field, key, data)
	if err != nil {
		return resd.Error(err, resd.ErrRedisSet)
	}
	return nil

}

// HsetCtx 设置哈希值，带上下文
func (t *Redisd) HsetCtx(ctx context.Context, field string, key string, data string) (danErr error) {
	err := t.redisConn.HsetCtx(ctx, t.prefix+":"+field, key, data)
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrRedisSet)
	}
	return nil

}

// Get 获取值, 单个时key用id，多个时key可以用list、info之类的字符串标识
func (t *Redisd) Get(field string, key string) (str string, danErr error) {
	str, err := t.redisConn.Get(t.prefix + ":" + field + ":" + key)
	if err != nil && err != redis.Nil {
		//报错返回错误信息
		return "", resd.Error(err, resd.ErrRedisGet)
	} else if str == "" {
		//没找到数据，按空返回，不报错，暂时没场景需要区分是否为你nil，如果要区分到时用原声client处理
		return "", nil
	} else {
		return str, nil
	}
}

// GetCtx 获取值，带上下文, 单个时key用id，多个时key可以用list、info之类的字符串标识
func (t *Redisd) GetCtx(ctx context.Context, field string, key string) (str string, danErr error) {
	k := t.FieldKey(field, key)
	str, err := t.redisConn.GetCtx(ctx, k)
	if err != nil && err != redis.Nil {
		//报错返回错误信息
		return "", resd.ErrorCtx(ctx, err, resd.ErrRedisGet)
	} else if str == "" {
		//没找到数据，按空返回
		return "", resd.ErrorCtx(ctx, err, resd.ErrRedisKeyNil) //&NotFound{Msg: t.prefix + ":" + field + ":" + key}
	} else {
		return str, nil
	}
}

// Hget 获取哈希值
func (t *Redisd) Hget(field string, key string) (str string, danErr error) {
	str, err := t.redisConn.Hget(t.prefix+":"+field, key)
	if err != nil {
		return "", resd.Error(err, resd.ErrRedisGet)
	}
	return str, nil
}

// HgetCtx 获取哈希值，带上下文
func (t *Redisd) HgetCtx(ctx context.Context, field string, key string) (str string, danErr error) {
	str, err := t.redisConn.HgetCtx(ctx, t.prefix+":"+field, key)
	if err != nil {
		return "", resd.Error(err, resd.ErrRedisGet)
	}
	return str, nil
}

// Hgetall 获取哈希全部
func (t *Redisd) Hgetall(field string, key string) (data map[string]string, danErr error) {
	data, err := t.redisConn.Hgetall(t.prefix + ":" + field + ":" + key)
	if err != nil {
		return nil, resd.Error(err, resd.ErrRedisGet)
	}
	return data, nil
}

// HgetallCtx 获取哈希值全部，带上下文
func (t *Redisd) HgetallCtx(ctx context.Context, field string) (data map[string]string, danErr error) {
	data, err := t.redisConn.HgetallCtx(ctx, t.prefix+":"+field)
	if err != nil {
		return nil, resd.Error(err, resd.ErrRedisGet)
	}
	return data, nil
}

// Zadd zadd方法
func (t *Redisd) Zadd(field string, key string, score int64, value string) (res bool, danErr error) {
	res, err := t.redisConn.Zadd(t.FieldKey(field, key), score, value)
	if err != nil {
		return false, resd.Error(err, resd.ErrRedisSet)
	}
	return res, nil
}

// ZaddCtx zadd方法,带ctx
func (t *Redisd) ZaddCtx(ctx context.Context, field string, key string, score int64, value string) (res bool, danErr error) {
	res, err := t.redisConn.ZaddCtx(ctx, t.FieldKey(field, key), score, value)
	if err != nil {
		return false, resd.Error(err, resd.ErrRedisSet)
	}
	return res, nil
}

// ZpageCtx 根据zadd数据进行score分页，带ctx
func (t *Redisd) ZpageCtx(ctx context.Context, field string, key string, page int, size int, isDesc bool) (list []redisx.Pair, danErr error) {
	if isDesc {
		list, danErr = t.redisConn.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, t.FieldKey(field, key), 0, 9999999999, page-1, size)
	} else {
		list, danErr = t.redisConn.ZrangebyscoreWithScoresAndLimitCtx(ctx, t.FieldKey(field, key), 0, 9999999999, page-1, size)
	}
	if danErr != nil {
		return list, resd.Error(danErr, resd.ErrRedisGet)
	}
	return list, nil
}

// Zpage 根据zadd数据进行score分页
func (t *Redisd) Zpage(field string, key string, page int, size int, isDesc bool) (list []redisx.Pair, danErr error) {
	if isDesc {
		list, danErr = t.redisConn.ZrevrangebyscoreWithScoresAndLimit(t.FieldKey(field, key), 0, 9999999999, page-1, size)
	} else {
		list, danErr = t.redisConn.ZrangebyscoreWithScoresAndLimit(t.FieldKey(field, key), 0, 9999999999, page-1, size)
	}
	if danErr != nil {
		return list, resd.Error(danErr, resd.ErrRedisGet)
	}
	return list, nil
}

// CursorPageDesc 游标降序分页
func (t *Redisd) CursorPageDesc(field string, key string, page int, size int) (list []redisx.Pair, danErr error) {
	list, danErr = t.redisConn.ZrevrangebyscoreWithScoresAndLimit(t.FieldKey(field, key), 0, 9999999999, page-1, size)
	if danErr != nil {
		return list, resd.Error(danErr, resd.ErrRedisGet)
	}
	return list, nil
}

// CursorPageAsc 游标升序分页
func (t *Redisd) CursorPageAsc(field string, key string, page int, size int) (list []redisx.Pair, danErr error) {
	list, danErr = t.redisConn.ZrangebyscoreWithScoresAndLimit(t.FieldKey(field, key), 0, 9999999999, page-1, size)
	if danErr != nil {
		return list, resd.Error(danErr, resd.ErrRedisGet)
	}
	return list, nil
}

// CursorPageAscCtx 游标升序分页，带ctx
func (t *Redisd) CursorPageAscCtx(ctx context.Context, field string, key string, cursor int64, size int) (list []redisx.Pair, danErr error) {
	list, danErr = t.redisConn.ZrangebyscoreWithScoresAndLimitCtx(ctx, t.FieldKey(field, key), cursor, 0, 0, size)
	if danErr != nil {
		return list, resd.Error(danErr, resd.ErrRedisGet)
	}
	return list, nil
}

// SetData 将数据转成json设置
func (t *Redisd) SetData(field string, key string, data any) (danErr error) {
	str, err := json.Marshal(data)
	if err != nil {
		return resd.Error(err, resd.ErrJsonEncode)
	}
	return t.Set(field, key, string(str))
}

// SetDataCtx 将数据转成json设置,带上下文
func (t *Redisd) SetDataCtx(ctx context.Context, field string, key string, data any) (danErr error) {
	str, err := json.Marshal(data)
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrJsonEncode)
	}
	return t.SetCtx(ctx, field, key, string(str))
}

// SetDataEx 将数据转成json设置,并设置过期时间
func (t *Redisd) SetDataEx(field string, key string, data any, expireSec int) (danErr error) {
	str, err := json.Marshal(data)
	if err != nil {
		return resd.Error(err, resd.ErrJsonEncode)
	}
	return t.SetEx(field, key, string(str), expireSec)
}

// SetDataExCtx 将数据转成json设置,并设置过期时间,带上下文
func (t *Redisd) SetDataExCtx(ctx context.Context, field string, key string, data any, expireSec int) (danErr error) {
	str, err := json.Marshal(data)
	if err != nil {
		return resd.ErrorCtx(ctx, err, resd.ErrJsonEncode)
	}
	return t.SetExCtx(ctx, field, key, string(str), expireSec)
}

// GetData 获取数据并且转json
func (t *Redisd) GetData(field string, key string, targetStructPointer any) (isSucc bool, danErr error) {
	str, err := t.Get(field, key)
	if err != nil {
		return false, resd.Error(err)
	}
	if str == "" {
		return false, nil
	}
	err = json.Unmarshal([]byte(str), targetStructPointer)
	if err != nil {
		return false, resd.Error(err, resd.ErrJsonDecode)
	} else {
		return true, nil
	}
}

// GetDataCtx 获取数据并且转json,带上下文
func (t *Redisd) GetDataCtx(ctx context.Context, field string, key string, targetStructPointer any) (isSucc bool, danErr error) {
	str, err := t.GetCtx(ctx, field, key)
	if err != nil {
		return false, resd.ErrorCtx(ctx, err)
	}
	if str == "" {
		return false, nil
	}
	err = json.Unmarshal([]byte(str), targetStructPointer)
	if err != nil {
		return false, resd.ErrorCtx(ctx, err, resd.ErrJsonDecode)
	} else {
		return true, nil
	}
}

// Exists 校验key是否存在
func (t *Redisd) Exists(field string) (isExist bool, danErr error) {
	isExist, danErr = t.redisConn.Exists(t.prefix + ":" + field)
	if danErr != nil {
		return false, resd.Error(danErr, resd.ErrRedisGet)
	}
	return
}

// ExistsCtx 校验key是否存在，带上下文
func (t *Redisd) ExistsCtx(ctx context.Context, field string) (isExist bool, danErr error) {
	isExist, danErr = t.redisConn.ExistsCtx(ctx, t.prefix+":"+field)
	if danErr != nil {
		return false, resd.Error(danErr, resd.ErrRedisGet)
	}
	return
}

// Hexists 校验哈希中的key是否存在
func (t *Redisd) Hexists(field string, key string) (isExist bool, danErr error) {
	isExist, danErr = t.redisConn.Hexists(t.prefix+":"+field, key)
	if danErr != nil {
		return false, resd.Error(danErr, resd.ErrRedisGet)
	}
	return
}

// HexistsCtx 校验哈希中的key是否存在，带上下文
func (t *Redisd) HexistsCtx(ctx context.Context, field string, key string) (isExist bool, danErr error) {
	isExist, danErr = t.redisConn.HexistsCtx(ctx, t.prefix+":"+field, key)
	if danErr != nil {
		return false, resd.Error(danErr, resd.ErrRedisGet)
	}
	return
}

// SetEx 设置带过期时间的值
func (t *Redisd) SetEx(field string, key string, value string, expireSec int) (danErr error) {
	err := t.redisConn.Setex(t.FieldKey(field, key), value, expireSec)
	if err != nil {
		return resd.Error(err, resd.ErrRedisSet)
	}
	return
}

// SetExCtx 设置带过期时间的值，带上下文
func (t *Redisd) SetExCtx(ctx context.Context, field string, key string, value string, expireSec int) (danErr error) {
	err := t.redisConn.SetexCtx(ctx, t.FieldKey(field, key), value, expireSec)
	if err != nil {
		return resd.Error(err, resd.ErrRedisSet)
	}
	return
}

// Expire 给缓存续期
func (t *Redisd) Expire(field string, key string, expireSec int) (danErr error) {
	err := t.redisConn.Expire(t.FieldKey(field, key), expireSec)
	if err != nil {
		return resd.Error(err, resd.ErrRedisSet)
	}
	return
}

// ExpireCtx 给缓存续期，带上下文
func (t *Redisd) ExpireCtx(ctx context.Context, field string, key string, expireSec int) (danErr error) {
	err := t.redisConn.ExpireCtx(ctx, t.FieldKey(field, key), expireSec)
	if err != nil {
		return resd.Error(err, resd.ErrRedisSet)
	}
	return
}

// DelKeyByPrefix 使用eval方式删除执行前缀的key
func (t *Redisd) DelKeyByPrefix(keyPrefix string) (res any, danErr error) {
	script := `
		local keys = redis.call("KEYS", ARGV[1])
		for i = 1, #keys do
			redis.call("DEL", keys[i])
		end
		return #keys
	`
	args := []interface{}{t.prefix + ":" + keyPrefix}
	res, err := t.redisConn.Eval(script, []string{}, args...)
	if err != nil {
		return nil, resd.Error(err, resd.ErrRedis)
	}
	return res, err
}

// DelKeyByPrefixCtx 使用eval方式删除执行前缀的key，带上下文
func (t *Redisd) DelKeyByPrefixCtx(ctx context.Context, keyPrefix string) (res any, danErr error) {
	script := `
		local keys = redis.call("KEYS", ARGV[1])
		for i = 1, #keys do
			redis.call("DEL", keys[i])
		end
		return #keys
	`
	args := []interface{}{t.prefix + ":" + keyPrefix}
	res, err := t.redisConn.EvalCtx(ctx, script, []string{}, args...)
	if err != nil {
		return nil, resd.Error(err, resd.ErrRedis)
	}
	return res, err
}
func (t *Redisd) Client() *redisx.Redis {
	return t.redisConn
}

func (t *Redisd) FieldKey(field string, key string) string {
	k := t.prefix
	if field != "" {
		k = k + ":" + field
	}
	if key != "" {
		k = k + ":" + key
	}
	return k

}
