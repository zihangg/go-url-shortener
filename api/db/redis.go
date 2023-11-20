package db

import "github.com/redis/go-redis/v9"

func CreateRedisClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	return client
}