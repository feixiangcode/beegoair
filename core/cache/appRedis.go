package cache

import (
	"beegoair/core/log"
	red "github.com/gomodule/redigo/redis"
	"strconv"
)

var redis *Redis

func InitRedis(pri string) {
	redis = NewRedis(pri)
}

func ToInt(data interface{}) int {
	ret, err := red.Int(data, nil)
	if err != nil {
		return 0
	}
	return ret
}

func ToInt64(data interface{}) int64 {
	ret, err := red.Int64(data, nil)
	if err != nil {
		return 0
	}
	return ret
}

func ToString(data interface{}) string {
	ret, err := red.String(data, nil)
	if err != nil {
		return ""
	}
	return ret
}

func ToBool(data interface{}) bool {
	ret, err := red.Bool(data, nil)
	if err != nil {
		return false
	}
	return ret
}

type AppRedis struct {

}

func NewAppRedis() *AppRedis {
	appRedis := new(AppRedis)
	return appRedis
}

func (this *AppRedis) Get(key string) interface{} {
	res, err := redis.Exec("GET", key)
	if err != nil {
		log.Error("appredis get fail", err.Error(), key)
	}
	return res
}

func (this *AppRedis) Set(key string, value interface{}, arg ...interface{}) bool {
	var expire int
	if len(arg) > 0 {
		expire = arg[0].(int)
	} else {
		expire = 30
	}

	res, err := redis.Exec("SETEX", key, expire, value)
	if err != nil {
		log.Error("appredis Set fail", err.Error(), key, value, arg)
	}
	result := false
	if res == "OK" {
		result = true
	}
	return result
}

/**
MGET
*/
func (this *AppRedis) MGet(key ...string) map[string]interface{} {
	res, err := redis.Exec("MGET", key)
	if err != nil {
		log.Error("appredis MGet fail", err.Error(), key)
	}
	result := make(map[string]interface{})
	if res == nil {
		return result
	}
	list := res.([]interface{})
	if len(list) > 0 {
		for k, v := range key {
			if list[k] != nil {
				result[v] = list[k]
			}
		}
	}
	return result
}

/**
redis del
删除key
*/
func (this *AppRedis) Del(key string) bool {
	res, err := redis.Exec("DEL", key)
	if err != nil {
		log.Error("appredis DEL fail", err.Error(), key)
	}
	result, _ := red.Int64(res, err)
	if result > 0 {
		return true
	} else {
		return false
	}
}

func (this *AppRedis) Incrby(key string, increment int) int64 {
	res, err := redis.Exec("INCRBY", key, increment)
	if err != nil {
		log.Error("appredis Incrby fail", err.Error(), key, increment)
		return 0
	}
	result, errInt := red.Int64(res, err)
	if errInt != nil {
		log.Error("appredis Incrby toint64 fail", errInt.Error(), key, increment)
		return 0
	}
	return result
}

func (this *AppRedis) Expire(key string, seconds int64) bool {
	res, err := redis.Exec("EXPIRE", key, seconds)
	if err != nil {
		log.Error("redis Expire fail", err.Error(), key, seconds)
		return false
	}
	result, err2 := red.Int(res, err)
	if err2 != nil {
		log.Error("redis Expire to int fail", err.Error(), key, seconds)
		return false
	}
	if result > 0 {
		return true
	}
	return false
}

func (this *AppRedis) HIncrby(key string, field interface{}, inc int) int64 {
	res, err := redis.Exec("HINCRBY", key, field, inc)
	if err != nil {
		log.Error("redis HIncrby fail", err.Error(), key, field, inc)
		return 0
	}
	result, interr := red.Int64(res, err)
	if interr != nil {
		log.Error("redis HIncrby toint64 fail", err.Error(), key, field, inc)
		return 0
	}
	if result > 0 {
		return result
	} else {
		return 0
	}
}

/**
redis hset(key,field, value)
*/
func (this *AppRedis) HSet(key string, field interface{}, value interface{}) bool {
	res, err := redis.Exec("HSET", key, field, value)
	if err != nil {
		log.Error("redis HSet fail", err.Error(), key, field, value)
		return false
	}
	result, _ := red.Int64(res, err)
	if result > 0 {
		return true
	} else {
		return false
	}
}

