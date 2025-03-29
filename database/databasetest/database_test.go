package databasetest

import (
	"PhBook/database"
	"PhBook/logger"
	"database/sql"
	"testing"
)

// setupTestDB создает тестовую БД и возвращает функцию для очистки
func setupTestDB(t *testing.T) (*database.SQLiteDB, *MockLogger, func()) {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Ошибка создания тестовой БД: %v", err)
	}

	// Создание MockLogger
	mockLogger := NewMockLogger()

	// Преобразование в интерфейс logger.Logger
	var logger logger.Logger = mockLogger

	// Применение миграции
	if err := database.RunMigrations(db, logger); err != nil {
		db.Close()
		t.Fatalf("Ошибка применения миграций: %v", err)
	}

	// Создание SQLiteDB через конструктор
	sqliteDB := database.NewTestDB(db, logger)

	// Возврат MockLogger для проверки логов
	return sqliteDB, mockLogger, func() {
		if err := db.Close(); err != nil {
			t.Logf("Ошибка при закрытии БД: %v", err)
		}
	}
}

func TestRegisterUser(t *testing.T) {
	db, mockLogger, cleanup := setupTestDB(t)
	defer cleanup()

	tests := []struct {
		name        string
		username    string
		password    string
		wantErr     bool
		expectLog   string
		expectError string
	}{
		{
			name:     "Valid",
			username: "validuser",
			password: "validpass123",
			wantErr:  false,
		},
		{
			name:        "Empty username",
			username:    "",
			password:    "validpass123",
			wantErr:     true,
			expectLog:   "имя пользователя не может быть пустым",
			expectError: "имя пользователя не может быть пустым",
		},
		{
			name:        "Empty password",
			username:    "validuser",
			password:    "",
			wantErr:     true,
			expectLog:   "пароль не может быть пустым",
			expectError: "пароль не может быть пустым",
		},
		{
			name:        "Whitespace username",
			username:    "   ",
			password:    "validpass123",
			wantErr:     true,
			expectLog:   "имя пользователя не может быть пустым",
			expectError: "имя пользователя не может быть пустым",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Сброс логов перед каждым тестом
			mockLogger.ClearLogs()

			err := db.RegisterUser(tt.username, tt.password)

			// Проверка ошибок
			if (err != nil) != tt.wantErr {

				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {

				// Проверка сообщения об ошибке
				if err == nil || err.Error() != tt.expectError {

					t.Errorf("Ожидалась ошибка '%s', получено: '%v'", tt.expectError, err)
				}

				// Проверка логов
				if !mockLogger.Contains(tt.expectLog) {

					t.Errorf("Ожидалось сообщение '%s' в логах", tt.expectLog)
					t.Logf("Полученные логи: %v", mockLogger.GetLogs())
				}
			} else if err != nil {

				t.Errorf("Не ожидалась ошибка, но получено: %v", err)
			}
		})
	}
}

func TestAuthUser(t *testing.T) {
	db, _, cleanup := setupTestDB(t)
	defer cleanup()

	// Регистрация тестового пользователя
	if err := db.RegisterUser("authuser", "authpass"); err != nil {

		t.Fatal(err)
	}

	tests := []struct {
		name     string
		username string
		password string
		wantErr  bool
	}{
		{"Valid", "authuser", "authpass", false},
		{"Wrong password", "authuser", "wrong", true},
		{"Nonexistent user", "nouser", "nopass", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.AuthUser(tt.username, tt.password)
			if (err != nil) != tt.wantErr {

				t.Errorf("Ошибка функции AuthUser() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContactOperations(t *testing.T) {
	db, mockLogger, cleanup := setupTestDB(t)
	defer cleanup()

	userID := setupTestUser(t, db, "contactuser", "contactpass")

	t.Run("Add contact", func(t *testing.T) {
		err := db.AddContact(userID, "John Doe", "+1234567890")
		if err != nil {

			t.Fatal(err)
		}

		if !mockLogger.Contains("Миграции успешно применены") {

			t.Error("Ожидалась запись об успешном применении миграций")
		}
	})

	t.Run("Get contacts", func(t *testing.T) {
		contacts, err := db.GetContacts(userID)
		if err != nil {

			t.Fatal(err)
		}

		if len(contacts) != 1 {

			t.Fatalf("Ожидался 1 контакт, получено %d", len(contacts))
		}

		if contacts[0].Name != "John Doe" {

			t.Errorf("Ожидался контакт 'John Doe', получен '%s'", contacts[0].Name)
		}
	})

	t.Run("Find contact", func(t *testing.T) {
		contacts, err := db.FindContact(userID, "Doe")
		if err != nil {

			t.Fatal(err)
		}

		if len(contacts) != 1 {

			t.Fatalf("Ожидался 1 контакт, получено %d", len(contacts))
		}
	})

	t.Run("Delete contact", func(t *testing.T) {
		err := db.DelContact(userID, "John Doe")
		if err != nil {

			t.Fatal(err)
		}

		contacts, err := db.GetContacts(userID)
		if err != nil {

			t.Fatal(err)
		}

		if len(contacts) != 0 {

			t.Fatalf("Ожидалось 0 контактов, получено %d", len(contacts))
		}
	})
}

func TestErrorCases(t *testing.T) {
	db, mockLogger, cleanup := setupTestDB(t)
	defer cleanup()

	t.Run("Add contact to nonexistent user", func(t *testing.T) {
		err := db.AddContact(999, "Test", "+123")
		if err == nil {

			t.Error("Ожидалась ошибка о несуществующем пользователе")
		} else if err.Error() != "пользователь не существует" {

			t.Errorf("Ожидалась ошибка 'пользователь не существует', получено: %v", err)
		}

		if !mockLogger.Contains("Пользователь с ID 999 не существует") {

			t.Error("Ожидалось сообщение об ошибке в логах")
		}
	})

	t.Run("Get contacts for nonexistent user", func(t *testing.T) {
		_, err := db.GetContacts(999)
		if err == nil {

			t.Error("Ожидалась ошибка о несуществующем пользователе")
		} else if err.Error() != "пользователь не существует" {

			t.Errorf("Ожидалась ошибка 'пользователь не существует', получено: %v", err)
		}

		if !mockLogger.Contains("Пользователь с ID 999 не существует") {

			t.Error("Ожидалось сообщение об ошибке в логах")
		}
	})
}

func setupTestUser(t *testing.T, db *database.SQLiteDB, username, password string) int {
	t.Helper()

	if err := db.RegisterUser(username, password); err != nil {

		t.Fatal(err)
	}

	userID, err := db.AuthUser(username, password)
	if err != nil {

		t.Fatal(err)
	}

	return userID
}
