package postgres

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
// 		UserType: "customer",
// 		Password: "1234",
// 	}

// 	_, err := a.Register(context.Background(), &user)
// 	if err != nil {
// 		panic(err)
// 	}

// }

// func TestLogin(t *testing.T){
// 	a := newAuthRepo()

// 	log := pb.ReqLogin{
// 		Email: "string@gmail.com",
// 		Password: "!qwerty2345Q",
// 	}

// 	_, err := a.Login(context.Background(), &log)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestLogout(t *testing.T) {
// 	a := newAuthRepo()

// 	err := a.LogOut(context.Background(), "gf")
// 	if err == nil {
// 		log.Println("verifying invalid token")
// 	}
// }

// func TestRefreshToken(t *testing.T) {
// 	a := newAuthRepo()

// 	err := a.RefreshToken(context.Background(), "gvcxf")
// 	if err == nil {
// 		panic(fmt.Errorf("verifying invalid token"))
// 	}

// 	err = a.RefreshToken(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN0cmluZ0BnbWFpbC5jb20iLCJleHAiOjE3MjE0OTIwODEsImZ1bGxfbmFtZSI6InN0cmluZyIsImlhdCI6MTcyMDg4NzI4MSwidXNlcl9pZCI6ImJlZjJkMWU5LWEzYjAtNDBhMC04Y2E1LTM4ZDI4NjQ3ZTYzNyIsInVzZXJfdHlwZSI6ImN1c3RvbWVyIiwidXNlcm5hbWUiOiJzdHJpbmcifQ.P6v72xJGUW8U47y3J3wx86aHqjRHEB7BbBypV0uQDlY")
// 	if err != nil {
// 		panic(fmt.Errorf("not verifying valid token %v", err))
// 	}
// }
