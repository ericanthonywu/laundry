package Lib

import (
	"github.com/go-redis/redis"
	"laundry/Config"
	"time"
)

var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(Config.RedisClientConfig())
}

func RDBSet(key string, value interface{}, expiration time.Duration) {
	if err := rdb.Set(key, value, expiration).Err(); err != nil {
		panic(err)
	}
}

func RDBGet(key string) (string, bool) {
	value, err := rdb.Get(key).Result()

	if err != nil && err != redis.Nil {
		panic(err)
	}

	return value, err == redis.Nil
}

func RDBGetUint(key string) (uint64, bool) {
	value, err := rdb.Get(key).Uint64()

	if err != nil && err != redis.Nil {
		panic(err)
	}

	return value, err == redis.Nil
}

func RDBDel(key string) {
	if err := rdb.Del(key).Err(); err != nil {
		panic(err)
	}
}
