package userCase

import (
	"PhBook/domen"
)

// Интерфейс для работы с БД
type Database interface {
	RegisterUser(username, password string) error
	AuthUser(username, password string) (int, error)
	AddContact(userID int, name, phone string) error
	DelContact(userID int, name string) error
	FindContact(userID int, name string) ([]domen.Contact, error)
	GetContacts(userID int) ([]domen.Contact, error)
}

// Реализация бизнес логики
type PhoneBook struct {
	db Database
}

// Создание нового экземпляра телефонной книги
func NewPhoneBook(db Database) *PhoneBook {
	return &PhoneBook{db: db}
}

// Регистрация
func (pb *PhoneBook) RegisterUser(username, password string) error {
	return pb.db.RegisterUser(username, password)
}

// Вход
func (pb *PhoneBook) AuthUser(username, password string) (int, error) {
	return pb.db.AuthUser(username, password)
}

// Создание контакта
func (pb *PhoneBook) AddContact(userID int, name, phone string) error {
	return pb.db.AddContact(userID, name, phone)
}

// Удаление контакта
func (pb *PhoneBook) DelContact(userID int, name string) error {
	return pb.db.DelContact(userID, name)
}

// Поиск контакта
func (pb *PhoneBook) FindContact(userID int, name string) ([]domen.Contact, error) {
	return pb.db.FindContact(userID, name)
}

// Список всех контактов
func (pb *PhoneBook) GetContacts(userID int) ([]domen.Contact, error) {
	return pb.db.GetContacts(userID)
}
