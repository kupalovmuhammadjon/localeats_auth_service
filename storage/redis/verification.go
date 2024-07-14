package redis

import (
	"auth_service/config"
	"auth_service/models"
	"auth_service/pkg/verification"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type VerificationRepo struct {
	Db     *redis.Client
	Config *config.Config
}

func NewVerificationRepo(sysConfig *models.SystemConfig) *VerificationRepo {
	return &VerificationRepo{
		Db:     sysConfig.RedisDb,
		Config: sysConfig.Config,
	}
}

func (v *VerificationRepo) SendVerificationToEmail(ctx context.Context, email string) error {
	url := "http://localhost:9999/localeats.uz/auth/updatepassword/"
	url += email

	err := v.Db.Set(ctx, email, url, time.Minute*10).Err()
	if err != nil {
		return err
	}

	err = verification.SendVerificationToEmail(v.Config, email, url)
	if err != nil {
		return err
	}
	return nil
}
