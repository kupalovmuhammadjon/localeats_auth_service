package service

import (
	pb "auth_service/genproto/kitchen"
	"auth_service/models"
	"auth_service/storage/postgres"
	"context"

	"go.uber.org/zap"
)

type KitchenService struct {
	kitchenRepo *postgres.KitchenRepo
	log      *zap.Logger
	pb.UnimplementedKitchenServer
}

func NewKitchenService(sysConfig *models.SystemConfig) *KitchenService {

	return &KitchenService{
		kitchenRepo: postgres.NewKitchenRepo(sysConfig.PostgresDb),
		log:      sysConfig.Logger,
	}
}

func (k *KitchenService) CreateKitchen(ctx context.Context, req *pb.ReqCreateKitchen) (*pb.KitchenInfo, error){
	kitchen, err := k.kitchenRepo.CreateKitchen(ctx, req)
	if err != nil {
		k.log.Error("failed to create kitchen ", zap.Error(err))
		return nil, err
	}

	return kitchen, nil
}

func (k *KitchenService) UpdateKitchen(ctx context.Context, req *pb.ReqUpdateKitchen) (*pb.KitchenInfo, error){
	kitchen, err := k.kitchenRepo.UpdateKitchen(ctx, req)
	if err != nil {
		k.log.Error("failed to update kitchen ", zap.Error(err))
		return nil, err
	}
	
	return kitchen, nil
}

func (k *KitchenService) GetKitchenById(ctx context.Context, req *pb.Id) (*pb.KitchenInfo, error){
	kitchen, err := k.kitchenRepo.GetKitchenById(ctx, req.Id)
	if err != nil {
		k.log.Error("failed to get kitchen by id ", zap.Error(err))
		return nil, err
	}
	
	return kitchen, nil
}

func (k *KitchenService) GetKitchens(ctx context.Context, req *pb.Pagination) (*pb.Kitchens, error){
	kitchens, err := k.kitchenRepo.GetKitchens(ctx, req)
	if err != nil {
		k.log.Error("failed to get kitchens ", zap.Error(err))
		return nil, err
	}
	
	return kitchens, nil
}

func (k *KitchenService) SearchKitchens(ctx context.Context, req *pb.Search) (*pb.Kitchens, error){
	kitchens, err := k.kitchenRepo.SearchKitchens(ctx, req)
	if err != nil {
		k.log.Error("failed to search kitchens ", zap.Error(err))
		return nil, err
	}
	
	return kitchens, nil
}

