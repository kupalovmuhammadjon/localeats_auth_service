package postgres

import (
	pb "auth_service/genproto/kitchen"
	"context"
	"database/sql"
	"fmt"
	"log"
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

func (k *KitchenRepo) UpdateKitchen(ctx context.Context, kitchen *pb.ReqUpdateKitchen) (*pb.KitchenInfo, error) {
	query := `
		UPDATE kitchens
		SET 
			owner_id = $1, 
			name = $2, 
			description = $3, 
			cuisine_type = $4, 
			address = $5, 
			phone_number = $6,
			updated_at = NOW()
		WHERE 
			id = $7 AND deleted_at IS NULL
		RETURNING 
			id, owner_id, name, description, cuisine_type, address, phone_number, rating, total_orders, created_at, updated_at
	`

	params := []interface{}{
		kitchen.OwnerId,
		kitchen.Name,
		kitchen.Description,
		kitchen.CuisineType,
		kitchen.Address,
		kitchen.PhoneNumber,
		kitchen.Id,
	}

	res := pb.KitchenInfo{}
	row := k.Db.QueryRowContext(ctx, query, params...)
	err := row.Scan(&res.Id, &res.OwnerId, &res.Name, &res.Description, &res.CuisineType, &res.Address,
		&res.PhoneNumber, &res.Rating, &res.TotalOrders, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		log.Printf("Query execution error: %v", err)
		return nil, err
	}

	return &res, nil
}

func (k *KitchenRepo) GetKitchenById(ctx context.Context, id string) (*pb.KitchenInfo, error) {
	query := `
        SELECT 
            id, owner_id, name, description, cuisine_type, address, phone_number, rating, total_orders, created_at, updated_at 
        FROM 
            kitchens 
        WHERE 
            id = $1 AND deleted_at IS NULL
    `

	var res pb.KitchenInfo
	row := k.Db.QueryRowContext(ctx, query, id)
	err := row.Scan(&res.Id, &res.OwnerId, &res.Name, &res.Description, &res.CuisineType, &res.Address,
		&res.PhoneNumber, &res.Rating, &res.TotalOrders, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No kitchen found with id: %s", id)
			return nil, nil
		}
		log.Printf("Query execution error: %v", err)
		return nil, err
	}

	return &res, nil
}

func (k *KitchenRepo) GetKitchens(ctx context.Context, filter *pb.Pagination) (*pb.Kitchens, error) {
	query := `
	SELECT 
            id, name, cuisine_type, address, rating, total_orders 
        FROM 
            kitchens 
        WHERE 
            deleted_at IS NULL
	`
	query += fmt.Sprintf(" offset %d", (filter.Page-1)*filter.Limit)
	query += fmt.Sprintf(" limit %d", filter.Limit)

	rows, err := k.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	kitchens := pb.Kitchens{}

	for rows.Next() {
		kitchen := pb.KitchenShortInfo{}

		err := rows.Scan(&kitchen.Id, &kitchen.Name, &kitchen.CuisineType, &kitchen.Address, &kitchen.Rating, &kitchen.TotalOrders)
		if err != nil {
			return nil, err
		}
		kitchens.Kitchens = append(kitchens.Kitchens, &kitchen)
	}

	return &kitchens, rows.Err()
}

func (k *KitchenRepo) SearchKitchens(ctx context.Context, search *pb.Search) (*pb.Kitchens, error) {
	query := `
	select 
            id, name, cuisine_type, address, rating, total_orders 
        from 
            kitchens 
        where 
            deleted_at IS NULL and 
			to_tsvector(name || ' ' || cuisine_type || ' ' || address || ' ' || rating::text) @@ plainto_tsquery($1)
	`

	query += fmt.Sprintf(" offset %d", (search.Page-1)*search.Limit)
	query += fmt.Sprintf(" limit %d", search.Limit)

	text := fmt.Sprintf("%s %s %s %v.00", search.Name, search.CuisineType, search.Address, search.Rating)
	rows, err := k.Db.QueryContext(ctx, query, text)
	fmt.Println(query)
	fmt.Println(text)
	if err != nil {
		return nil, err
	}

	kitchens := pb.Kitchens{}

	for rows.Next() {
		kitchen := pb.KitchenShortInfo{}

		err := rows.Scan(&kitchen.Id, &kitchen.Name, &kitchen.CuisineType, &kitchen.Address, &kitchen.Rating, &kitchen.TotalOrders)
		if err != nil {
			return nil, err
		}
		kitchens.Kitchens = append(kitchens.Kitchens, &kitchen)
	}

	return &kitchens, rows.Err()
}

func (k *KitchenRepo) DeleteKitchen(ctx context.Context, id string) error {
	query := `
		UPDATE 
			kitchens
		SET 
			deleted_at = NOW()
		WHERE 
			id = $1 AND deleted_at IS NULL
	`

	_, err := k.Db.ExecContext(ctx, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No kitchen found with id: %s or it is already deleted", id)
			return nil 
		}
		return err
	}

	return nil
}


func (k *KitchenRepo) ValidateKitchenId(ctx context.Context, id string) error {
	query := `
	SELECT 
		1
	FROM 
		kitchens
	WHERE 
		id = $1

	`

	var exists int
	err := k.Db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("kitchen ID %s does not exist", id)
		}
		return fmt.Errorf("error checking kitchen ID %s: %v", id, err)
	}

	return nil
}