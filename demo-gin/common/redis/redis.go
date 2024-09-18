package redis

import (
	"fmt"
	"go-gin/config"

	"github.com/go-redis/redis/v8"
)

func Client() *redis.Client {
	return con
}

var con *redis.Client

func InitRedis() {
	ds := config.GetRedisConfig()
	con = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ds.Host, ds.Port),
		Password: ds.Password,
		PoolSize: 10,
		//
	})
}
