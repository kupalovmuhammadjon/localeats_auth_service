package main

import (
	"auth_service/api"
	"auth_service/config"
	"auth_service/models"
	logger "auth_service/pkg/logger"
	"auth_service/storage/postgres"
	"auth_service/storage/redis"

	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log, err := logger.New("debug", "development", cfg.LOG_PATH)
	if err != nil {
		panic(err)
	}

	postgresDb, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to Postgres", zap.Error(err))
		return
	}
	defer postgresDb.Close()

	redisDb, err := redis.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to Redis", zap.Error(err))
		return
	}
	defer redisDb.Close()

	systemConfig := &models.SystemConfig{
		Config:     cfg,
		PostgresDb: postgresDb,
		RedisDb:    redisDb,
		Logger:     log,
	}

	router := api.NewRouter(systemConfig)

	err = router.Run(cfg.HTTP_PORT)
	if err != nil {
		log.Fatal("Server failed to run ", zap.Error(err))
		return
	}
}
