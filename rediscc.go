package rediscc

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func Connect(redisUri string, dbNumber string) (*RedisDataStore, error) {
	redisUrl := fmt.Sprintf("%s/%s", redisUri, dbNumber)
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	fmt.Printf("You successfully connected to Redis: %s\n", redisUrl)

	redisDataStore := &RedisDataStore{
		client: client,
		Debug:  true,
	}

	return redisDataStore, nil
}

type RedisDataStore struct {
	client *redis.Client
	Debug  bool
}

func (redisDataStore *RedisDataStore) Publish(ctx context.Context, channel string, payload interface{}) error {
	if redisDataStore.Debug {
		fmt.Println("Publish", channel, payload)
	}
	return redisDataStore.client.Publish(ctx, channel, payload).Err()
}

func (redisDataStore *RedisDataStore) Get(ctx context.Context, key string) *redis.StringCmd {
	if redisDataStore.Debug {
		fmt.Println("Get", key)
	}
	return redisDataStore.client.Get(ctx, key)
}

func (redisDataStore *RedisDataStore) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if redisDataStore.Debug {
		fmt.Println("Set", key, value, expiration)
	}
	return redisDataStore.client.Set(ctx, key, value, expiration).Err()
}

func (redisDataStore *RedisDataStore) SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if redisDataStore.Debug {
		fmt.Println("SetWithExpiration", key, value, expiration)
	}
	return redisDataStore.client.Set(ctx, key, value, expiration).Err()
}
