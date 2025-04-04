package database

import (
	"PhBook/domen"
	"PhBook/logger"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

type SQLiteDB struct {
	db     *sql.DB
	logger logger.Logger
}

func NewSQLiteDB(logger logger.Logger) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite", "phonebook.db")
	if err != nil {

		logger.LogError("Failed to open DB: %v", err)
		return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	if err := db.Ping(); err != nil {

		logger.LogError("Failed to ping DB: %v", err)
		return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	if err := RunMigrations(db, logger); err != nil {

		return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	return &SQLiteDB{db: db, logger: logger}, nil
}

func NewTestDB(db *sql.DB, logger logger.Logger) *SQLiteDB {
	return &SQLiteDB{db: db, logger: logger}
}

func RunMigrations(db *sql.DB, logger logger.Logger) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {

		logger.LogError("Failed to create SQLite driver: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	absPath, err := filepath.Abs("migrations")
	if err != nil {

		logger.LogError("Failed to get absolute path: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	absPath = filepath.ToSlash(absPath)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {

		logger.LogError("Migrations directory does not exist: %s", absPath)
		return fmt.Errorf("%w: migrations directory missing", domen.ErrOperationFailed)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"sqlite3",
		driver,
	)
	if err != nil {

		logger.LogError("Failed to create migrator: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {

		logger.LogError("Failed to apply migrations: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	logger.LogInfo("Миграции применены успешно")
	return nil
}

func (s *SQLiteDB) RegisterUser(username, password string) error {
	if strings.TrimSpace(username) == "" {

		s.logger.LogError(domen.ErrEmptyUsername.Error())
		return domen.ErrEmptyUsername
	}

	if strings.TrimSpace(password) == "" {

		s.logger.LogError(domen.ErrEmptyPassword.Error())
		return domen.ErrEmptyPassword
	}

	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
	if err != nil {

		s.logger.LogError("User check error: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	if exists {

		s.logger.LogError(domen.ErrUserExists.Error())
		return domen.ErrUserExists
	}

	_, err = s.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {

		s.logger.LogError("Registration error: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	return nil
}

func (s *SQLiteDB) AuthUser(username, password string) (int, error) {
	var id int
	var storedPassword string

	err := s.db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&id, &storedPassword)
	if err != nil {

		if err == sql.ErrNoRows {

			s.logger.LogError(domen.ErrUserNotFound.Error())
			return 0, domen.ErrUserNotFound
		}
		s.logger.LogError("Authentication error: %v", err)
		return 0, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}

	if storedPassword != password {

		s.logger.LogError(domen.ErrInvalidCredentials.Error())
		return 0, domen.ErrInvalidCredentials
	}

	return id, nil
}

func (s *SQLiteDB) AddContact(userID int, name, phone string) error {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", userID).Scan(&exists)
	if err != nil {

		s.logger.LogError("User check error: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}
	if !exists {

		s.logger.LogError(domen.ErrUserNotFound.Error())
		return domen.ErrUserNotFound
	}

	_, err = s.db.Exec("INSERT INTO contacts (name, phone, user_id) VALUES (?, ?, ?)", name, phone, userID)
	if err != nil {

		s.logger.LogError("Failed to add contact: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}
	return nil
}

func (s *SQLiteDB) GetContacts(userID int) ([]domen.Contact, error) {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", userID).Scan(&exists)
	if err != nil {

		s.logger.LogError("User check error: %v", err)
		return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}
	if !exists {

		s.logger.LogError(domen.ErrUserNotFound.Error())
		return nil, domen.ErrUserNotFound
	}

	rows, err := s.db.Query("SELECT id, name, phone FROM contacts WHERE user_id = ?", userID)
	if err != nil {

		s.logger.LogError("Failed to get contacts: %v", err)
		return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}
	defer rows.Close()

	var contacts []domen.Contact
	for rows.Next() {
		var contact domen.Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {

			s.logger.LogError("Contact scan error: %v", err)
			return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (s *SQLiteDB) FindContact(userID int, name string) ([]domen.Contact, error) {
	rows, err := s.db.Query("SELECT id, name, phone FROM contacts WHERE user_id = ? AND name LIKE ?", userID, "%"+name+"%")
	if err != nil {

		s.logger.LogError("Contact search error: %v", err)
		return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}
	defer rows.Close()

	var contacts []domen.Contact
	for rows.Next() {
		var contact domen.Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {

			s.logger.LogError("Contact scan error: %v", err)
			return nil, fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
		}
		contacts = append(contacts, contact)
	}

	if len(contacts) == 0 {

		return nil, domen.ErrContactNotFound
	}

	return contacts, nil
}

func (s *SQLiteDB) DelContact(userID int, name string) error {
	_, err := s.db.Exec("DELETE FROM contacts WHERE name=? AND user_id=?", name, userID)
	if err != nil {

		s.logger.LogError("Failed to delete contact: %v", err)
		return fmt.Errorf("%w: %v", domen.ErrOperationFailed, err)
	}
	return nil
}
