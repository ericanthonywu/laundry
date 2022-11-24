package Lib

import (
	"fmt"
	"github.com/go-redis/redis"
	"laundry/Config"
	"time"
)

var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(Config.RedisClientConfig())
}

func RDBSet(key string, value string, expiration time.Duration) {
	if err := rdb.Set(key, value, expiration).Err(); err != nil {
		panic(err)
	}
}

func RDBGet(key string) string {
	value, err := rdb.Get(key).Result()

	if err == redis.Nil {
		fmt.Println(key + ": does not exist")
		return ""
	} else if err != nil {
		panic(err)
	}

	return value
}

func RDBDel(key string) {
	if err := rdb.Del(key).Err(); err != nil {
		panic(err)
	}
}
