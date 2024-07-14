package postgres

import (
	pb "auth_service/genproto/user"
	"context"
	"testing"
)

func newUserRepo() *UserRepo {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	return &UserRepo{
		Db: db,
	}
}

func TestGetProfile(t *testing.T) {
	u := newUserRepo()

	_, err := u.GetProfile(context.Background(), "18483d59-75e8-4cbc-a2e9-22d981b09d34")
	if err != nil {
		panic(err)
	}
}

func TestUpdateProfile(t *testing.T) {
	u := newUserRepo()

	user := pb.ReqUpdateUser{
		Id: "18483d59-75e8-4cbc-a2e9-22d981b09d34",
		Username: "qwerty",
		Email:    "qwerty@gmail.com",
		FullName: "gfd gfr",
	}

	_, err := u.UpdateProfile(context.Background(), &user)
	if err != nil {
		panic(err)
	}
}

func TestDeleteUser(t *testing.T) {
	u := newUserRepo()

	err := u.DeleteUser(context.Background(), "18483d59-75e8-4cbc-a2e9-22d981b09d34")
	if err != nil {
		panic(err)
	}
}
