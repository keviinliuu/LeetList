package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/keviinliuu/leetlist/graph/model"
)

func GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.Email 
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}