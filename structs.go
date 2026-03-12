package rediscc

import "github.com/redis/go-redis/v9"

type RedisOptions struct {
	Debug        bool
	DebugPayload bool
}

type RedisDataStore struct {
	Client  *redis.Client
	Options *RedisOptions
}
