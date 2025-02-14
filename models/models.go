package models

// User представляет структуру пользователя
type User struct {
	ID       int
	Username string
	Password string
}

// Contact представляет структуру контакта
type Contact struct {
	ID     int
	Name   string
	Phone  string
	UserID int
}

// Интерфейс для работы с БД
type Database interface {
	RegisterUser(username, password string) error
	AuthUser(username, password string) (int, error)
	AddContact(userID int, name, phone string) error
	DelContact(userID int, name string) error
	FindContact(userID int, name string) ([]Contact, error)
	GetContacts(userID int) ([]Contact, error)
}
