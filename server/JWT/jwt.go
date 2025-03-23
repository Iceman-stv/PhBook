package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Загрузка переменных окружения из .env файла
func init() {
	if err := godotenv.Load(); err != nil {
		panic("Ошибка загрузки .env файла")
	}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtSecret, nil
	})
	if err != nil {

		return nil, err
	}
	if !token.Valid {

		return nil, jwt.ErrSignatureInvalid
	}
	if time.Now().After(claims.ExpiresAt.Time) {

		return nil, err
	}
	return claims, nil
}
