package redis

import (
	"auth_service/config"
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectDB() (*redis.Client, error) {
	config := config.Load()
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST + ":" + config.REDIS_PORT,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	
	return client, nil
}
