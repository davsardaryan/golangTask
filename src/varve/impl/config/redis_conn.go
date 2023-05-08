package config

import "github.com/go-redis/redis/v8"

func ConnToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   15,
	})
}
