package middlewareGrpc

import (
	"context"
	"strings"

	"PhBook/logger"
	"PhBook/server/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor создаёт gRPC-интерцептор для проверки JWT
func AuthInterceptor(logger logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Игнорирование проверки JWT для метода AuthUser
		if info.FullMethod == "/phonebook.PhoneBookService/AuthUser" {

			return handler(ctx, req)
		}

		// Получение метаданных из контекста
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {

			logger.LogError("Метаданные не предоставлены")
			return nil, status.Error(codes.Unauthenticated, "метаданные не предоставлены")
		}

		// Извлечение токена из заголовка
		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {

			logger.LogError("Токен не предоставлен")
			return nil, status.Error(codes.Unauthenticated, "токен не предоставлен")
		}

		// Удаление префикса "Bearer "
		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
		if tokenString == "" {

			logger.LogError("Неправильный формат токена")
			return nil, status.Error(codes.Unauthenticated, "неправильный формат токена")
		}

		// Валидирование токена
		claims, err := jwt.ValidateJWT(tokenString)
		if err != nil {

			logger.LogError("Неправильный токен: %v", err)
			return nil, status.Error(codes.Unauthenticated, "неправильный токен")
		}

		// Добавление userID в контекст
		ctx = context.WithValue(ctx, "userID", claims.UserID)

		// Продолжение выполнения запроса
		return handler(ctx, req)
	}
}
