package gRPC

import (
	"PhBook/domen"
	"PhBook/interface/interfaceGrpc"
	"PhBook/logger"
	"PhBook/server/middlewareGrpc"
	"PhBook/userCase"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pB "PhBook/gen/github.com/Iceman-stv/PhBook/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// ServerConfig содержит конфигурацию gRPC сервера
type ServerConfig struct {
	Port             string        // Порт в формате ":50051"
	EnableReflection bool          // Включить reflection API
	EnableAuth       bool          // Включить аутентификацию
	Timeout          time.Duration // Таймаут соединения
}

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		Port:             ":50051",
		EnableReflection: true,
		EnableAuth:       true,
		Timeout:          30 * time.Second,
	}
}

// GRPCServer реализует gRPC сервер телефонной книги
type GRPCServer struct {
	pB.UnimplementedPhoneBookServiceServer
	handlers *interfaceGrpc.PhoneBookHandlers
	logger   logger.Logger
	config   *ServerConfig
	server   *grpc.Server
}

// New создает новый экземпляр gRPC сервера
func New(pb *userCase.PhoneBook, l logger.Logger, cfg *ServerConfig) *GRPCServer {
	if pb == nil {

		panic("phonebook use case cannot be nil")
	}
	if l == nil {

		panic("logger cannot be nil")
	}
	if cfg == nil {

		cfg = DefaultConfig()
	}

	return &GRPCServer{
		handlers: interfaceGrpc.NewPhoneBookHandlers(pb, l),
		logger:   l,
		config:   cfg,
	}
}

// Start запускает gRPC сервер
func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {

		return fmt.Errorf("failed to listen: %w", err)
	}

	// Настройка интерцепторов
	interceptors := []grpc.UnaryServerInterceptor{
		middlewareGrpc.AuthInterceptor(s.logger),
	}
	if s.config.EnableAuth {

		interceptors = append(interceptors, middlewareGrpc.AuthInterceptor(s.logger))
	}

	s.server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors...),
		grpc.ConnectionTimeout(s.config.Timeout),
	)

	pB.RegisterPhoneBookServiceServer(s.server, s)

	if s.config.EnableReflection {

		reflection.Register(s.server)
		s.logger.LogInfo("gRPC reflection работает")
	}

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errChan := make(chan error, 1)
	go func() {
		s.logger.LogInfo("Старт gRPC сервера на %s", s.config.Port)
		fmt.Printf("Старт gRPC сервера на %s\n", s.config.Port)
		if err := s.server.Serve(lis); err != nil {

			errChan <- fmt.Errorf("Ошибка gRPC сервера: %v", err)
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		s.logger.LogInfo("сервер gRPC завершил работу успешно...")
		s.server.GracefulStop()
		return nil
	}
}

// Stop останавливает сервер
func (s *GRPCServer) Stop() {
	if s.server != nil {

		s.server.GracefulStop()
	}
}

// AuthUser обрабатывает аутентификацию пользователя
func (s *GRPCServer) AuthUser(ctx context.Context, req *pB.AuthUserRequest) (*pB.AuthUserResponse, error) {
	resp, err := s.handlers.AuthUser(ctx, req)
	if err != nil {

		return nil, s.mapError(err)
	}
	return resp, nil
}

// RegisterUser обрабатывает регистрацию пользователя
func (s *GRPCServer) RegisterUser(ctx context.Context, req *pB.RegisterUserRequest) (*pB.RegisterUserResponse, error) {
	resp, err := s.handlers.RegisterUser(ctx, req)
	if err != nil {

		return nil, s.mapError(err)
	}
	return resp, nil
}

// AddContact добавляет новый контакт
func (s *GRPCServer) AddContact(ctx context.Context, req *pB.AddContactRequest) (*pB.AddContactResponse, error) {
	resp, err := s.handlers.AddContact(ctx, req)
	if err != nil {

		return nil, s.mapError(err)
	}
	return resp, nil
}

// DelContact удаляет контакт
func (s *GRPCServer) DelContact(ctx context.Context, req *pB.DelContactRequest) (*pB.DelContactResponse, error) {
	resp, err := s.handlers.DelContact(ctx, req)
	if err != nil {

		return nil, s.mapError(err)
	}
	return resp, nil
}

// FindContact ищет контакт
func (s *GRPCServer) FindContact(ctx context.Context, req *pB.FindContactRequest) (*pB.FindContactResponse, error) {
	resp, err := s.handlers.FindContact(ctx, req)
	if err != nil {

		return nil, s.mapError(err)
	}
	return resp, nil
}

// GetContacts возвращает список контактов
func (s *GRPCServer) GetContacts(ctx context.Context, req *pB.GetContactsRequest) (*pB.GetContactsResponse, error) {
	resp, err := s.handlers.GetContacts(ctx, req)
	if err != nil {

		return nil, s.mapError(err)
	}
	return resp, nil
}

// mapError преобразует доменные ошибки в gRPC статусы
func (s *GRPCServer) mapError(err error) error {
	switch err {
	case domen.ErrInvalidCredentials, domen.ErrEmptyUsername, domen.ErrEmptyPassword:
		return status.Error(codes.InvalidArgument, err.Error())
	case domen.ErrUserExists, domen.ErrContactExists:
		return status.Error(codes.AlreadyExists, err.Error())
	case domen.ErrUserNotFound, domen.ErrContactNotFound:
		return status.Error(codes.NotFound, err.Error())
	default:
		s.logger.LogError("Unhandled error: %v", err)
		return status.Error(codes.Internal, domen.ErrOperationFailed.Error())
	}
}
