package utils

import (
	"reflect"
	"time"

	red "github.com/gomodule/redigo/redis"
)

// Redis Redis
type Redis struct {
	pool *red.Pool
}

var redis *Redis

// InitRedis 初始化redis
func InitRedis() {
	redis = new(Redis)
	redis.pool = &red.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				RedisConnStr,
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(0),
				red.DialPassword(RedisPassword),
			)
		},
	}
}

// Exec Exec
func (redis Redis) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

// Get 查询数据
func (redis Redis) Get(key string) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	return red.String(con.Do("get", key))
}

// Set 插入
func (redis Redis) Set(key string, value string) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	return red.String(con.Do("set", key, value))
}

// Lpush 列表插入
func (redis Redis) Lpush(key string, values ...interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	args := make([]interface{}, 0)
	args = append(args, key)

	if len(args) > 0 {
		for _, v := range args {
			args = append(args, v)
		}
	}

	return red.String(con.Do("lpush", args...))
}

// Lpop Lpop
func (redis Redis) Lpop(key string) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()

	return red.String(con.Do("lpop", key))
}

// Rpush 列表插入
func (redis Redis) Rpush(key string, values ...interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	args := make([]interface{}, 0)
	args = append(args, key)

	if len(args) > 0 {
		for _, v := range args {
			args = append(args, v)
		}
	}

	return red.String(con.Do("rpush", args...))
}

// Rpop Rpop
func (redis Redis) Rpop(key string) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()

	return red.String(con.Do("rpop", key))
}

// Hset Hset
func (redis Redis) Hset(key string, dataMap map[string]interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	args := make([]interface{}, 0)
	args = append(args, key)

	m := reflect.ValueOf(dataMap)
	// 不是map
	if m.Kind() != reflect.Map {
		UtilsLogger.Error("params error")
		return nil, nil
	}

	// 通过反射拿到所有的key
	keys := m.MapKeys()
	for _, key := range keys {
		value := m.MapIndex(key)
		args = append(args, key.Interface())
		args = append(args, value.Interface())
	}

	return red.String(con.Do("hset", args...))
}

// Hget Hget
func (redis Redis) Hget(key string, keys ...interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	args := make([]interface{}, 0)
	args = append(args, key)

	if len(keys) > 0 {
		for _, v := range keys {
			args = append(args, v)
		}
	}

	return red.String(con.Do("hget", args...))
}
