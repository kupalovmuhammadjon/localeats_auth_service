package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT           string
	AUTH_SERVICE_PORT   string
	DB_HOST             string
	DB_PORT             string
	DB_NAME             string
	DB_USER             string
	DB_PASSWORD         string
	ACCESS_SIGNING_KEY  string
	REFRESH_SIGNING_KEY string
	REDIS_HOST          string
	REDIS_PORT          string
	REDIS_PASSWORD      string
	LOG_PATH            string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error while loading .env file")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":23456"))
	config.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":50053"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "name"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))
	config.ACCESS_SIGNING_KEY = cast.ToString(coalesce("ACCESS_SIGNING_KEY", "root"))
	config.REFRESH_SIGNING_KEY = cast.ToString(coalesce("REFRESH_SIGNING_KEY", "root"))
	config.REDIS_HOST = cast.ToString(coalesce("REDIS_HOST", "root"))
	config.REDIS_PORT = cast.ToString(coalesce("REDIS_PORT", "root"))
	config.REDIS_PASSWORD = cast.ToString(coalesce("REDIS_PASSWORD", "root"))
	config.LOG_PATH = cast.ToString(coalesce("REDIS_PASSWORD", "areyouinterested.log"))

	return &config
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
