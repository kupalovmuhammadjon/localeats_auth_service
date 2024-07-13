package postgres

import (
	pb "auth_service/genproto/auth"
	"context"

	"testing"
)

func newAuthRepo() *AuthRepo {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	return &AuthRepo{Db: db}
}

// func TestRegister(t *testing.T) {
// 	a := newAuthRepo()

// 	user := pb.ReqCreateUser{
// 		Username: "qwerty",
// 		Email:    "qwerty@gmail.com",
// 		FullName: "gfd gfr",
// 		UserType: "customer" ,
// 		Password: "1234",
// 	}

// 	_, err := a.Register(&user)
// 	if err != nil {
// 		panic(err)
// 	}
	
// }

func TestLogin(t *testing.T){
	a := newAuthRepo()

	log := pb.ReqLogin{
		Email: "qwerty@gmail.com",
		Password: "1234",
	}

	_, err := a.Login(context.Background(), &log)
	if err != nil {
		panic(err)
	}
}