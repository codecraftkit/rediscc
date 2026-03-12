package rediscc

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context, redisUri string, dbNumber string, options *RedisOptions) (*RedisDataStore, error) {
	redisUrl := fmt.Sprintf("%s/%s", redisUri, dbNumber)
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	fmt.Printf("Successfully connected to Redis db: %s\n", dbNumber)

	if options == nil {
		options = &RedisOptions{}
	}

	redisDataStore := &RedisDataStore{
		Client:  client,
		Options: options,
	}

	return redisDataStore, nil
}

func (redisDataStore *RedisDataStore) Publish(ctx context.Context, channel string, payload interface{}) error {
	if redisDataStore.Options.Debug {
		fmt.Println("[LOG] Publish", channel)
	}
	if redisDataStore.Options.DebugPayload {
		fmt.Println("[LOG] Payload", payload)
	}
	return redisDataStore.Client.Publish(ctx, channel, payload).Err()
}

func (redisDataStore *RedisDataStore) Get(ctx context.Context, key string) (string, error) {
	if redisDataStore.Options.Debug {
		fmt.Println("[LOG] Get", key)
	}
	value, err := redisDataStore.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (redisDataStore *RedisDataStore) GetRaw(ctx context.Context, key string) *redis.StringCmd {
	if redisDataStore.Options.Debug {
		fmt.Println("[LOG] GetRaw", key)
	}
	return redisDataStore.Client.Get(ctx, key)
}

func (redisDataStore *RedisDataStore) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if redisDataStore.Options.Debug {
		fmt.Println("[LOG] Set", key, value, expiration)
	}
	return redisDataStore.Client.Set(ctx, key, value, expiration).Err()
}

func (redisDataStore *RedisDataStore) Del(ctx context.Context, key string) error {
	if redisDataStore.Options.Debug {
		fmt.Println("[LOG] Del", key)
	}
	_, err := redisDataStore.Client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (redisDataStore *RedisDataStore) Keys(ctx context.Context, pattern string) ([]string, error) {
	if redisDataStore.Options.Debug {
		fmt.Println("[LOG] Keys", pattern)
	}
	keys, err := redisDataStore.Client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
