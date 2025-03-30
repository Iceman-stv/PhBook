package mockjwt

import "github.com/stretchr/testify/mock"

type MockJWTGenerator struct {
	mock.Mock
}

func (m *MockJWTGenerator) GenerateJWT(userID int) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}
