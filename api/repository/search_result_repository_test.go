package repository_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"server-side/model"
	"server-side/repository"
	"testing"
	"time"
)

func setupMockDBForSearchResult() *gorm.DB {
	uniqueID := uuid.New().String()
	connectionString := "file:" + uniqueID + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.SearchResult{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func TestSearchResultRepositoryCreate(t *testing.T) {
	db := setupMockDBForSearchResult()
	repo := repository.NewSearchResultRepository(db)

	searchDate := time.Now()
	adWordsCount, totalLinks := 10, 100
	totalResults := "About 1,000,000 results"
	searchResult := &model.SearchResult{
		KeywordID:    1,
		AdWordsCount: &adWordsCount,
		TotalLinks:   &totalLinks,
		TotalResults: &totalResults,
		SearchDate:   &searchDate,
		Status:       model.Pending,
	}

	createdSearchResult, err := repo.Create(searchResult)
	require.NoError(t, err)
	assert.NotNil(t, createdSearchResult)
	assert.NotZero(t, createdSearchResult.ID)
	assert.Equal(t, model.Pending, createdSearchResult.Status)
}

func TestSearchResultRepositoryFindByID(t *testing.T) {
	db := setupMockDBForSearchResult()
	repo := repository.NewSearchResultRepository(db)

	searchDate := time.Now()
	adWordsCount, totalLinks := 10, 100
	totalResults := "About 1,000,000 results"
	searchResult := &model.SearchResult{
		KeywordID:    1,
		AdWordsCount: &adWordsCount,
		TotalLinks:   &totalLinks,
		TotalResults: &totalResults,
		SearchDate:   &searchDate,
		Status:       model.Pending,
	}
	createdSearchResult, _ := repo.Create(searchResult)

	foundSearchResult, err := repo.FindByID(createdSearchResult.ID)
	require.NoError(t, err)
	assert.NotNil(t, foundSearchResult)
	assert.Equal(t, createdSearchResult.ID, foundSearchResult.ID)
	assert.Equal(t, model.Pending, foundSearchResult.Status)
}

func TestSearchResultRepositoryUpdate(t *testing.T) {
	db := setupMockDBForSearchResult()
	repo := repository.NewSearchResultRepository(db)

	searchDate := time.Now()
	adWordsCount, totalLinks := 10, 100
	totalResults := "About 1,000,000 results"
	searchResult := &model.SearchResult{
		KeywordID:    1,
		AdWordsCount: &adWordsCount,
		TotalLinks:   &totalLinks,
		TotalResults: &totalResults,
		SearchDate:   &searchDate,
		Status:       model.Pending,
	}
	createdSearchResult, _ := repo.Create(searchResult)

	createdSearchResult.Status = model.Completed
	updatedSearchResult, err := repo.Update(createdSearchResult)
	require.NoError(t, err)
	assert.Equal(t, model.Completed, updatedSearchResult.Status)
}

func TestSearchResultRepositoryBulkUpdate(t *testing.T) {
	db := setupMockDBForSearchResult()
	repo := repository.NewSearchResultRepository(db)

	searchDate := time.Now()
	adWordsCount, totalLinks := 10, 100
	totalResults := "About 1,000,000 results"
	searchResults := []*model.SearchResult{
		{
			KeywordID:    1,
			AdWordsCount: &adWordsCount,
			TotalLinks:   &totalLinks,
			TotalResults: &totalResults,
			SearchDate:   &searchDate,
			Status:       model.Pending,
		},
		{
			KeywordID:    2,
			AdWordsCount: &adWordsCount,
			TotalLinks:   &totalLinks,
			TotalResults: &totalResults,
			SearchDate:   &searchDate,
			Status:       model.Pending,
		},
	}
	for _, sr := range searchResults {
		repo.Create(sr)
	}

	for _, sr := range searchResults {
		sr.Status = model.Completed
	}
	updatedSearchResults, err := repo.BulkUpdate(searchResults)
	require.NoError(t, err)
	assert.Len(t, updatedSearchResults, len(searchResults))
	for _, usr := range updatedSearchResults {
		assert.Equal(t, model.Completed, usr.Status)
	}
}

func TestSearchResultRepositoryDelete(t *testing.T) {
	db := setupMockDBForSearchResult()
	repo := repository.NewSearchResultRepository(db)

	searchDate := time.Now()
	adWordsCount, totalLinks := 10, 100
	totalResults := "About 1,000,000 results"
	searchResult := &model.SearchResult{
		KeywordID:    1,
		AdWordsCount: &adWordsCount,
		TotalLinks:   &totalLinks,
		TotalResults: &totalResults,
		SearchDate:   &searchDate,
		Status:       model.Pending,
	}
	createdSearchResult, _ := repo.Create(searchResult)

	err := repo.Delete(createdSearchResult.ID)
	require.NoError(t, err)

	_, err = repo.FindByID(createdSearchResult.ID)
	assert.Error(t, err)
}
