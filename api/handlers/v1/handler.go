package v1

import (
	"auth_service/models"
	"auth_service/service"
	"go.uber.org/zap"
)

type HandlerV1 struct {
	log         *zap.Logger
	authService *service.AuthService
}

func NewHandlerV1(sysConfig *models.SystemConfig) *HandlerV1 {
	return &HandlerV1{
		log:         sysConfig.Logger,
		authService: service.NewAuthService(sysConfig),
	}
}
