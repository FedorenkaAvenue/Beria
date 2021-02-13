package db

import (
	"os"

	"github.com/go-redis/redis"
)

// RedisClient to connect to Redis
var RedisClient = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	Password: os.Getenv("REDIS_PASS"),
	DB:       0,
})
