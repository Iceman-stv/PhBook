package userCase

import (
	"PhBook/domen"
	"PhBook/domen/testconst"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) RegisterUser(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

func (m *MockDatabase) AuthUser(username, password string) (int, error) {
	args := m.Called(username, password)
	return args.Int(0), args.Error(1)
}

func (m *MockDatabase) AddContact(userID int, name, phone string) error {
	args := m.Called(userID, name, phone)
	return args.Error(0)
}

func (m *MockDatabase) DelContact(userID int, name string) error {
	args := m.Called(userID, name)
	return args.Error(0)
}

func (m *MockDatabase) FindContact(userID int, name string) ([]domen.Contact, error) {
	args := m.Called(userID, name)
	return args.Get(0).([]domen.Contact), args.Error(1)
}

func (m *MockDatabase) GetContacts(userID int) ([]domen.Contact, error) {
	args := m.Called(userID)
	return args.Get(0).([]domen.Contact), args.Error(1)
}

func TestPhoneBook(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	t.Run("RegisterUser", func(t *testing.T) {
		tests := []struct {
			name     string
			username string
			password string
			mockErr  error
			wantErr  error
		}{
			{
				name:     "Success",
				username: testconst.TestUsername,
				password: testconst.TestPassword,
			},
			{
				name:     "User exists",
				username: "existing",
				password: "pass",
				mockErr:  domen.ErrUserExists,
				wantErr:  domen.ErrUserExists,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockDB.On("RegisterUser", tt.username, tt.password).Return(tt.mockErr)
				err := pb.RegisterUser(tt.username, tt.password)

				if tt.wantErr != nil {

					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)
				}
				mockDB.AssertExpectations(t)
			})
		}
	})

	t.Run("AuthUser", func(t *testing.T) {
		tests := []struct {
			name     string
			username string
			password string
			mockID   int
			mockErr  error
			wantID   int
			wantErr  error
		}{
			{
				name:     "Success",
				username: testconst.TestUsername,
				password: testconst.TestPassword,
				mockID:   testconst.TestUserID,
				wantID:   testconst.TestUserID,
			},
			{
				name:     "Invalid credentials",
				username: "invalid",
				password: "wrong",
				mockErr:  domen.ErrInvalidCredentials,
				wantErr:  domen.ErrInvalidCredentials,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockDB.On("AuthUser", tt.username, tt.password).Return(tt.mockID, tt.mockErr)
				id, err := pb.AuthUser(tt.username, tt.password)

				if tt.wantErr != nil {

					assert.ErrorIs(t, err, tt.wantErr)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.wantID, id)
				}
				mockDB.AssertExpectations(t)
			})
		}
	})

	t.Run("ContactOperations", func(t *testing.T) {
		testContact := domen.Contact{
			ID:     1,
			Name:   testconst.TestContactName,
			Phone:  testconst.TestContactPhone,
			UserID: testconst.TestUserID,
		}

		t.Run("AddContact", func(t *testing.T) {
			mockDB.On("AddContact", testconst.TestUserID, testconst.TestContactName, testconst.TestContactPhone).Return(nil)
			err := pb.AddContact(testconst.TestUserID, testconst.TestContactName, testconst.TestContactPhone)
			assert.NoError(t, err)
			mockDB.AssertExpectations(t)
		})

		t.Run("FindContact", func(t *testing.T) {
			mockDB.On("FindContact", testconst.TestUserID, "Doe").Return([]domen.Contact{testContact}, nil)
			contacts, err := pb.FindContact(testconst.TestUserID, "Doe")
			assert.NoError(t, err)
			assert.Len(t, contacts, 1)
			assert.Equal(t, testconst.TestContactName, contacts[0].Name)
			mockDB.AssertExpectations(t)
		})

		t.Run("GetContacts", func(t *testing.T) {
			mockDB.On("GetContacts", testconst.TestUserID).Return([]domen.Contact{testContact}, nil)
			contacts, err := pb.GetContacts(testconst.TestUserID)
			assert.NoError(t, err)
			assert.Len(t, contacts, 1)
			mockDB.AssertExpectations(t)
		})

		t.Run("DelContact", func(t *testing.T) {
			mockDB.On("DelContact", testconst.TestUserID, testconst.TestContactName).Return(nil)
			err := pb.DelContact(testconst.TestUserID, testconst.TestContactName)
			assert.NoError(t, err)
			mockDB.AssertExpectations(t)
		})
	})
}
