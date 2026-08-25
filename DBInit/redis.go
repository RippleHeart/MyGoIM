package DBInit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"mygoim/Conf"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Conf.Conf.R.Addr,
		Password: Conf.Conf.R.Password,
		DB:       Conf.Conf.R.DB,
		PoolSize: Conf.Conf.R.PoolSize,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis连接失败:", err)
	}
	log.Println("Redis连接成功:", pong)
}
