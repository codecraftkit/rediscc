package rediscc

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepositoryPorts interface {
	Publish(ctx context.Context, channel string, payload interface{}) error
	Get(ctx context.Context, key string) (string, error)
	GetRaw(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, channel string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
}
