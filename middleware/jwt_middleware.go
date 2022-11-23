package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(adminID uint, roleID uint, medicalID uint, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["adminID"] = adminID
	claims["roleID"] = roleID
	claims["medicalID"] = medicalID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
