package rediscc

import (
	"github.com/redis/go-redis/v9"
)

type RedisDataStore struct {
	Client *redis.Client
}
