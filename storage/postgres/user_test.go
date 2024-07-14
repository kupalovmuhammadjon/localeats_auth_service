package postgres

import (
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

	_, err := u.GetProfile(context.Background(), "0f7c65e6-03a9-415f-ab63-804a5654068a")
	if err != nil {
		panic(err)
	}
}
