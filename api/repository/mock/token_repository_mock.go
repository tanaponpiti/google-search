package repository_test

import (
	"github.com/stretchr/testify/mock"
	"server-side/model"
	"time"
)

type TokenRepositoryMock struct {
	mock.Mock
}

func (m *TokenRepositoryMock) GetTokenByTokenString(tokenString string) (*model.Token, error) {
	args := m.Called(tokenString)
	return args.Get(0).(*model.Token), args.Error(1)
}

func (m *TokenRepositoryMock) CreateToken(userId int64, token string, duration time.Duration) (*model.Token, error) {
	args := m.Called(userId, token, duration)
	return args.Get(0).(*model.Token), args.Error(1)
}

func (m *TokenRepositoryMock) DeleteToken(tokenString string) error {
	args := m.Called(tokenString)
	return args.Error(0)
}
