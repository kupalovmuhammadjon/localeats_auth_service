package v1

import (
	"auth_service/api/handlers/tokens"
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"auth_service/pkg/validations"
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
			Error:   err.Error(),
		})
		h.log.Error("Error while decoding request json ", zap.Error(err))
		return
	}

	err = validations.ValidateEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Email",
			Error:   err.Error(),
		})
	}
	err = validations.ValidatePassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Password",
			Error:   err.Error(),
		})
	}

	user, err := h.authService.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while creating new user",
			Error:   err.Error(),
		})
		h.log.Error("Error while creating new user ", zap.Error(err))
		return
	}

	ctx.JSON(201, user)
}

// @Summary Login user
// @Description checks the user and returns tokens
// @Tags Auth
// @ID login
// @Accept json
// @Produce json
// @Param user body auth.ReqLogin true "User Information to log in"
// @Success 200 {object} auth.Tokens  "Returns access and refresh tokens"
// @Failure 400 {object} models.Error "You did something wrong"
// @Failure 401 {object} models.Error "if Access token fails it will returns this"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /auth/login [post]
func (h *HandlerV1) Login(ctx *gin.Context) {
	req := pb.ReqLogin{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding request json",
			Error:   err.Error(),
		})
		h.log.Error("Error while decoding request json ", zap.Error(err))
		return
	}

	err = validations.ValidateEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Email",
			Error:   err.Error(),
		})
	}
	err = validations.ValidatePassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Password",
			Error:   err.Error(),
		})
	}

	tokens, err := h.authService.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while finding user",
			Error:   err.Error(),
		})
		h.log.Error("Error while finding user ", zap.Error(err))
		return
	}

	ctx.JSON(200, tokens)
}

// @Summary log outs user
// @Description removes refresh token gets token from header
// @Tags Auth
// @ID logout
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} models.Error "some thing wrong with what you sent"
// @Failure 401 {object} models.Error "Invalid token in header"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /auth/logout [post]
func (h *HandlerV1) Logout(ctx *gin.Context) {
	refresh := ctx.GetHeader("refresh_token")
	_, err := tokens.ExtractClaims(refresh, true)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		h.log.Info("Invalid token ", zap.Error(err))
		return
	}

	tokens, err := h.authService.Logout(ctx, &pb.Token{RefreshToken: refresh})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while finding user",
			Error:   err.Error(),
		})
		h.log.Error("Error while finding user ", zap.Error(err))
		return
	}

	ctx.JSON(200, tokens)
}

// @Summary refresh token
// @Description gives new access token through refresh token
// @Tags Auth
// @ID refresh
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} models.Error "some thing wrong with what you sent"
// @Failure 401 {object} models.Error "Invalid token in header"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /auth/refreshtoken [post]
func (h *HandlerV1) RefreshToken(ctx *gin.Context) {
	refresh := ctx.GetHeader("refresh_token")
	_, err := tokens.ExtractClaims(refresh, true)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		h.log.Info("Invalid token ", zap.Error(err))
		return
	}

	tokens, err := h.authService.RefreshToken(ctx, &pb.Token{RefreshToken: refresh})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while refreshing",
			Error:   err.Error(),
		})
		h.log.Error("Error while refreshing ", zap.Error(err))
		return
	}

	ctx.JSON(200, tokens)
}
