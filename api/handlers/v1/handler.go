package v1

import (
	"auth_service/storage/postgres"

	"go.uber.org/zap"
)

type HandlerV1 struct{
	log *zap.Logger
	authRepo *postgres.AuthRepo
	
}