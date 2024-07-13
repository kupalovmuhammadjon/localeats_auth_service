package postgres

import (
	pb "auth_service/genproto/auth"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct {
	Db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo{
	return &AuthRepo{Db: db}
}

func (a *AuthRepo) Register(ctx context.Context, user *pb.ReqCreateUser) (*pb.User, error) {
	query := `
	insert into
		users(id, username, email, password_hash, full_name, user_type, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7, $8)
	`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userRes := pb.User{
		Id:        uuid.NewString(),
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		UserType:  user.UserType,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	_, err = a.Db.Exec(query, userRes.Id, userRes.Username, userRes.Email, string(hashedPassword),
		userRes.FullName, userRes.UserType, userRes.CreatedAt, userRes.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &userRes, nil
}

func (a *AuthRepo) Login(ctx context.Context, credentials *pb.ReqLogin) (*pb.UserClaims, error) {
	query := `
	select
		id, username, email, full_name, user_type, password_hash
	from
		users
	where 
		deleted_at is null and email = $1
	`

	row := a.Db.QueryRowContext(ctx, query, credentials.Email)

	var hashedPassword string
	user := pb.UserClaims{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.UserType, &hashedPassword)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		return nil, err
	}

	return &user, row.Err()
}

func (a *AuthRepo) LogOut(ctx context.Context, token string) error {
	query := `
	update 
	set
		deleted_at = now()
	from
		refresh_tokens
	where
		token = $1
	`

	res, err := a.Db.ExecContext(ctx, query, token)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return  err
	}
	if affectedRows == 0{
		return fmt.Errorf("refresh token not found")
	}

	return nil
}

func (a *AuthRepo) RefreshToken(ctx context.Context, token string) error {
	query := `
	select
		case when token = $1 then true else false
	from
		refresh_tokens
	where
		deleted_at is null and token = $1
	`

	exists := false
	err := a.Db.QueryRowContext(ctx, query, token).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists{
		return fmt.Errorf("token does not exists")
	}

	return nil
}
