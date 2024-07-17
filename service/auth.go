package service

import (
	"auth_service/api/handlers/tokens"
	pb "auth_service/genproto/auth"
	pbu "auth_service/genproto/user"
	"auth_service/models"
	"auth_service/storage/postgres"
	"auth_service/storage/redis"
	"context"

	"go.uber.org/zap"
)

type AuthService struct {
	authRepo         *postgres.AuthRepo
	userRepo         *postgres.UserRepo
	verificationRepo *redis.VerificationRepo
	log              *zap.Logger
	pb.UnimplementedAuthServer
}

func NewAuthService(sysConfig *models.SystemConfig) *AuthService {

	return &AuthService{
		authRepo:         postgres.NewAuthRepo(sysConfig.PostgresDb),
		userRepo:         postgres.NewUserRepo(sysConfig.PostgresDb),
		verificationRepo: redis.NewVerificationRepo(sysConfig),
		log:              sysConfig.Logger,
	}
}

func (a *AuthService) Register(ctx context.Context, req *pb.ReqCreateUser) (*pb.User, error) {
	user, err := a.authRepo.Register(ctx, req)
	if err != nil {
		a.log.Error("user is not registerd", zap.Error(err))
		return nil, err
	}
	_, err = a.userRepo.CreateUserPreference(ctx, &pbu.Preferences{UserId: user.Id})
	if err != nil {
		a.log.Error("failed to create user precerences", zap.Error(err))
		return nil, err
	}
	return user, nil
}
func (a *AuthService) Login(ctx context.Context, req *pb.ReqLogin) (*pb.Tokens, error) {
	user, err := a.authRepo.Login(ctx, req)
	if err != nil {
		a.log.Error("user is not registerd", zap.Error(err))
		return nil, err
	}
	tokens, err := tokens.GenerateJWT(user)
	if err != nil {
		a.log.Error("Tokens are not generated ", zap.Error(err))
		return nil, err
	}
	err = a.authRepo.WriteRefreshToken(ctx, user, tokens.RefreshToken)
	if err != nil {
		a.log.Error("Refresh token is not inserted ", zap.Error(err))
		return nil, err
	}

	return tokens, nil
}
func (a *AuthService) Logout(ctx context.Context, req *pb.Token) (*pb.Void, error) {
	err := a.authRepo.LogOut(ctx, req.RefreshToken)

	return &pb.Void{}, err
}

func (a *AuthService) RefreshToken(ctx context.Context, req *pb.Token) (*pb.Tokens, error) {

	err := a.authRepo.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		a.log.Error("Error while refresh token ", zap.Error(err))
		return nil, err
	}
	tokens, err := tokens.GenerateAccessToken(req.RefreshToken)
	if err != nil {
		a.log.Error("Access token is not generated ", zap.Error(err))
		return nil, err
	}

	return tokens, nil
}

func (a *AuthService) ResetPassword(ctx context.Context, req *pb.ReqResetPassword) (*pb.Status, error) {

	err := a.authRepo.CheckUserExists(ctx, req.Email)
	if err != nil {
		a.log.Info("User does not exists ", zap.Error(err))
		return &pb.Status{Message: "User does not exists"}, err
	}

	err = a.verificationRepo.SendVerificationToEmail(ctx, req.Email)
	if err != nil {
		a.log.Error("Failed to send email ", zap.Error(err))
		return &pb.Status{Message: "Failed to send email"}, err
	}

	return &pb.Status{Message: "Email sent"}, nil
}

func (a *AuthService) UpdatePassword(ctx context.Context, req *pb.ReqUpdatePassword) (*pb.Status, error) {

	err := a.authRepo.CheckUserExists(ctx, req.Email)
	if err != nil {
		a.log.Info("User does not exists ", zap.Error(err))
		return &pb.Status{Message: "User does not exists"}, err
	}

	err = a.authRepo.UpdatePassword(ctx, req)
	if err != nil {
		a.log.Error("Failed to update password ", zap.Error(err))
		return &pb.Status{Message: "Failed to update password "}, err
	}

	return &pb.Status{Message: "Password updated"}, nil
}
