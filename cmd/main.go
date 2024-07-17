package main

import (
	"auth_service/api"
	"auth_service/config"
	pba "auth_service/genproto/auth"
	pbk "auth_service/genproto/kitchen"
	pbu "auth_service/genproto/user"
	"auth_service/models"
	logger "auth_service/pkg/logger"
	"auth_service/service"
	"auth_service/storage/postgres"
	"auth_service/storage/redis"
	"log"
	"net"
	"sync"

	"github.com/gin-gonic/gin"
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

	router := api.NewRouter(systemConfig)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go runHttpServer(router, cfg, &wg)
	go runGrpcServer(systemConfig, &wg)

	wg.Wait()
}

func runHttpServer(router *gin.Engine, cfg *config.Config, wait *sync.WaitGroup) {
	defer wait.Done()

	err := router.Run(cfg.HTTP_PORT)
	if err != nil {
		log.Fatal("Server failed to run ", zap.Error(err))
		return
	}
}

func runGrpcServer(sysConfig *models.SystemConfig, wait *sync.WaitGroup) {
	defer wait.Done()

	listener, err := net.Listen("tcp", sysConfig.Config.AUTH_SERVICE_PORT)
	if err != nil {
		log.Fatal("can not create listener ", zap.Error(err))
		return
	}

	server := grpc.NewServer()

	pbu.RegisterUserServiceServer(server, service.NewUserService(sysConfig))
	pba.RegisterAuthServer(server, service.NewAuthService(sysConfig))
	pbk.RegisterKitchenServer(server, service.NewKitchenService(sysConfig))

	sysConfig.Logger.Info("User service is started working ")
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve to listener listener ", zap.Error(err))
		return
	}
}
