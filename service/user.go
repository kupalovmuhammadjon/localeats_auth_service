package service

import (
	pb "auth_service/genproto/user"
	"auth_service/models"
	"auth_service/storage/postgres"
	"context"

	"go.uber.org/zap"
)

type UserService struct {
	userRepo *postgres.UserRepo
	log      *zap.Logger
	pb.UnimplementedUserServiceServer
}

func NewUserService(sysConfig *models.SystemConfig) *UserService {

	return &UserService{
		userRepo: postgres.NewUserRepo(sysConfig.PostgresDb),
		log:      sysConfig.Logger,
	}
}

func (u *UserService) GetProfile(ctx context.Context, id *pb.Id) (*pb.User, error) {
	user, err := u.userRepo.GetProfile(ctx, id.Id)
	if err != nil {
		u.log.Error("can not get user profile ", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *UserService) UpdateProfile(ctx context.Context, user *pb.ReqUpdateUser) (*pb.User, error) {
	updatedUser, err := u.userRepo.UpdateProfile(ctx, user)
	if err != nil {
		u.log.Error("can not update user profile ", zap.Error(err))
		return nil, err
	}
	return updatedUser, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id *pb.Id) (*pb.Status, error) {
	err := u.userRepo.DeleteUser(ctx, id.Id)
	if err != nil {
		u.log.Error("Error while deleting user ", zap.Error(err))
		return nil, err
	}
	return &pb.Status{Message: "Deleted successfully"}, nil
}

