package interfaceGrpc

import (
	"PhBook/domen"
	pB "PhBook/gen/github.com/Iceman-stv/PhBook/gen"
	"PhBook/logger/mocklog"
	"PhBook/server/jwt/mockjwt"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPhoneBook struct {
	mock.Mock
}

func (m *MockPhoneBook) RegisterUser(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

func (m *MockPhoneBook) AuthUser(username, password string) (int, error) {
	args := m.Called(username, password)
	return args.Int(0), args.Error(1)
}

func (m *MockPhoneBook) AddContact(userID int, name, phone string) error {
	args := m.Called(userID, name, phone)
	return args.Error(0)
}

func (m *MockPhoneBook) DelContact(userID int, name string) error {
	args := m.Called(userID, name)
	return args.Error(0)
}

func (m *MockPhoneBook) FindContact(userID int, name string) ([]domen.Contact, error) {
	args := m.Called(userID, name)
	return args.Get(0).([]domen.Contact), args.Error(1)
}

func (m *MockPhoneBook) GetContacts(userID int) ([]domen.Contact, error) {
	args := m.Called(userID)
	return args.Get(0).([]domen.Contact), args.Error(1)
}

/* ----- AuthUser тесты ----- */
func TestAuthUser_Success(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()
	mockJWT := new(mockjwt.MockJWTGenerator)

	mockPB.On("AuthUser", "valid", "pass").Return(1, nil)
	mockJWT.On("GenerateJWT", 1).Return("test-token", nil)

	h := NewTestPhoneBookHandlers(mockPB, mockLogger, mockJWT)
	resp, err := h.AuthUser(context.Background(), &pB.AuthUserRequest{
		Username: "valid",
		Password: "pass",
	})

	assert.NoError(t, err)
	assert.Equal(t, "test-token", resp.Token)
	mockPB.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

func TestAuthUser_InvalidCredentials(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	authErr := errors.New("invalid credentials")
	mockPB.On("AuthUser", "invalid", "pass").Return(0, authErr)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.AuthUser(context.Background(), &pB.AuthUserRequest{
		Username: "invalid",
		Password: "pass",
	})

	assert.EqualError(t, err, "invalid credentials")
	assert.True(t, mockLogger.Contains("invalid credentials"))
}

func TestAuthUser_JWTGenerationError(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()
	mockJWT := new(mockjwt.MockJWTGenerator)

	mockPB.On("AuthUser", "valid", "pass").Return(1, nil)
	jwtErr := errors.New("jwt ошибка")
	mockJWT.On("GenerateJWT", 1).Return("", jwtErr)

	h := NewTestPhoneBookHandlers(mockPB, mockLogger, mockJWT)
	_, err := h.AuthUser(context.Background(), &pB.AuthUserRequest{
		Username: "valid",
		Password: "pass",
	})

	assert.EqualError(t, err, "jwt ошибка")
	assert.True(t, mockLogger.Contains("jwt ошибка"))
}

/* ----- RegisterUser тесты ----- */
func TestRegisterUser_Success(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	mockPB.On("RegisterUser", "newuser", "pass123").Return(nil)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.RegisterUser(context.Background(), &pB.RegisterUserRequest{
		Username: "newuser",
		Password: "pass123",
	})

	assert.NoError(t, err)
	assert.False(t, mockLogger.Contains("ERROR"))
}

func TestRegisterUser_EmptyCredentials(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	h := NewPhoneBookHandlers(mockPB, mockLogger)

	// пустой username
	_, err := h.RegisterUser(context.Background(), &pB.RegisterUserRequest{
		Username: "",
		Password: "pass",
	})
	assert.Error(t, err)
	assert.True(t, mockLogger.Contains("пустое имя или пароль"))
	mockPB.AssertNotCalled(t, "RegisterUser")

	// пустой password
	mockLogger.ClearLogs() // Очищаем логи между тестами
	_, err = h.RegisterUser(context.Background(), &pB.RegisterUserRequest{
		Username: "user",
		Password: "",
	})
	assert.Error(t, err)
	assert.True(t, mockLogger.Contains("пустое имя или пароль"))
	mockPB.AssertNotCalled(t, "RegisterUser")
}

func TestRegisterUser_DBError(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	dbErr := errors.New("user exists")
	mockPB.On("RegisterUser", "existing", "pass").Return(dbErr)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.RegisterUser(context.Background(), &pB.RegisterUserRequest{
		Username: "existing",
		Password: "pass",
	})

	assert.EqualError(t, err, "user exists")
	assert.True(t, mockLogger.Contains("user exists"))
}

/* ----- AddContact тесты ----- */
func TestAddContact_Success(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	mockPB.On("AddContact", 1, "John", "+123456789").Return(nil)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.AddContact(context.Background(), &pB.AddContactRequest{
		UserId: 1,
		Name:   "John",
		Phone:  "+123456789",
	})

	assert.NoError(t, err)
	assert.False(t, mockLogger.Contains("ERROR"))
	mockPB.AssertExpectations(t)
}

func TestAddContact_ValidationError(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	h := NewPhoneBookHandlers(mockPB, mockLogger)

	// пустой name
	_, err := h.AddContact(context.Background(), &pB.AddContactRequest{
		UserId: 1,
		Name:   "",
		Phone:  "+123",
	})
	assert.Error(t, err)
	assert.True(t, mockLogger.Contains("пустое имя или телефон"))
	mockPB.AssertNotCalled(t, "AddContact")

	// пустой phone
	mockLogger.ClearLogs()
	_, err = h.AddContact(context.Background(), &pB.AddContactRequest{
		UserId: 1,
		Name:   "John",
		Phone:  "",
	})
	assert.Error(t, err)
	assert.True(t, mockLogger.Contains("пустое имя или телефон"))
	mockPB.AssertNotCalled(t, "AddContact")
}

func TestAddContact_DBError(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	dbErr := errors.New("database error")
	mockPB.On("AddContact", 1, "John", "+123").Return(dbErr)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.AddContact(context.Background(), &pB.AddContactRequest{
		UserId: 1,
		Name:   "John",
		Phone:  "+123",
	})

	assert.EqualError(t, err, "database error")
	assert.True(t, mockLogger.Contains("database error"))
	mockPB.AssertExpectations(t)
}

/* ----- DelContact тесты ----- */
func TestDelContact_Success(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	mockPB.On("DelContact", 1, "John").Return(nil)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.DelContact(context.Background(), &pB.DelContactRequest{
		UserId: 1,
		Name:   "John",
	})

	assert.NoError(t, err)
}

func TestDelContact_NotFound(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	notFoundErr := errors.New("not found")
	mockPB.On("DelContact", 1, "Unknown").Return(notFoundErr)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.DelContact(context.Background(), &pB.DelContactRequest{
		UserId: 1,
		Name:   "Unknown",
	})

	assert.EqualError(t, err, "not found")
	assert.True(t, mockLogger.Contains("not found"))
}

/* ----- FindContact тесты ----- */
func TestFindContact_Success(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	expected := []domen.Contact{
		{Name: "John", Phone: "+123"},
	}
	mockPB.On("FindContact", 1, "John").Return(expected, nil)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	resp, err := h.FindContact(context.Background(), &pB.FindContactRequest{
		UserId: 1,
		Name:   "John",
	})

	assert.NoError(t, err)
	assert.Len(t, resp.Contacts, 1)
	assert.Equal(t, "John", resp.Contacts[0].Name)
}

func TestFindContact_EmptyResult(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	mockPB.On("FindContact", 1, "Unknown").Return([]domen.Contact{}, nil)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	resp, err := h.FindContact(context.Background(), &pB.FindContactRequest{
		UserId: 1,
		Name:   "Unknown",
	})

	assert.NoError(t, err)
	assert.Empty(t, resp.Contacts)
}

/* ----- GetContacts тесты ----- */
func TestGetContacts_Success(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	expected := []domen.Contact{
		{Name: "John", Phone: "+123"},
		{Name: "Alice", Phone: "+456"},
	}
	mockPB.On("GetContacts", 1).Return(expected, nil)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	resp, err := h.GetContacts(context.Background(), &pB.GetContactsRequest{
		UserId: 1,
	})

	assert.NoError(t, err)
	assert.Len(t, resp.Contacts, 2)
}

func TestGetContacts_DBError(t *testing.T) {
	mockPB := new(MockPhoneBook)
	mockLogger := mocklog.NewMockLogger()

	dbErr := errors.New("db error")
	mockPB.On("GetContacts", 1).Return([]domen.Contact{}, dbErr)

	h := NewPhoneBookHandlers(mockPB, mockLogger)
	_, err := h.GetContacts(context.Background(), &pB.GetContactsRequest{
		UserId: 1,
	})

	assert.EqualError(t, err, "db error")
	assert.True(t, mockLogger.Contains("db error"))
}
