package domen

import "errors"

// Общие ошибки приложения
var (
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmptyUsername      = errors.New("username cannot be empty")
	ErrEmptyPassword      = errors.New("password cannot be empty")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrContactNotFound    = errors.New("contact not found")
	ErrContactExists      = errors.New("contact already exists")
	ErrOperationFailed    = errors.New("operation failed")
)

// User представляет структуру пользователя
type User struct {
	ID       int    `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Contact представляет структуру контакта
type Contact struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID int    `json:"userID"`
}
