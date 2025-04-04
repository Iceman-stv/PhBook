package interfaceGrpc

import (
	"PhBook/domen"
	pB "PhBook/gen/github.com/Iceman-stv/PhBook/gen"
	"PhBook/logger"
	"PhBook/server/jwt"
	"PhBook/userCase"
	"context"
)

// JWTGenerator определяет интерфейс для генерации JWT
type JWTGenerator interface {
	GenerateJWT(userID int) (string, error)
}

// Реальная реализация JWTGenerator
type defaultJWTGenerator struct{}

func (g *defaultJWTGenerator) GenerateJWT(userID int) (string, error) {
	return jwt.GenerateJWT(userID)
}

// PhoneBookHandlers содержит методы для обработки запросов gRPC
type PhoneBookHandlers struct {
	PhoneBook userCase.Database
	Logger    logger.Logger
	jwtGen    JWTGenerator
}

// NewPhoneBookHandlers создаёт production-экземпляр
func NewPhoneBookHandlers(pb userCase.Database, l logger.Logger) *PhoneBookHandlers {
	return &PhoneBookHandlers{
		PhoneBook: pb,
		Logger:    l,
		jwtGen:    &defaultJWTGenerator{},
	}
}

// NewTestPhoneBookHandlers создаёт экземпляр для тестов
func NewTestPhoneBookHandlers(pb userCase.Database, l logger.Logger, jwtGen JWTGenerator) *PhoneBookHandlers {
	return &PhoneBookHandlers{
		PhoneBook: pb,
		Logger:    l,
		jwtGen:    jwtGen,
	}
}

// AuthUser обрабатывает аутентификацию
func (h *PhoneBookHandlers) AuthUser(ctx context.Context, req *pB.AuthUserRequest) (*pB.AuthUserResponse, error) {
	if req.Username == "" || req.Password == "" {

		h.Logger.LogError("Ошибка аутентификации: %v", domen.ErrInvalidCredentials)
		return nil, domen.ErrInvalidCredentials
	}

	userID, err := h.PhoneBook.AuthUser(req.Username, req.Password)
	if err != nil {

		h.Logger.LogError("Ошибка аутентификации: %v", err)
		return nil, err
	}

	token, err := h.jwtGen.GenerateJWT(userID)
	if err != nil {

		h.Logger.LogError("Ошибка генерации jwt: %v", err)
		return nil, domen.ErrOperationFailed
	}

	return &pB.AuthUserResponse{
		UserId: int32(userID),
		Token:  token,
	}, nil
}

// RegisterUser обрабатывает регистрацию
func (h *PhoneBookHandlers) RegisterUser(ctx context.Context, req *pB.RegisterUserRequest) (*pB.RegisterUserResponse, error) {
	if req.Username == "" {

		h.Logger.LogError("Ошибка регистрации: %v", domen.ErrEmptyUsername)
		return nil, domen.ErrEmptyUsername
	}
	if req.Password == "" {

		h.Logger.LogError("Ошибка регистрации: %v", domen.ErrEmptyPassword)
		return nil, domen.ErrEmptyPassword
	}

	err := h.PhoneBook.RegisterUser(req.Username, req.Password)
	if err != nil {

		h.Logger.LogError("Ошибка регистрации: %v", err)
		return nil, err
	}

	return &pB.RegisterUserResponse{}, nil
}

// AddContact добавляет контакт
func (h *PhoneBookHandlers) AddContact(ctx context.Context, req *pB.AddContactRequest) (*pB.AddContactResponse, error) {
	if req.Name == "" || req.Phone == "" {

		h.Logger.LogError("Ошибка при добавлении контакта: %v", domen.ErrOperationFailed)
		return nil, domen.ErrOperationFailed
	}

	err := h.PhoneBook.AddContact(int(req.UserId), req.Name, req.Phone)
	if err != nil {

		h.Logger.LogError("Ошибка при добавлении контакта: %v", err)
		return nil, err
	}

	return &pB.AddContactResponse{}, nil
}

// DelContact удаляет контакт
func (h *PhoneBookHandlers) DelContact(ctx context.Context, req *pB.DelContactRequest) (*pB.DelContactResponse, error) {
	err := h.PhoneBook.DelContact(int(req.UserId), req.Name)
	if err != nil {

		h.Logger.LogError("Ошибка при удалении контакта: %v", err)
		return nil, err
	}
	return &pB.DelContactResponse{}, nil
}

// FindContact ищет контакт
func (h *PhoneBookHandlers) FindContact(ctx context.Context, req *pB.FindContactRequest) (*pB.FindContactResponse, error) {
	contacts, err := h.PhoneBook.FindContact(int(req.UserId), req.Name)
	if err != nil {

		h.Logger.LogError("Ошибка при поиске контакта: %v", err)
		return nil, err
	}
	return &pB.FindContactResponse{Contacts: toProtoContacts(contacts)}, nil
}

// GetContacts возвращает все контакты
func (h *PhoneBookHandlers) GetContacts(ctx context.Context, req *pB.GetContactsRequest) (*pB.GetContactsResponse, error) {
	contacts, err := h.PhoneBook.GetContacts(int(req.UserId))
	if err != nil {

		h.Logger.LogError("Ошибка при получении контактов: %v", err)
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
