package middleware

import (
	"log"
	"projek_fisioterapi/configs"

	"github.com/golang-jwt/jwt"
)

func GetToken(id uint, name string, role int, level int) (string, error) {

	log.Println(id, name)
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["name"] = name
	claims["role"] = role
	claims["level"] = level

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Cfg.TokenSecret))
}
