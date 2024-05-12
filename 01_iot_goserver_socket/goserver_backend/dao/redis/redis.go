package redis

import (
	"fmt"

	"github.com/miaogu-go/bluebell/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client
var Rdb *redis.Client

func Init(redisConf *settings.RedisConf) {
	redisOptions := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.Db,
		PoolSize: redisConf.PoolSize,
	}
	rdb = redis.NewClient(redisOptions)
	Rdb = rdb
	fmt.Printf("\nredisOptions info: %v\n", *redisOptions)
	vle, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n ping redis sucessfully: %v\n", vle)
}

func Close() {
	_ = rdb.Close()
}
