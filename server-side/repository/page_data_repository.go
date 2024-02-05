package repository

import (
	"gorm.io/gorm"
	"server-side/database"
	"server-side/model"
)

type IPageDataRepository interface {
	Create(pageData *model.PageData) (*model.PageData, error)
	FindByID(id uint) (*model.PageData, error)
	Update(pageData *model.PageData) (*model.PageData, error)
	Delete(id uint) error
}

var PageDataRepositoryInstance IPageDataRepository

type PageDataRepository struct {
	db *gorm.DB
}

func InitPageDataRepository() {
	PageDataRepositoryInstance = NewPageDataRepository(database.GormDB)
}

func NewPageDataRepository(db *gorm.DB) *PageDataRepository {
	return &PageDataRepository{db: db}
}

func (r *PageDataRepository) Create(pageData *model.PageData) (*model.PageData, error) {
	if err := r.db.Create(&pageData).Error; err != nil {
		return nil, err
	}
	return pageData, nil
}

func (r *PageDataRepository) FindByID(id uint) (*model.PageData, error) {
	var pageData model.PageData
	if err := r.db.First(&pageData, id).Error; err != nil {
		return nil, err
	}
	return &pageData, nil
}

func (r *PageDataRepository) Update(pageData *model.PageData) (*model.PageData, error) {
	if err := r.db.Save(&pageData).Error; err != nil {
		return nil, err
	}
	return pageData, nil
}

func (r *PageDataRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.PageData{}, id).Error; err != nil {
		return err
	}
	return nil
}
