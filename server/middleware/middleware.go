package middleware

import (
	"context"
	"net/http"
	"strings"

	"PhBook/logger"
	"PhBook/server/utils"
)

func AuthMiddleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {

				logger.LogError("Отсутствует заголовок Authorization")
				http.Error(w, "Ошибка в заголовке запроса", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {

				logger.LogError("Неправильный формат токена")
				http.Error(w, "Неправильный формат токена", http.StatusUnauthorized)
				return
			}

			claims, err := utils.ValidateJWT(tokenString)
			if err != nil {

				logger.LogError("Неправильный токен: %v", err)
				http.Error(w, "Неправильный токен", http.StatusUnauthorized)
				return
			}

			// Добавление userID в контекст запроса
			ctx := context.WithValue(r.Context(), "userID", claims.UserID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
