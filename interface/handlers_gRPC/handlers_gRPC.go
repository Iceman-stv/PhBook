package handlers_gRPC

import (
    "context"
    "PhBook/domen"
    pB "PhBook/gen/github.com/Iceman-stv/PhBook/gen"
    "PhBook/logger"
    "PhBook/server/jwt"
    "PhBook/userCase"
)

// PhoneBookHandlers содержит методы для обработки запросов gRPC
type PhoneBookHandlers struct {
    PhoneBook *userCase.PhoneBook
    Logger    logger.Logger
}

// NewPhoneBookHandlers создаёт новый экземпляр PhoneBookHandlers
func NewPhoneBookHandlers(pb *userCase.PhoneBook, l logger.Logger) *PhoneBookHandlers {
    return &PhoneBookHandlers{
        PhoneBook: pb,
        Logger:    l,
    }
}

// AuthUser аутентифицирует пользователя и возвращает JWT
func (h *PhoneBookHandlers) AuthUser(ctx context.Context, req *pB.AuthUserRequest) (*pB.AuthUserResponse, error) {
    userID, err := h.PhoneBook.AuthUser(req.Username, req.Password)
    if err != nil {
        h.Logger.LogError("Ошибка при аутентификации пользователя: %v", err)
        return nil, err
    }

    // Генерация JWT
    token, err := jwt.GenerateJWT(userID)
    if err != nil {
        h.Logger.LogError("Ошибка при генерации JWT: %v", err)
        return nil, err
    }

    return &pB.AuthUserResponse{
        UserId: int32(userID),
        Token:  token,
    }, nil
}

// AddContact добавляет контакт
func (h *PhoneBookHandlers) AddContact(ctx context.Context, req *pB.AddContactRequest) (*pB.AddContactResponse, error) {
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