package repository_test

import (
	"github.com/stretchr/testify/mock"
	"server-side/model"
)

type KeywordRepositoryMock struct {
	mock.Mock
}

func (m *KeywordRepositoryMock) Migrate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *KeywordRepositoryMock) Create(keyword string) (*model.Keyword, error) {
	args := m.Called(keyword)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.Keyword), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *KeywordRepositoryMock) CreateKeywordJob(keywords []string) ([]model.Keyword, error) {
	args := m.Called(keywords)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.([]model.Keyword), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *KeywordRepositoryMock) GetFilteredKeywords(filter *model.KeywordFilter, page, pageSize int) ([]model.Keyword, int64, error) {
	args := m.Called(filter, page, pageSize)
	arg0 := args.Get(0)
	arg1 := args.Get(1)
	if arg0 != nil && !args.Is(nil, 1) {
		return arg0.([]model.Keyword), arg1.(int64), args.Error(2)
	} else {
		return nil, 0, args.Error(2)
	}
}

func (m *KeywordRepositoryMock) FindByID(id uint) (*model.Keyword, error) {
	args := m.Called(id)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.Keyword), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *KeywordRepositoryMock) Update(keyword *model.Keyword) (*model.Keyword, error) {
	args := m.Called(keyword)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.Keyword), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *KeywordRepositoryMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
