package connections

import (
	"ginDemo/config"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

func NewRedisConnectionTest() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         "loclhost:6379",
		Password:     "xxxx",
		DB:           0,
		DialTimeout:  time.Second * 3,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 5,
		PoolSize:     5,
	})
	if err := client.Ping().Err(); err != nil {
		log.Panic("ping redis error", "info:", err)
		return nil
	}
	return client
}

func NewRedisConnection(conf *config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.DB,
		DialTimeout:  time.Second * 3,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 5,
		PoolSize:     conf.PoolSize,
	})
	if err := client.Ping().Err(); err != nil {
		log.Panic("ping redis error", "info:", err)
		return nil
	}
	return client
}