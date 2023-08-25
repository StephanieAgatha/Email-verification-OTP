package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secret = []byte("blablablaaa")

func GenerateJWT(email string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["email"] = email
	claims["iss"] = "Sora project"

	//signed token
	tokenstr, err := token.SignedString(secret)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenstr, nil
}
