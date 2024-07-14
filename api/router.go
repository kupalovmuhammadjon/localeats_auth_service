package api

import (
	_ "auth_service/api/docs"
	v1 "auth_service/api/handlers/v1"
	"auth_service/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title LocalEats API
// @version 1.0
// @description LocalEats is a program to order local and homemade food with quality and precise delivery.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9999
// @BasePath /localeats.uz

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewRouter(sysConfig *models.SystemConfig) *gin.Engine {

	handlerV1 := v1.NewHandlerV1(sysConfig)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	main := router.Group("/localeats.uz")

	auth := main.Group("/auth")

	auth.POST("/register", handlerV1.Register)
	auth.POST("/login", handlerV1.Login)
	auth.POST("/logout", handlerV1.Logout)
	auth.POST("/refreshtoken", handlerV1.RefreshToken)
	auth.POST("/resetpassword", handlerV1.ResetPassword)
	auth.POST("/updatepassword/:email", handlerV1.UpdatePassword)

	return router
}
