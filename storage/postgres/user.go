package postgres

import (
	pb "auth_service/genproto/user"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

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
		is_verified,
		created_at,
		updated_at
	from
		users
	where 
		id = $1
	`
	user := pb.User{}
	var address sql.NullString
	var bio sql.NullString
	var yearsOfExperience sql.NullInt32
	var isVerified sql.NullBool
	row := u.Db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.FullName, &user.UserType, &address,
		&user.PhoneNumber, &bio, pq.Array(&user.Specialties), &yearsOfExperience, &isVerified,
		&user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	user.Address = address.String
	user.Bio = bio.String
	user.YearsOfExperience = yearsOfExperience.Int32
	user.IsVerified = isVerified.Bool

	return &user, row.Err()
}

func (u *UserRepo) UpdateProfile(ctx context.Context, user *pb.ReqUpdateUser) (*pb.User, error) {
	// Convert the specialties slice to a comma-separated string
	specialties := strings.Join(user.Specialties, ",")

	query := `
	UPDATE
		users
	SET
		username = $1,
		email = $2,
		full_name = $3,
		address = $4,
		phone_number = $5,
		bio = $6,
		specialties = string_to_array($7, ','),
		years_of_experience = $8,
		is_verified = $9,
		updated_at = now()
	WHERE
		id = $10
	RETURNING
		id, username, email, full_name, address, phone_number, bio, specialties, years_of_experience, is_verified, created_at, updated_at
	`

	var updatedUser pb.User

	err := u.Db.QueryRowContext(ctx, query, user.Username, user.Email, user.FullName, user.Address, user.PhoneNumber,
		user.Bio, specialties, user.YearsOfExperience, user.IsVerified, user.Id).Scan(
		&updatedUser.Id,
		&updatedUser.Username,
		&updatedUser.Email,
		&updatedUser.FullName,
		&updatedUser.Address,
		&updatedUser.PhoneNumber,
		&updatedUser.Bio,
		pq.Array(&updatedUser.Specialties),
		&updatedUser.YearsOfExperience,
		&updatedUser.IsVerified,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, id string) error {
	query := `
	update
		users
	set
		deleted_at = now()
	where
		id = $1
	`
	res, err := u.Db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("nothing deleted")
	}
	return nil
}

func (u *UserRepo) ValidateUserId(ctx context.Context, id string) error {
	query := `
	SELECT 
		1
	FROM 
		users
	WHERE 
		id = $1

	`

	var exists int
	err := u.Db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user ID %s does not exist", id)
		}
		return fmt.Errorf("error checking user ID %s: %v", id, err)
	}

	return nil
}

func (u *UserRepo) CreateUserPreference(ctx context.Context, pref *pb.Preferences) (*pb.PreferencesRes, error) {
	query := `
	insert into
		user_preferences(
		user_id,
		cuisine_type,
		dietary_preferences,
		favorite_kitchen_ids,
		created_at,
		updated_at)
	values($1, $2, $3, $4, $5, $6)
	`

	currentTime := time.Now().Format(time.RFC3339)
	res := pb.PreferencesRes{
		UserId:             pref.UserId,
		CuisineType:        pref.CuisineType,
		DietaryPreferences: pref.DietaryPreferences,
		FavoriteKitchenIds: pref.FavoriteKitchenIds,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}

	_, err := u.Db.ExecContext(ctx, query, res.UserId, res.CuisineType, pq.Array(res.DietaryPreferences),
		pq.Array(res.FavoriteKitchenIds), res.CreatedAt, res.UpdatedAt)

	return &res, err
}

func (u *UserRepo) UpdateUserPreference(ctx context.Context, pref *pb.Preferences) (*pb.PreferencesRes, error) {
	query := `
	update
		user_preferences
	set
		cuisine_type = $1,
		dietary_preferences = $2,
		favorite_kitchen_ids = $3,
		updated_at = now()
	where
		user_id = $4
	returning
		user_id,
		cuisine_type,
		dietary_preferences,
		favorite_kitchen_ids,
		created_at,
		updated_at
	`

	res := pb.PreferencesRes{}

	row := u.Db.QueryRowContext(ctx, query, pref.CuisineType, pq.Array(pref.DietaryPreferences),
		pq.Array(pref.FavoriteKitchenIds), pref.UserId)

	err := row.Scan(&res.UserId, &res.CuisineType, pq.Array(&res.DietaryPreferences), pq.Array(&res.FavoriteKitchenIds),
		&res.CreatedAt, &res.UpdatedAt)

	return &res, err
}

func (u *UserRepo) GetUserPreference(ctx context.Context, id string) (*pb.PreferencesRes, error) {
	query := `
	select
		user_id,
		cuisine_type,
		dietary_preferences,
		favorite_kitchen_ids,
		created_at,
		updated_at
	from
		user_preferences
	where
		user_id = $1
	`

	res := pb.PreferencesRes{}

	row := u.Db.QueryRowContext(ctx, query, id)

	err := row.Scan(&res.UserId, &res.CuisineType, pq.Array(&res.DietaryPreferences), pq.Array(&res.FavoriteKitchenIds),
		&res.CreatedAt, &res.UpdatedAt)

	return &res, err
}

func (u *UserRepo) DeleteUserPreference(ctx context.Context, id string) error {
	query := `
	update
		user_preferences
	set
		deleted_at = now()
	where
		user_id = $1
	`

	_, err := u.Db.ExecContext(ctx, query, id)

	return err
}
