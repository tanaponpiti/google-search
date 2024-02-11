package repository_test

import (
	"github.com/stretchr/testify/mock"
	"server-side/model"
)

type SearchResultRepositoryMock struct {
	mock.Mock
}

func (m *SearchResultRepositoryMock) Create(searchResult *model.SearchResult) (*model.SearchResult, error) {
	args := m.Called(searchResult)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.SearchResult), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *SearchResultRepositoryMock) FindByID(id uint) (*model.SearchResult, error) {
	args := m.Called(id)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.SearchResult), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *SearchResultRepositoryMock) Update(searchResult *model.SearchResult) (*model.SearchResult, error) {
	args := m.Called(searchResult)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*model.SearchResult), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *SearchResultRepositoryMock) BulkUpdate(searchResults []*model.SearchResult) ([]*model.SearchResult, error) {
	args := m.Called(searchResults)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.([]*model.SearchResult), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *SearchResultRepositoryMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
