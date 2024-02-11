package repository_test

import (
	"github.com/stretchr/testify/mock"
	"server-side/model"
)

type PageDataRepositoryMock struct {
	mock.Mock
}

func (m *PageDataRepositoryMock) Create(pageData *model.PageData) (*model.PageData, error) {
	args := m.Called(pageData)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.PageData), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *PageDataRepositoryMock) FindByID(id uint) (*model.PageData, error) {
	args := m.Called(id)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.PageData), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *PageDataRepositoryMock) Update(pageData *model.PageData) (*model.PageData, error) {
	args := m.Called(pageData)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.PageData), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *PageDataRepositoryMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *PageDataRepositoryMock) FindBySearchResultID(searchResultID uint) (*model.PageData, error) {
	args := m.Called(searchResultID)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.PageData), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}
