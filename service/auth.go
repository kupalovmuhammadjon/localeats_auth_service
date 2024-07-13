package service

import (
	"auth_service/api/handlers/tokens"
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"auth_service/storage/postgres"
	"context"

	"go.uber.org/zap"
)

type AuthService struct {
	authRepo *postgres.AuthRepo
	log      *zap.Logger
	pb.UnimplementedAuthServer
}

func NewAuthService(sysConfig *models.SystemConfig) *AuthService {

	return &AuthService{
		authRepo: postgres.NewAuthRepo(sysConfig.PostgresDb),
		log:      sysConfig.Logger,
	}
}

func (a *AuthService) Register(ctx context.Context, req *pb.ReqCreateUser) (*pb.User, error) {
	user, err := a.authRepo.Register(ctx, req)
	if err != nil {
		a.log.Error("user is not registerd", zap.Error(err))
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
