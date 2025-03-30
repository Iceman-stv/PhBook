package userCase

import (
	"PhBook/domen"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDatabase реализует интерфейс Database
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

func TestNewPhoneBook(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	assert.NotNil(t, pb)
	assert.Equal(t, mockDB, pb.db)
}

func TestRegisterUser_Success(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	// Настройка ожиданий
	mockDB.On("RegisterUser", "testuser", "testpass").Return(nil)

	// Вызов методов
	err := pb.RegisterUser("testuser", "testpass")

	// Проверка
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestRegisterUser_Error(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	expectedErr := errors.New("user already exists")
	mockDB.On("RegisterUser", "existing", "pass").Return(expectedErr)

	err := pb.RegisterUser("existing", "pass")

	assert.EqualError(t, err, expectedErr.Error())
	mockDB.AssertExpectations(t)
}

func TestAuthUser_Success(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	expectedID := 1
	mockDB.On("AuthUser", "valid", "pass").Return(expectedID, nil)

	id, err := pb.AuthUser("valid", "pass")

	assert.NoError(t, err)
	assert.Equal(t, expectedID, id)
	mockDB.AssertExpectations(t)
}

func TestAddContact_Success(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	userID := 1
	name := "John"
	phone := "+1234567890"

	mockDB.On("AddContact", userID, name, phone).Return(nil)

	err := pb.AddContact(userID, name, phone)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDelContact_Success(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	userID := 1
	name := "John"

	mockDB.On("DelContact", userID, name).Return(nil)

	err := pb.DelContact(userID, name)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestFindContact_Success(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	userID := 1
	name := "John"
	expectedContacts := []domen.Contact{
		{ID: 1, Name: "John Doe", Phone: "+1234567890"},
	}

	mockDB.On("FindContact", userID, name).Return(expectedContacts, nil)

	contacts, err := pb.FindContact(userID, name)

	assert.NoError(t, err)
	assert.Equal(t, expectedContacts, contacts)
	mockDB.AssertExpectations(t)
}

func TestGetContacts_Success(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	userID := 1
	expectedContacts := []domen.Contact{
		{ID: 1, Name: "John Doe", Phone: "+1234567890"},
		{ID: 2, Name: "Jane Smith", Phone: "+9876543210"},
	}

	mockDB.On("GetContacts", userID).Return(expectedContacts, nil)

	contacts, err := pb.GetContacts(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedContacts, contacts)
	mockDB.AssertExpectations(t)
}

func TestAuthUser_InvalidCredentials(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	expectedErr := errors.New("invalid credentials")
	mockDB.On("AuthUser", "invalid", "pass").Return(0, expectedErr)

	_, err := pb.AuthUser("invalid", "pass")

	assert.EqualError(t, err, expectedErr.Error())
	mockDB.AssertExpectations(t)
}

func TestFindContact_NotFound(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	userID := 1
	name := "Nonexistent"
	var emptyContacts []domen.Contact

	mockDB.On("FindContact", userID, name).Return(emptyContacts, nil)

	contacts, err := pb.FindContact(userID, name)

	assert.NoError(t, err)
	assert.Empty(t, contacts)
	mockDB.AssertExpectations(t)
}

func TestGetContacts_Empty(t *testing.T) {
	mockDB := new(MockDatabase)
	pb := NewPhoneBook(mockDB)

	userID := 2
	var emptyContacts []domen.Contact

	mockDB.On("GetContacts", userID).Return(emptyContacts, nil)

	contacts, err := pb.GetContacts(userID)

	assert.NoError(t, err)
	assert.Empty(t, contacts)
	mockDB.AssertExpectations(t)
}
