package v1

import (
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Register User
// @Description Registers user
// @Tags Auth
// @ID register
// @Accept json
// @Produce json
// @Param user body auth.ReqCreateUser true "User information to create it"
// @Success 201
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /auth/register [post]
func (h *HandlerV1) Register(ctx *gin.Context) {
	req := pb.ReqCreateUser{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding request json",
			Error: err.Error(),
		})
		h.log.Error("Error while decoding request json ", zap.Error(err))
		return
	}

	user, err := h.authService.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while creating new user",
			Error: err.Error(),
		})
		h.log.Error("Error while creating new user ", zap.Error(err))
		return
	}

	ctx.JSON(201, user)
}
