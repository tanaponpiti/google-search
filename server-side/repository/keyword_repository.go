package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server-side/database"
	"server-side/model"
	"time"
)

type IKeywordRepository interface {
	Migrate() error
	Create(keyword string) (*model.Keyword, error)
	CreateKeywordJob(keywords []string) ([]model.Keyword, error)
	GetFilteredKeywords(filter *model.KeywordFilter, page, pageSize int) ([]model.Keyword, int64, error)
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

func (r *KeywordRepository) Create(keywordString string) (*model.Keyword, error) {
	var keyword = &model.Keyword{
		KeywordText: keywordString,
	}
	if err := r.db.Create(&keyword).Error; err != nil {
		return nil, err
	}
	return keyword, nil
}

func (r *KeywordRepository) CreateKeywordJob(keywords []string) ([]model.Keyword, error) {
	// Start a transaction.
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Prepare bulk upsert for keywords.
	var keywordRecords []model.Keyword
	for _, keywordText := range keywords {
		keywordRecords = append(keywordRecords, model.Keyword{KeywordText: keywordText, UpdatedAt: time.Now()})
	}

	// Perform bulk upsert.
	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "keyword_text"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}).Create(&keywordRecords).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Fetch all relevant keywords with their IDs.
	var updatedKeywords []model.Keyword
	if err := tx.Where("keyword_text IN ?", keywords).Find(&updatedKeywords).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Process SearchResult scenarios.
	var pendingKeyword []model.Keyword
	for _, keyword := range updatedKeywords {
		var pendingResult model.SearchResult
		result := tx.Where("keyword_id = ? AND status = ?", keyword.ID, model.Pending).First(&pendingResult)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newSearchResult := model.SearchResult{
				KeywordID: keyword.ID,
				Status:    model.Pending,
			}
			if err := tx.Create(&newSearchResult).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
			keyword.SearchResults = append(keyword.SearchResults, newSearchResult)
			pendingKeyword = append(pendingKeyword, keyword)
		} else if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		// If a pending SearchResult already exists, do nothing.
	}

	// Commit the transaction.
	err := tx.Commit().Error
	return pendingKeyword, err
}

func (r *KeywordRepository) GetFilteredKeywords(filter *model.KeywordFilter, page, pageSize int) ([]model.Keyword, int64, error) {
	var keywords []model.Keyword
	query := r.db.Model(&model.Keyword{})
	if filter != nil {
		if filter.KeywordSearch != nil {
			query = query.Where("keyword_text LIKE ?", *filter.KeywordSearch+"%")
		}
	}
	var totalMatched int64
	if err := query.Count(&totalMatched).Error; err != nil {
		return nil, 0, err
	}
	query = query.Order("updated_at DESC").Offset((page - 1) * pageSize).Limit(pageSize)
	if err := query.Find(&keywords).Error; err != nil {
		return nil, 0, err
	}
	for i := range keywords {
		if err := r.db.Model(&keywords[i]).
			Preload("SearchResults", func(db *gorm.DB) *gorm.DB {
				return db.Order("search_results.updated_at DESC").Limit(5)
			}).Find(&keywords[i]).Error; err != nil {
			return nil, 0, err
		}
	}
	return keywords, totalMatched, nil
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
