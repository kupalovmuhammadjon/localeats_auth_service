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

	return tokens, err
}
func (a *AuthService) Logout(ctx context.Context, req *pb.Token) (*pb.Void, error) {
	err := a.authRepo.LogOut(ctx, req.RefreshToken)

	return &pb.Void{}, err
}
func (a *AuthService) RefreshToken(ctx context.Context, req *pb.Token) (*pb.Tokens, error) {
	return &pb.Tokens{}, nil
}
