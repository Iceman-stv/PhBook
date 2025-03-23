package gRPC

import (
	"PhBook/domen"
	pB "PhBook/gen/github.com/Iceman-stv/PhBook/gen"
	"PhBook/logger"
	"PhBook/server/jwt"
	"PhBook/server/middleware_gRPC"
	"PhBook/userCase"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// GRPCServer представляет gRPC-сервер
type GRPCServer struct {
	pB.UnimplementedPhoneBookServiceServer
	phoneBook *userCase.PhoneBook
	logger    logger.Logger
	server    grpc.Server
}

// NewGRPCServer создаёт новый экземпляр GRPCServer
func NewGRPCServer(pb *userCase.PhoneBook, l logger.Logger) *GRPCServer {
	return &GRPCServer{
		phoneBook: pb,
		logger:    l,
	}
}

// Start запускает gRPC-сервер
func (s *GRPCServer) Start() error {
	// Создание gRPC-сервер
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware_gRPC.AuthInterceptor(s.logger)),
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
	fmt.Printf("gRPC-сервер запущен на %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {

		s.logger.LogError("Ошибка при работе gRPC-сервера: %v", err)
		return err
	}

	return nil
}

// AuthUser аутентифицирует пользователя и возвращает JWT
func (s *GRPCServer) AuthUser(ctx context.Context, req *pB.AuthUserRequest) (*pB.AuthUserResponse, error) {
	userID, err := s.phoneBook.AuthUser(req.Username, req.Password)
	if err != nil {

		s.logger.LogError("Ошибка при аутентификации пользователя: %v", err)
		return nil, err
	}

	// Генерация JWT
	token, err := jwt.GenerateJWT(userID)
	if err != nil {

		s.logger.LogError("Ошибка при генерации JWT: %v", err)
		return nil, err
	}

	return &pB.AuthUserResponse{
		UserId: int32(userID),
		Token:  token,
	}, nil
}

// Добавление контакта
func (s *GRPCServer) AddContact(ctx context.Context, req *pB.AddContactRequest) (*pB.AddContactResponse, error) {
	err := s.phoneBook.AddContact(int(req.UserId), req.Name, req.Phone)
	if err != nil {

		s.logger.LogError("Ошибка при добавлении контакта: %v", err)
		return nil, err
	}
	return &pB.AddContactResponse{}, nil
}

// Удаление контакта
func (s *GRPCServer) DelContact(ctx context.Context, req *pB.DelContactRequest) (*pB.DelContactResponse, error) {
	err := s.phoneBook.DelContact(int(req.UserId), req.Name)
	if err != nil {

		s.logger.LogError("Ошибка при удалении контакта: %v", err)
		return nil, err
	}
	return &pB.DelContactResponse{}, nil
}

// Поиск контакта
func (s *GRPCServer) FindContact(ctx context.Context, req *pB.FindContactRequest) (*pB.FindContactResponse, error) {
	contacts, err := s.phoneBook.FindContact(int(req.UserId), req.Name)
	if err != nil {

		s.logger.LogError("Ошибка при поиске контакта: %v", err)
		return nil, err
	}
	return &pB.FindContactResponse{Contacts: toProtoContacts(contacts)}, nil
}

// Вывод всех контактов
func (s *GRPCServer) GetContacts(ctx context.Context, req *pB.GetContactsRequest) (*pB.GetContactsResponse, error) {
	contacts, err := s.phoneBook.GetContacts(int(req.UserId))
	if err != nil {

		s.logger.LogError("Ошибка при получении контактов: %v", err)
		return nil, err
	}
	return &pB.GetContactsResponse{Contacts: toProtoContacts(contacts)}, nil
}

// toProtoContacts преобразует []domen.Contact в []*pB.Contact
func toProtoContacts(contacts []domen.Contact) []*pB.Contact {
	var protoContacts []*pB.Contact
	for _, c := range contacts {
		protoContacts = append(protoContacts, &pB.Contact{
			Name:  c.Name,
			Phone: c.Phone,
		})
	}
	return protoContacts
}
