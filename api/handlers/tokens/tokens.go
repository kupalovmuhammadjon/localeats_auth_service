package tokens

import (
	"auth_service/config"
	pb "auth_service/genproto/auth"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user *pb.UserClaims) (*pb.Tokens, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["full_name"] = user.FullName
	claims["user_type"] = user.UserType
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	cfg := config.Load()
	access, err := accessToken.SignedString([]byte(cfg.ACCESS_SIGNING_KEY))
	if err != nil {
		return nil, err
	}

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["user_id"] = user.Id
	rftclaims["username"] = user.Username
	rftclaims["email"] = user.Email
	rftclaims["full_name"] = user.FullName
	rftclaims["user_type"] = user.UserType
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	refresh, err := refreshToken.SignedString([]byte(cfg.REFRESH_SIGNING_KEY))
	if err != nil {
		return nil, err
	}

	return &pb.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func GenerateAccessToken(refresh string) (*pb.Tokens, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)

	rftclaims, err := ExtractClaims(refresh, true)
	if err != nil {
		return nil, err
	}

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = rftclaims["user_id"]
	claims["username"] = rftclaims["username"]
	claims["email"] = rftclaims["email"]
	claims["full_name"] = rftclaims["full_name"]
	claims["user_type"] = rftclaims["user_type"]
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	cfg := config.Load()
	access, err := accessToken.SignedString([]byte(cfg.ACCESS_SIGNING_KEY))
	if err != nil {
		return nil, err
	}

	return &pb.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func ExtractClaims(tokenStr string, isRefresh bool) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		if isRefresh{
			return []byte(config.Load().REFRESH_SIGNING_KEY), nil
		}
		return []byte(config.Load().ACCESS_SIGNING_KEY), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
