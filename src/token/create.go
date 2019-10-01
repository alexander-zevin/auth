package token

import (
	"github.com/dgrijalva/jwt-go"
	"log"
)

// Секретный ключ
var mySigningKey = []byte("secret")

func CreateToken(login, password string) string {

	// Создаем новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":    login,
		"password": password,
	})
	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