func (this *AppRedis) HGet(key string, field interface{}) interface{} {
	res, err := redis.Exec("HGET", key, field)
	if err != nil {
		log.Error("redis HGet fail", err.Error(), key, field)
		return nil
	}
	return res
}

func (this *AppRedis) HMGet(key string, fields []interface{}) map[interface{}]interface{} {
	res, err := redis.Exec("HMGET", key, fields...)
	if err != nil {
		log.Error("redis HMGet fail", err.Error(), key, fields)
	}
	result := make(map[interface{}]interface{})
	list := res.([]interface{})
	if len(list) > 0 {
		for k, v := range fields {
			if list[k] != nil {
				result[v] = list[k]
			}
		}
	}
	return result
}

func (this *AppRedis) HGetAll(key string) map[interface{}]interface{} {
	res, err := redis.Exec("HGETALL", key)
	if err != nil {
		log.Error("redis HGETALL fail", err.Error(), key, res)
		return nil
	}
	list := res.([]interface{})
	len := len(list)
	result := make(map[interface{}]interface{})
	for i := 0; i < len; i = i + 2 {
		id := list[i]
		result[id] = list[i+1]
	}
	return result
}

func (this *AppRedis) HMSet(key string, args ...interface{}) bool {
	res, err := redis.Exec("HMSET", key, args...)
	if err != nil {
		log.Error("redis HMSet fail", err.Error(), key, args)
		return false
	}
	result := false
	if res == "OK" {
		result = true
	}
	return result
}

func (this *AppRedis) HLen(key string) int64 {
	res, err := redis.Exec("HLEN", key)
	if err != nil {
		log.Error("redis HLen fail", err.Error(), key)
		return 0
	}
	result, err := red.Int64(res, err)
	if err != nil {
		log.Error("redis HLen toint64 fail", err.Error(), key)
		return 0
	}
	return result
}

func (this *AppRedis) HDel(key string, field interface{}) bool {
	res, err := redis.Exec("HDEL", key, field)
	if err != nil {
		log.Error("redis HDel fail", err.Error(), key, field)
	}
	result, err := red.Int64(res, err)
	if err != nil {
		log.Error("redis HDel toint64 fail", err.Error(), key, field)
		return false
	}
	if result > 0 {
		return true
	} else {
		return false
	}
}

func (this *AppRedis) RPush(key string, args ...interface{}) bool {
	res, err := redis.Exec("RPUSH", key, args...)
	if err != nil {
		log.Error("redis RPush fail", err.Error(), key, args)
		return false
	}
	result, err := red.Int64(res, err)
	if err != nil {
		log.Error("redis RPush toint64 fail", err.Error(), key, args)
		return false
	}
	if result > 0 {
		return true
	} else {
		return false
	}
}

func (this *AppRedis) LPush(key string, args ...interface{}) bool {
	res, err := redis.Exec("LPUSH", key, args...)
	if err != nil {
		log.Error("redis LPUSH fail", err.Error(), key, args)
		return false
	}
	result, err := red.Int64(res, err)
	if err != nil {
		log.Error("redis LPUSH toint64 fail", err.Error(), key, args)
		return false
	}
	if result > 0 {
		return true
	} else {
		return false
	}
}

func (this *AppRedis) LLen(key string) int64 {
	res, err := redis.Exec("LLEN", key)
	if err != nil {
		log.Error("redis LLen fail", err.Error(), key)
		return 0
	}
	result, err := red.Int64(res, err)
	if err != nil {
		log.Error("redis LLen toint64 fail", err.Error(), key)
		return 0
	}
	return result
}

func (this *AppRedis) LRange(key string, start int64, stop int64) []interface{} {
	res, err := redis.Exec("LRANGE", key, start, stop)
	result := make([]interface{}, 0)
	if err != nil {
		log.Error("redis LRANGE fail", err.Error(), key, start, stop)
		return result
	}
	result = res.([]interface{})
	return result
}

func (this *AppRedis) ZAdd(key string, args ...interface{}) int64 {
	res, err := redis.Exec("ZADD", key, args...)
	if err != nil {
		log.Error("redis zadd fail", err.Error(), key, args)
		return 0
	}
	result, err := red.Int64(res, err)
	return result
}

