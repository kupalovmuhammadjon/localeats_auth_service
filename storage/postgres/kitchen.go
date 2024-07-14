package postgres

import (
	pb "auth_service/genproto/kitchen"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type KitchenRepo struct {
	Db *sql.DB
}

func NewKitchenRepo(db *sql.DB) *KitchenRepo {
	return &KitchenRepo{Db: db}
}

func (k *KitchenRepo) CreateKitchen(ctx context.Context, kitchen *pb.ReqCreateKitchen) (*pb.KitchenInfo, error) {
	query := `
	insert into
		kitchens(
		id,
    	owner_id,
    	name,
    	description,
    	cuisine_type,
    	address,
    	phone_number,
    	created_at,
    	updated_at)
	values($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	kt := pb.KitchenInfo{
		Id:          uuid.NewString(),
		OwnerId:     kitchen.OwnerId,
		Name:        kitchen.Name,
		Description: kitchen.Description,
		CuisineType: kitchen.CuisineType,
		Address:     kitchen.Address,
		PhoneNumber: kitchen.PhoneNumber,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	_, err := k.Db.ExecContext(ctx, query, kt.Id, kt.OwnerId, kt.Name, kt.Description, kt.CuisineType, kt.Address,
		kt.PhoneNumber, kt.CreatedAt, kt.UpdatedAt)

	return &kt, err
}
