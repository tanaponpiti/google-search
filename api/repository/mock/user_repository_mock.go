package repository_test

import (
	"github.com/stretchr/testify/mock"
	"server-side/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Migrate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UserRepositoryMock) InsertUser(userCreateInput model.UserCreate) (*model.User, error) {
	args := m.Called(userCreateInput)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.User), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *UserRepositoryMock) GetUser(id int64) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.User), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *UserRepositoryMock) UpdateUser(id int64, updateData model.UserUpdate) (*model.User, error) {
	args := m.Called(id, updateData)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.User), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *UserRepositoryMock) DeleteUser(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