func (this *AppRedis) ZCount(key string) int64 {
	res, err := redis.Exec("ZCOUNT", key, "-inf", "+inf")
	if err != nil {
		log.Error("redis zcount fail", err.Error(), key)
		return 0
	}
	result, err := red.Int64(res, err)
	if err != nil {
		log.Error("redis zcount toint64 fail", err.Error(), key)
		return 0
	}
	return result
}

func (this *AppRedis) ZRem(key string, args ...interface{}) int64 {
	res, err := redis.Exec("ZREM", key, args...)
	if err != nil {
		log.Error("redis ZRem fail", err.Error(), key, args)
		return 0
	}
	result, err := red.Int64(res, err)
	if err != nil {
		log.Error("redis ZRem toint64 fail", err.Error(), key, args)
		return 0
	}
	return result
}

func (this *AppRedis) ZRevRangeDesc(key string, score int64, limit int) []interface{} {
	start := "+inf"
	if score > 0 {
		start = "(" + strconv.FormatInt(score, 10)
	}
	res, err := redis.Exec("ZREVRANGEBYSCORE", key, start, "-inf", "limit", 0, limit)
	result := make([]interface{}, 0)
	if err != nil {
		log.Error("redis ZREVRANGEBYSCORE fail", key, err.Error(), start, limit)
		return result
	}
	result = res.([]interface{})
	return result
}

func (this *AppRedis) ZRevRangeDescWithScore(key string, score int64, limit int) ([]interface{}, []int64) {
	start := "+inf"
	if score > 0 {
		start = "(" + strconv.FormatInt(score, 10)
	}
	res, err := redis.Exec("ZREVRANGEBYSCORE", key, start, "-inf", "withscores", "limit", 0, limit)
	result := make([]interface{}, 0)
	scores := make([]int64, 0)
	if err != nil {
		log.Error("redis ZREVRANGEBYSCORE fail", key, err.Error(), start, limit)
		return result, scores
	}
	resList := res.([]interface{})
	for k, v := range resList {
		if k%2 == 0 {
			result = append(result, v)
		} else {
			scores = append(scores, ToInt64(v))
		}
	}
	return result, scores
}

func (this *AppRedis) ZRangeAsc(key string, score int64, limit int) []interface{} {
	start := "-inf"
	if score > 0 {
		start = "(" + strconv.FormatInt(score, 10)
	}
	res, err := redis.Exec("ZRANGEBYSCORE", key, start, "+inf", "limit", 0, limit)
	result := make([]interface{}, 0)
	if err != nil {
		log.Error("redis ZRangeAsc fail", key, err.Error(), start, limit)
		return result
	}
	result = res.([]interface{})
	return result
}

func (this *AppRedis) ZRangeByScore(key string, min int64, max int64, limit int) []interface{} {
	var res interface{}
	var err error
	if limit > 0 {
		res, err = redis.Exec("ZRANGEBYSCORE", key, min, max, "limit", 0, limit)
	} else {
		res, err = redis.Exec("ZRANGEBYSCORE", key, min, max)
	}
	result := make([]interface{}, 0)
	if err != nil {
		log.Error("redis ZRangeByScore fail", key, err.Error(), min, max, limit)
		return result
	}
	result = res.([]interface{})
	return result
}
func (this *AppRedis) ZRangeAscWithScore(key string, score int64, limit int) ([]interface{}, []int64) {
	start := "-inf"
	if score > 0 {
		start = "(" + strconv.FormatInt(score, 10)
	}
	res, err := redis.Exec("ZRANGEBYSCORE", key, start, "+inf", "withscores", "limit", 0, limit)
	result := make([]interface{}, 0)
	scores := make([]int64, 0)
	if err != nil {
		log.Error("redis ZREVRANGEBYSCORE fail", key, err.Error(), start, limit)
		return result, scores
	}
	resList := res.([]interface{})
	for k, v := range resList {
		if k%2 == 0 {
			result = append(result, v)
		} else {
			scores = append(scores, ToInt64(v))
		}
	}
	return result, scores
}
