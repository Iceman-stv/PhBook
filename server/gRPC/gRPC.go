package gRPC

import (
	"PhBook/interface/interfaceGrpc"
	"PhBook/logger"
	"PhBook/server/middlewareGrpc"
	"PhBook/userCase"
	"context"
	"fmt"
	"net"

	pB "PhBook/gen/github.com/Iceman-stv/PhBook/gen"

	"google.golang.org/grpc"
)

// GRPCServer представляет gRPC-сервер
type GRPCServer struct {
	pB.UnimplementedPhoneBookServiceServer
	handlers *interfaceGrpc.PhoneBookHandlers
	logger   logger.Logger
}

// NewGRPCServer создаёт новый экземпляр GRPCServer
func NewGRPCServer(pb *userCase.PhoneBook, l logger.Logger) *GRPCServer {
	return &GRPCServer{
		handlers: interfaceGrpc.NewPhoneBookHandlers(pb, l),
		logger:   l,
	}
}

// Start запускает gRPC-сервер
func (s *GRPCServer) Start() error {
	// Создание gRPC-сервера
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middlewareGrpc.AuthInterceptor(s.logger)),
	)

	// Регистрация сервиса на gRPC-сервере
	pB.RegisterPhoneBookServiceServer(grpcServer, s)

	// Запуск сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		s.logger.LogError("Ошибка при запуске gRPC-сервера: %v", err)
		return err
	}

	s.logger.LogInfo("gRPC-сервер запущен на %v", lis.Addr())
	mes := fmt.Sprintf("gRPC-сервер запущен на %v", lis.Addr())
	fmt.Println(mes)
	if err := grpcServer.Serve(lis); err != nil {
		s.logger.LogError("Ошибка при работе gRPC-сервера: %v", err)
		return err
	}

	return nil
}

// AuthUser аутентифицирует пользователя и возвращает JWT
func (s *GRPCServer) AuthUser(ctx context.Context, req *pB.AuthUserRequest) (*pB.AuthUserResponse, error) {
	return s.handlers.AuthUser(ctx, req)
}

// AddContact добавляет контакт
func (s *GRPCServer) AddContact(ctx context.Context, req *pB.AddContactRequest) (*pB.AddContactResponse, error) {
	return s.handlers.AddContact(ctx, req)
}

// DelContact удаляет контакт
func (s *GRPCServer) DelContact(ctx context.Context, req *pB.DelContactRequest) (*pB.DelContactResponse, error) {
	return s.handlers.DelContact(ctx, req)
}

// FindContact ищет контакт
func (s *GRPCServer) FindContact(ctx context.Context, req *pB.FindContactRequest) (*pB.FindContactResponse, error) {
	return s.handlers.FindContact(ctx, req)
}

// GetContacts возвращает все контакты
func (s *GRPCServer) GetContacts(ctx context.Context, req *pB.GetContactsRequest) (*pB.GetContactsResponse, error) {
	return s.handlers.GetContacts(ctx, req)
}
