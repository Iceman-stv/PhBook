package database

import (
	"PhBook/domen"
	"PhBook/logger"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite" // Драйвер SQLite
)

// SQLiteDB реализует интерфейс Database
type SQLiteDB struct {
	db     *sql.DB
	logger logger.Logger
}

// NewSQLiteDB создает новое подключение к SQLite и применяет миграции
func NewSQLiteDB(logger logger.Logger) (*SQLiteDB, error) {
	// Создание БД
	db, err := sql.Open("sqlite", "phonebook.db")
	if err != nil {

		logger.LogError("Ошибка при создании/открывании БД: %v", err)
		return nil, fmt.Errorf("Ошибка при создании/открывании БД: %v", err)
	}

	// Проверка соединения с базой данных
	if err := db.Ping(); err != nil {

		logger.LogError("Ошибка при подключении к БД: %v", err)
		return nil, fmt.Errorf("Ошибка при подключении к БД: %v", err)
	}

	// Применение миграций
	if err := runMigrations(db, logger); err != nil {

		logger.LogError("Ошибка при применении миграций: %v", err)
		return nil, fmt.Errorf("Ошибка при применении миграций: %v", err)
	}

	return &SQLiteDB{db: db, logger: logger}, nil
}

// runMigrations применяет миграции к базе данных
func runMigrations(db *sql.DB, logger logger.Logger) error {
	// Создание экземпляра драйвера для SQLite
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {

		logger.LogError("Ошибка при создании драйвера SQLite: %v", err)
		return fmt.Errorf("Ошибка при создании драйвера SQLite: %v", err)
	}

	// Получение абсолютного пути к папке с миграциями
	absPath, err := filepath.Abs("PhBook/migrations")
	if err != nil {
		logger.LogError("Ошибка при получении абсолютного пути: %v", err)
		return fmt.Errorf("Ошибка при получении абсолютного пути: %v", err)
	}

	absPath = filepath.ToSlash(absPath)

	// Проверка существования папки с миграциями
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		logger.LogError("Папка с миграциями не существует: %s", absPath)
		return fmt.Errorf("Папка с миграций не существует: %s", absPath)
	}

	logger.LogInfo("Папка с миграциями найдена: %s", absPath)

	// Создание экземпляра мигратора
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath, // Абсолютный путь к папке с миграциями
		"sqlite3",         // Тип базы данных
		driver,
	)
	if err != nil {
		logger.LogError("Ошибка при создании миграции: %v", err)
		return fmt.Errorf("Ошибка при создании миграции: %v", err)
	}

	// Применение миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.LogError("Ошибка при применении миграций: %v", err)
		return fmt.Errorf("Ошибка при применении миграций: %v", err)
	}

	logger.LogInfo("Миграции успешно применены!")
	
	// Проверка текущей версии миграций
	version, dirty, err := m.Version()
	if err != nil {
		
		logger.LogError("Ошибка при получении версии миграций: %v", err)
		return fmt.Errorf("Ошибка при получении версии миграций: %v", err)
	}
	logger.LogInfo("Текущая версия миграций: %v, %v", version, dirty)
	
	return nil
}

// Регистрация нового пользователя
func (s *SQLiteDB) RegisterUser(username, password string) error {
	_, err := s.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {

		s.logger.LogError("Ошибка при регистрации пользователя: %v", err)
		return fmt.Errorf("Ошибка при регистрации пользователя")
	}
	return nil
}

// Аутентификация пользователя
func (s *SQLiteDB) AuthUser(username, password string) (int, error) {
	var id int
	var storedPassword string
	err := s.db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&id, &storedPassword)
	if err != nil {

		s.logger.LogError("Ошибка при аутентификации пользователя: %v", err)
		return 0, fmt.Errorf("Ошибка при аутентификации: %v", err)
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

		s.logger.LogError("Ошибка при создании контакта: %v", err)
		return fmt.Errorf("Ошибка при создании контакта")
	}
	return nil
}

// Выводит все контакты пользователя
func (s *SQLiteDB) GetContacts(userID int) ([]domen.Contact, error) {
	rows, err := s.db.Query("SELECT id, name, phone FROM contacts WHERE user_id = ?", userID)
	if err != nil {

		s.logger.LogError("Ошибка при получении контактов: %v", err)
		return nil, fmt.Errorf("Ошибка при получении контактов")
	}
	defer rows.Close()

	var contacts []domen.Contact
	for rows.Next() {
		var contact domen.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {

			s.logger.LogError("Ошибка при сканировании контакта (GetContacts): %v", err)
			return nil, fmt.Errorf("Ошибка при сканировании контакта")
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// Поиск контакта по имени
func (s *SQLiteDB) FindContact(userID int, name string) ([]domen.Contact, error) {
	rows, err := s.db.Query("SELECT id, name, phone FROM contacts WHERE user_id = ? AND name LIKE ?", userID, "%"+name+"%")
	if err != nil {

		s.logger.LogError("Ошибка при поиске контакта: %v", err)
		return nil, fmt.Errorf("Ошибка при поиске контакта")
	}
	defer rows.Close()

	var contacts []domen.Contact

	for rows.Next() {
		var contact domen.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {

			s.logger.LogError("Ошибка при сканировании контакта (FindContact): %v", err)
			return nil, fmt.Errorf("Ошибка при сканировании контакта")
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// Удаление контакта
func (s *SQLiteDB) DelContact(userID int, name string) error {
	_, err := s.db.Exec("DELETE FROM contacts WHERE name=? AND user_id=?", name, userID)
	if err != nil {

		s.logger.LogError("Ошибка при удалении контакта: %v", err)
		return fmt.Errorf("Ошибка при удалении контакта")
	}
	return nil
}
