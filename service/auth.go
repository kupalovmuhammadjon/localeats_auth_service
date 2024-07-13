package service

import (
	pb "auth_service/genproto/auth"
	"auth_service/storage/postgres"
	"context"

	"go.uber.org/zap"
)

type AuthService struct {
	AuthRepo *postgres.AuthRepo
	log      *zap.Logger
	pb.UnimplementedAuthServer
}

func (a *AuthService) Register(ctx context.Context, req *pb.ReqCreateUser) (*pb.User, error) {
	user, err := a.AuthRepo.Register(ctx, req)
	if err != nil {
		a.log.Error("user is not registerd", zap.Error(err))
		return nil, err
	}
	return user, nil
}
func (a *AuthService) Login(ctx context.Context, req *pb.ReqLogin) (*pb.Tokens, error) {
	user, err := a.AuthRepo.Login(ctx, req)
	if err != nil {
		a.log.Error("user is not registerd", zap.Error(err))
		return nil, err
	}
	
	return user, nil
}
func (a *AuthService) Logout(ctx context.Context, req *pb.Token) (*pb.Void, error) {

}
func (a *AuthService) RefreshToken(ctx context.Context, req *pb.Token) (*pb.Tokens, error) {

}
