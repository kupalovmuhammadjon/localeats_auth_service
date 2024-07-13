package models

type ReqCreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	UserType string `json:"user_type"`
}

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	UserType  string `json:"user_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type ReqRefresh struct {
	RefreshToken string `json:"refresh_token"`
}
