package repository

import (
	"gorm.io/gorm"
	"server-side/database"
	"server-side/model"
)

type IKeywordRepository interface {
	Migrate() error
	Create(keyword *model.Keyword) (*model.Keyword, error)
	FindByID(id uint) (*model.Keyword, error)
	Update(keyword *model.Keyword) (*model.Keyword, error)
	Delete(id uint) error
}

var KeywordRepositoryInstance IKeywordRepository

type KeywordRepository struct {
	db *gorm.DB
}

func InitKeywordRepository() error {
	KeywordRepositoryInstance = NewKeywordRepository(database.GormDB)
	err := KeywordRepositoryInstance.Migrate()
	return err
}

func NewKeywordRepository(db *gorm.DB) *KeywordRepository {
	return &KeywordRepository{db: db}
}

func (r *KeywordRepository) Migrate() error {
	return r.db.AutoMigrate(&model.Keyword{}, &model.SearchResult{}, &model.PageData{})
}

func (r *KeywordRepository) Create(keyword *model.Keyword) (*model.Keyword, error) {
	if err := r.db.Create(&keyword).Error; err != nil {
		return nil, err
	}
	return keyword, nil
}

func (r *KeywordRepository) FindByID(id uint) (*model.Keyword, error) {
	var keyword model.Keyword
	if err := r.db.First(&keyword, id).Error; err != nil {
		return nil, err
	}
	return &keyword, nil
}

func (r *KeywordRepository) Update(keyword *model.Keyword) (*model.Keyword, error) {
	if err := r.db.Save(&keyword).Error; err != nil {
		return nil, err
	}
	return keyword, nil
}

func (r *KeywordRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Keyword{}, id).Error; err != nil {
		return err
	}
	return nil
}
