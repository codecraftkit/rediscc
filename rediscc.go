package rediscc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Connect(redisUri string, dbNumber string, dataStore *RedisDataStore) error {
	redisUrl := fmt.Sprintf("%s/%s", redisUri, dbNumber)
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return err
	}

	client := redis.NewClient(opt)

	fmt.Printf("You successfully connected to Redis: %s\n", redisUrl)

	dataStore.Client = client

	return nil
}
