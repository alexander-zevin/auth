package token

import (
	"github.com/dgrijalva/jwt-go"
)

func DecodeToken(value string) (string, string) {

	// Разбираем токен
	token, _ := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	login := claims["login"].(string)
	password := claims["password"].(string)
	return login, password
}