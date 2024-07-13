package redis

import (
	"auth_service/config"

	"github.com/go-redis/redis"
)

func ConnectDB() (*redis.Client, error) {
	config := config.Load()
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST + ":" + config.REDIS_PORT,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return client, nil
}
