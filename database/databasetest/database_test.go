package databasetest

import (
	"PhBook/database"
	"PhBook/domen"
	"PhBook/logger"
	"PhBook/logger/mocklog"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*database.SQLiteDB, *mocklog.MockLogger, func()) {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {

		t.Fatalf("Failed to create test DB: %v", err)
	}

	mockLogger := mocklog.NewMockLogger()
	var logger logger.Logger = mockLogger

	if err := database.RunMigrations(db, logger); err != nil {

		db.Close()
		t.Fatalf("Failed to apply migrations: %v", err)
	}

	sqliteDB := database.NewTestDB(db, logger)
	return sqliteDB, mockLogger, func() {
		if err := db.Close(); err != nil {

			t.Logf("Failed to close DB: %v", err)
		}
	}
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

func TestDatabase(t *testing.T) {
	db, mockLogger, cleanup := setupTestDB(t)
	defer cleanup()

	t.Run("UserOperations", func(t *testing.T) {
		t.Run("RegisterUser", func(t *testing.T) {
			tests := []struct {
				name     string
				username string
				password string
				wantErr  error
			}{
				{"Success", domen.TestUsername, domen.TestPassword, nil},
				{"Empty username", "", domen.TestPassword, domen.ErrEmptyUsername},
				{"Empty password", domen.TestUsername, "", domen.ErrEmptyPassword},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					mockLogger.ClearLogs()
					err := db.RegisterUser(tt.username, tt.password)

					if tt.wantErr != nil {

						assert.ErrorIs(t, err, tt.wantErr)
						assert.True(t, mockLogger.Contains(tt.wantErr.Error()),
							"Expected log to contain '%s', got logs: %v",
							tt.wantErr.Error(), mockLogger.GetLogs())
					} else {
						assert.NoError(t, err)
					}
				})
			}
		})

		t.Run("AuthUser", func(t *testing.T) {
			authUsername := "authuser"
			authPassword := "authpass"
			assert.NoError(t, db.RegisterUser(authUsername, authPassword))

			tests := []struct {
				name     string
				username string
				password string
				wantErr  error
			}{
				{"Success", authUsername, authPassword, nil},
				{"Wrong password", authUsername, "wrong", domen.ErrInvalidCredentials},
				{"Nonexistent user", "nouser", "nopass", domen.ErrUserNotFound},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					_, err := db.AuthUser(tt.username, tt.password)
					if tt.wantErr != nil {

						assert.ErrorIs(t, err, tt.wantErr)
					} else {
						assert.NoError(t, err)
					}
				})
			}
		})
	})

	t.Run("ContactOperations", func(t *testing.T) {
		userID := setupTestUser(t, db, "contactuser", "contactpass")

		t.Run("AddContact", func(t *testing.T) {
			err := db.AddContact(userID, domen.TestContactName, domen.TestContactPhone)
			assert.NoError(t, err)
		})

		t.Run("GetContacts", func(t *testing.T) {
			contacts, err := db.GetContacts(userID)
			assert.NoError(t, err)
			assert.Len(t, contacts, 1)
			assert.Equal(t, domen.TestContactName, contacts[0].Name)
		})

		t.Run("FindContact", func(t *testing.T) {
			contacts, err := db.FindContact(userID, "Doe")
			assert.NoError(t, err)
			assert.Len(t, contacts, 1)
		})

		t.Run("DeleteContact", func(t *testing.T) {
			err := db.DelContact(userID, domen.TestContactName)
			assert.NoError(t, err)

			contacts, err := db.GetContacts(userID)
			assert.NoError(t, err)
			assert.Empty(t, contacts)
		})
	})

	t.Run("ErrorCases", func(t *testing.T) {
		t.Run("AddContactToNonexistentUser", func(t *testing.T) {
			mockLogger.ClearLogs()
			err := db.AddContact(999, "Test", "+123")
			assert.ErrorIs(t, err, domen.ErrUserNotFound)
			assert.True(t, mockLogger.Contains(domen.ErrUserNotFound.Error()),
				"Expected log to contain '%s', got logs: %v",
				domen.ErrUserNotFound.Error(), mockLogger.GetLogs())
		})

		t.Run("GetContactsForNonexistentUser", func(t *testing.T) {
			mockLogger.ClearLogs()
			contacts, err := db.GetContacts(999)
			assert.Nil(t, contacts)
			assert.ErrorIs(t, err, domen.ErrUserNotFound)
			assert.True(t, mockLogger.Contains(domen.ErrUserNotFound.Error()),
				"Expected log to contain '%s', got logs: %v",
				domen.ErrUserNotFound.Error(), mockLogger.GetLogs())
		})
	})
}
