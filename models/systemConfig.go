package models

import (
	"auth_service/config"
	"database/sql"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type SystemConfig struct {
	Config     *config.Config
	PostgresDb *sql.DB
	RedisDb    *redis.Client
	Logger     *zap.Logger
}
