package main

import (
	pba "auth_service/genproto/auth"
	pbu "auth_service/genproto/user"
	"auth_service/service"

	"auth_service/config"
	"auth_service/models"
	"auth_service/pkg/logger"
	"auth_service/storage/postgres"
	"auth_service/storage/redis"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	listener, err := net.Listen("tcp", cfg.AUTH_SERVICE_PORT)
	if err != nil {
		log.Fatal("can not create listener ", zap.Error(err))
		return
	}

	server := grpc.NewServer()

	pbu.RegisterUserServiceServer(server, service.NewUserService(systemConfig))
	pba.RegisterAuthServer(server, service.NewAuthService(systemConfig))

	log.Info("User service is started working ")
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve to listener listener ", zap.Error(err))
		return
	}
}
