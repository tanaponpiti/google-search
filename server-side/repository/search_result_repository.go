package repository

import (
	"gorm.io/gorm"
	"server-side/database"
	"server-side/model"
)

type ISearchResultRepository interface {
	Create(searchResult *model.SearchResult) (*model.SearchResult, error)
	FindByID(id uint) (*model.SearchResult, error)
	Update(searchResult *model.SearchResult) (*model.SearchResult, error)
	Delete(id uint) error
}

var SearchResultRepositoryInstance ISearchResultRepository

type SearchResultRepository struct {
	db *gorm.DB
}

func InitSearchResultRepository() {
	SearchResultRepositoryInstance = NewSearchResultRepository(database.GormDB)
}

func NewSearchResultRepository(db *gorm.DB) *SearchResultRepository {
	return &SearchResultRepository{db: db}
}

func (r *SearchResultRepository) Create(searchResult *model.SearchResult) (*model.SearchResult, error) {
	if err := r.db.Create(&searchResult).Error; err != nil {
		return nil, err
	}
	return searchResult, nil
}

func (r *SearchResultRepository) FindByID(id uint) (*model.SearchResult, error) {
	var searchResult model.SearchResult
	if err := r.db.First(&searchResult, id).Error; err != nil {
		return nil, err
	}
	return &searchResult, nil
}

func (r *SearchResultRepository) Update(searchResult *model.SearchResult) (*model.SearchResult, error) {
	if err := r.db.Save(&searchResult).Error; err != nil {
		return nil, err
	}
	return searchResult, nil
}

func (r *SearchResultRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.SearchResult{}, id).Error; err != nil {
		return err
	}
	return nil
}
