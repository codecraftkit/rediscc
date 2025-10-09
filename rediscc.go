package rediscc

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func Connect(redisUri string, dbNumber string, options *RedisOptions) (*RedisDataStore, error) {
	redisUrl := fmt.Sprintf("%s/%s", redisUri, dbNumber)
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	fmt.Printf("You successfully connected to Redis: %s\n", redisUrl)

	if options == nil {
		options = &RedisOptions{}
	}

	redisDataStore := &RedisDataStore{
		client:  client,
		options: options,
	}

	return redisDataStore, nil
}

func (redisDataStore *RedisDataStore) Publish(ctx context.Context, channel string, payload interface{}) error {
	if redisDataStore.options.Debug {
		fmt.Println("[LOG] Publish", channel)
	}
	if redisDataStore.options.DebugPayload {
		fmt.Println("[LOG] Payload", payload)
	}
	return redisDataStore.client.Publish(ctx, channel, payload).Err()
}

func (redisDataStore *RedisDataStore) Get(ctx context.Context, key string) (string, error) {
	if redisDataStore.options.Debug {
		fmt.Println("[LOG] Get", key)
	}
	value, err := redisDataStore.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (redisDataStore *RedisDataStore) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if redisDataStore.options.Debug {
		fmt.Println("[Redis Client]", redisDataStore.client)
		fmt.Println("[LOG] Set", key, value, expiration)
	}
	return redisDataStore.client.Set(ctx, key, value, expiration).Err()
}

func (redisDataStore *RedisDataStore) SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if redisDataStore.options.Debug {
		fmt.Println("[LOG] SetWithExpiration", key, value, expiration)
	}
	return redisDataStore.client.Set(ctx, key, value, expiration).Err()
}

func (redisDataStore *RedisDataStore) Del(ctx context.Context, channel string) error {
	if redisDataStore.options.Debug {
		fmt.Println("[LOG] Del", channel)
	}
	_, err := redisDataStore.client.Del(ctx, channel).Result()
	if err != nil {
		return err
	}
	return nil
}

func (redisDataStore *RedisDataStore) Keys(ctx context.Context, pattern string) ([]string, error) {
	if redisDataStore.options.Debug {
		fmt.Println("[LOG] Keys", pattern)
	}
	keys, err := redisDataStore.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
