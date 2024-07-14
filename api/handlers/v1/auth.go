package v1

import (
	"auth_service/api/handlers/tokens"
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"auth_service/pkg/validations"
	"encoding/json"
	"net/http"
	"time"

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
		h.log.Info("invalid Email ", zap.Error(err))
		return
	}
	err = validations.ValidatePassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Password",
			Error:   err.Error(),
		})
		h.log.Info("invalid password ", zap.Error(err))
		return
	}
	err = validations.ValidatePhoneNumber(req.PhoneNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid phone number",
			Error:   err.Error(),
		})
		h.log.Info("invalid phone number ", zap.Error(err))
		return
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
		h.log.Info("invalid Email ", zap.Error(err))
		return
	}
	err = validations.ValidatePassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Password",
			Error:   err.Error(),
		})
		h.log.Info("invalid password ", zap.Error(err))
		return
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

	ctx.SetCookie("access_token", tokens.AccessToken, int(time.Hour), "/", "", false, true)
	ctx.SetCookie("refresh_token", tokens.RefreshToken, int(time.Hour*24*7), "/", "", false, true)

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
	refresh, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		h.log.Info("Invalid token ", zap.Error(err))
		return
	}
	_, err = tokens.ExtractClaims(refresh, true)
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
	ctx.SetCookie("access_token", "", -1, "/", "", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)

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
	refresh, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while getting refresh token from cookie",
			Error:   err.Error(),
		})
		h.log.Info("Error while getting refresh token from cookie ", zap.Error(err))
		return
	}
	_, err = tokens.ExtractClaims(refresh, true)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid token",
			Error:   err.Error(),
		})
		h.log.Info("Invalid token ", zap.Error(err))
		h.log.Info("Invalid token ", zap.Any("input", refresh))
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

// @Summary resets password
// @Description send info about reserttting poassword to email
// @Tags Auth
// @ID reset
// @Accept json
// @Produce json
// @Param email body auth.ReqResetPassword true "email of the user"
// @Success 200
// @Failure 400 {object} models.Error "some thing wrong with what you sent"
// @Failure 401 {object} models.Error "Invalid token in header"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /auth/resetpassword [post]
func (h *HandlerV1) ResetPassword(ctx *gin.Context) {

	var body pb.ReqResetPassword

	err := json.NewDecoder(ctx.Request.Body).Decode(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding request json",
			Error:   err.Error(),
		})
		h.log.Error("Error while decoding request json ", zap.Error(err))
		return
	}
	err = validations.ValidateEmail(body.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Email",
			Error:   err.Error(),
		})
		h.log.Info("invalid Email ", zap.Error(err))
		return
	}

	status, err := h.authService.ResetPassword(ctx, &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Failed to send email",
			Error:   err.Error(),
		})
		h.log.Error("Failed to send email ", zap.Error(err))
		return
	}

	ctx.JSON(200, status)
}

// @Summary update password
// @Description updates password
// @Tags Auth
// @ID updatepassword
// @Accept json
// @Produce json
// @Param email path string true "email of the user"
// @Param password body auth.ReqUpdatePassword true "email of the user"
// @Success 200
// @Failure 400 {object} models.Error "some thing wrong with what you sent"
// @Failure 401 {object} models.Error "Invalid token in header"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /auth/updatepassword/{email} [post]
func (h *HandlerV1) UpdatePassword(ctx *gin.Context) {

	req := pb.ReqUpdatePassword{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding request json",
			Error:   err.Error(),
		})
		h.log.Error("Error while decoding request json ", zap.Error(err))
		return
	}

	email := ctx.Param("email")
	err = validations.ValidateEmail(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Email",
			Error:   err.Error(),
		})
		h.log.Info("invalid Email ", zap.Error(err))
		return
	}
	err = validations.ValidatePassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid Password",
			Error:   err.Error(),
		})
		h.log.Info("invalid password ", zap.Error(err))
		return
	}
	req.Email = email
	status, err := h.authService.UpdatePassword(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Failed to send email",
			Error:   err.Error(),
		})
		h.log.Error("Failed to send email ", zap.Error(err))
		return
	}

	ctx.JSON(200, status)
}
