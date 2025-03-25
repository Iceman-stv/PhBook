package jwt

import (
	"os"
	"time"

	"PhBook/logger"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	log       logger.Logger
	jwtSecret []byte
)

// Инициализация модуля с передачей логгера
func Init(l logger.Logger) {
	log = l

	// Загрузка .env файлов
	if err := godotenv.Load(".env.example", ".env"); err != nil {
		log.LogError("JWT: Ошибка загрузки .env файлов", err)
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.LogError("JWT: JWT_SECRET не установлен", nil)
		panic("JWT_SECRET должен быть установлен в .env файле")
	}

	jwtSecret = []byte(secret)
	log.LogInfo("JWT: Модуль инициализирован")
}

type Claims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int) (string, error) {
	startTime := time.Now()

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
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		log.LogError("JWT: Ошибка генерации токена", err,
			"userID", userID,
			"duration", time.Since(startTime))
		return "", err
	}

	log.LogInfo("JWT: Токен успешно сгенерирован",
		"userID", userID,
		"expiresAt", expirationTime,
		"duration", time.Since(startTime))

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	startTime := time.Now()

	if tokenString == "" {
		log.LogWarn("JWT: Пустой токен")
		return nil, jwt.ErrTokenMalformed
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		log.LogError("JWT: Ошибка валидации токена", err,
			"token", maskToken(tokenString),
			"duration", time.Since(startTime))
		return nil, err
	}

	if !token.Valid {
		log.LogWarn("JWT: Невалидный токен",
			"token", maskToken(tokenString),
			"reason", getInvalidReason(token))
		return nil, jwt.ErrTokenUnverifiable
	}

	log.LogInfo("JWT: Токен валиден",
		"userID", claims.UserID,
		"expiresAt", claims.ExpiresAt,
		"duration", time.Since(startTime))

	return claims, nil
}

// Вспомогательные функции для логирования

func maskToken(token string) string {
	if len(token) < 10 {
		return "*****"
	}
	return token[:5] + "****" + token[len(token)-5:]
}

func getInvalidReason(token *jwt.Token) string {
	if token.Claims.(*Claims).ExpiresAt.Before(time.Now()) {
		return "token expired"
	}
	return "invalid signature"
}