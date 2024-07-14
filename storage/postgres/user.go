package postgres

import (
	pb "auth_service/genproto/user"
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) GetProfile(ctx context.Context, id string) (*pb.User, error) {
	query := `
	select
		id,
		username,
		email,
		full_name,
		user_type, 
		address,
		phone_number,
		bio,
		specialties, 
		years_of_experience,
		is_verified
	from
		users
	where 
		id = $1
	`
	user := pb.User{}
	var address sql.NullString
	row := u.Db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.UserType, &address, &user.PhoneNumber,
		&user.Bio, pq.Array(&user.Specialties), &user.YearsOfExperience, &user.IsVerified)

	if err != nil {
		return nil, err
	}

	return &user, row.Err()
}
