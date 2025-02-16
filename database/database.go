package database

import (
	"PhBook/domen"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" //Драйвер SQLite
)

// Реализация интерфейса Database
type SQLiteDB struct {
	db *sql.DB
}

// Создание нового подключения к SQLite
func NewSQLiteDB() (*SQLiteDB, error) {
	db, err := sql.Open("sqlite", "phonebook.db")

	if err != nil {

		return nil, fmt.Errorf("Ошибка при открытии БД %v", err)
	}

	//Создание таблицы пользователей, если она не существует
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
	)
	`)

	if err != nil {

		return nil, fmt.Errorf("Ошибка при создании таблицы пользователей %v", err)
	}

	//Создание таблицы контактов, если она не существует
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS contacts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	phone TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id)
	)
	`)

	if err != nil {

		return nil, fmt.Errorf("Ошибка при создании таблицы контактов %v", err)
	}

	return &SQLiteDB{db: db}, nil
}

// Регистрация нового пользователя
func (s *SQLiteDB) RegisterUser(username, password string) error {
	_, err := s.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)

	if err != nil {

		return fmt.Errorf("Ошибка при регистрации пользователя %v", err)
	}

	return nil
}

// Аутентификация пользователя
func (s *SQLiteDB) AuthUser(username, password string) (int, error) {
	var id int
	var storedPassword string
	err := s.db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&id, &storedPassword)

	if err != nil {

		return 0, fmt.Errorf("Ошибка при аутентификации %v", err)
	}

	if storedPassword != password {

		return 0, fmt.Errorf("Неверный пароль")
	}

	return id, nil
}

// Создание нового контакта
func (s *SQLiteDB) AddContact(userID int, name, phone string) error {
	_, err := s.db.Exec("INSERT INTO contacts (name, phone, user_id) VALUES (?, ?, ?)", name, phone, userID)

	if err != nil {

		return fmt.Errorf("Ошибка при создании контакта %v", err)
	}

	return nil
}

// Выводит все контакты пользователя
func (s *SQLiteDB) GetContacts(userID int) ([]domen.Contact, error) {
	rows, err := s.db.Query("SELECT id, name, phone FROM contacts WHERE user_id = ?", userID)

	if err != nil {

		return nil, fmt.Errorf("ошибка при получении контактов %v", err)
	}
	defer rows.Close()

	var contacts []domen.Contact

	for rows.Next() {
		var contact domen.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)

		if err != nil {

			return nil, fmt.Errorf("ошибка при сканировании контакта %v", err)
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// Поиск контакта по имени
func (s *SQLiteDB) FindContact(userID int, name string) ([]domen.Contact, error) {
	rows, err := s.db.Query("SELECT id, name, phone FROM contacts WHERE user_id = ? AND name LIKE ?", userID, "%"+name+"%")

	if err != nil {

		return nil, fmt.Errorf("ошибка при поиске контакта %v", err)
	}
	defer rows.Close()

	var contacts []domen.Contact

	for rows.Next() {
		var contact domen.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)

		if err != nil {

			return nil, fmt.Errorf("ошибка при сканировании контакта %v", err)
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// Удаление контакта
func (s *SQLiteDB) DelContact(userID int, name string) error {
	_, err := s.db.Exec("DELETE FROM contacts WHERE name=? AND user_id=?", name, userID)

	if err != nil {

		return fmt.Errorf("Ошибка при удалении контакта %v", err)
	}

	return nil
}
