package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"server-side/model"
)

// mockDB is a mock for the GORM DB
type mockDB struct {
	mock.Mock
	*gorm.DB
}

func NewMockDB() *mockDB {
	uniqueID := uuid.New().String()
	connectionString := fmt.Sprintf("file:%s?mode=memory&cache=shared", uniqueID)
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		// Handle error
		panic("failed to connect database")
	}
	return &mockDB{DB: db}
}

func TestMigrate(t *testing.T) {
	db := NewMockDB()
	repository := NewKeywordRepository(db.DB)

	err := repository.Migrate()
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	db := NewMockDB()
	repository := NewKeywordRepository(db.DB)

	// Migrate the database to ensure the Keyword table exists.
	assert.NoError(t, repository.Migrate())

	keywordString := "test keyword"
	keyword, err := repository.Create(keywordString)
	assert.NoError(t, err)
	assert.NotNil(t, keyword)
	assert.Equal(t, keywordString, keyword.KeywordText)

	// Verify keyword is created in the database
	var fetchedKeyword model.Keyword
	err = db.First(&fetchedKeyword, keyword.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, keywordString, fetchedKeyword.KeywordText)
}

func TestCreateKeywordJob(t *testing.T) {

	t.Run("Add keyword success", func(t *testing.T) {
		db := NewMockDB()
		keywordRepository := NewKeywordRepository(db.DB)
		require.NoError(t, keywordRepository.Migrate())
		keywords := []string{"keyword1", "keyword2"}
		createdKeywords, err := keywordRepository.CreateKeywordJob(keywords)
		require.NoError(t, err)
		assert.Len(t, createdKeywords, len(keywords))

		for _, keyword := range createdKeywords {
			var fetchedKeyword model.Keyword
			err := db.Preload("SearchResults").First(&fetchedKeyword, keyword.ID).Error
			require.NoError(t, err)
			assert.Equal(t, keyword.KeywordText, fetchedKeyword.KeywordText)
			assert.NotEmpty(t, fetchedKeyword.SearchResults)
		}
	})

	t.Run("Add keyword that already exists", func(t *testing.T) {
		db := NewMockDB()
		keywordRepository := NewKeywordRepository(db.DB)
		require.NoError(t, keywordRepository.Migrate())
		keywords := []string{"keyword1", "keyword2"}
		_, err := keywordRepository.CreateKeywordJob(keywords)
		require.NoError(t, err)

		newKeywords := []string{"keyword1", "keyword2", "keyword3"}
		newCreatedKeywords, err := keywordRepository.CreateKeywordJob(newKeywords)
		require.NoError(t, err)

		assert.Len(t, newCreatedKeywords, 1, "Only new keywords should be added")

		for _, keyword := range newCreatedKeywords {
			var fetchedKeyword model.Keyword
			err := db.Preload("SearchResults").First(&fetchedKeyword, keyword.ID).Error
			require.NoError(t, err)
			assert.Equal(t, "keyword3", fetchedKeyword.KeywordText, "The new keyword should match")
			assert.NotEmpty(t, fetchedKeyword.SearchResults, "Search results should not be empty")
		}
	})

	t.Run("Add keyword that already exists, and those keywords already scraped", func(t *testing.T) {
		db := NewMockDB()
		keywordRepository := NewKeywordRepository(db.DB)
		require.NoError(t, keywordRepository.Migrate())
		keywords := []string{"keyword1", "keyword2"}
		createdKeywords, err := keywordRepository.CreateKeywordJob(keywords)
		require.NoError(t, err)
		for _, keyword := range createdKeywords {
			searchResult := keyword.SearchResults[0]
			searchResult.Status = model.Completed
			db.Save(&searchResult)
		}

		newKeywords := []string{"keyword1", "keyword2", "keyword3"}
		newCreatedKeywords, err := keywordRepository.CreateKeywordJob(newKeywords)
		require.NoError(t, err)

		assert.Len(t, newCreatedKeywords, 3, "New keywords,and completed keywords should be added")
	})
}

func TestGetFilteredKeywords(t *testing.T) {
	t.Run("Get keyword without filter", func(t *testing.T) {
		db := NewMockDB()
		keywordRepository := NewKeywordRepository(db.DB)
		require.NoError(t, keywordRepository.Migrate())
		keywords := []string{"keyword1", "keyword2"}
		_, err := keywordRepository.CreateKeywordJob(keywords)
		require.NoError(t, err)
		searchResult, total, err := keywordRepository.GetFilteredKeywords(nil, 1, 10)
		assert.NoError(t, err)
		assert.Equal(t, total, int64(2), "Total keyword is not correct")
		assert.Len(t, searchResult, 2, "All of keywords should be visible")
	})
	t.Run("Get keyword with search filter", func(t *testing.T) {
		db := NewMockDB()
		keywordRepository := NewKeywordRepository(db.DB)
		require.NoError(t, keywordRepository.Migrate())
		keywords := []string{"apple", "banana", "alpaca", "cat"}
		_, err := keywordRepository.CreateKeywordJob(keywords)
		require.NoError(t, err)
		keywordFilter := "a"
		filter := model.KeywordFilter{KeywordSearch: &keywordFilter}
		searchResult, total, err := keywordRepository.GetFilteredKeywords(&filter, 1, 10)
		assert.NoError(t, err)
		assert.Equal(t, total, int64(2), "Total keyword start with given filter is not correct")
		for _, keyword := range searchResult {
			assert.True(t, strings.HasPrefix(keyword.KeywordText, keywordFilter), "Result keyword should start with given filter")
		}
	})
}

func TestFindByID(t *testing.T) {
	db := NewMockDB()
	repository := NewKeywordRepository(db.DB)
	require.NoError(t, repository.Migrate())
	keyword, _ := repository.Create("test")
	fetchedKeyword, err := repository.FindByID(keyword.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedKeyword)
	assert.Equal(t, "test", fetchedKeyword.KeywordText)
}

func TestUpdate(t *testing.T) {
	db := NewMockDB()
	repository := NewKeywordRepository(db.DB)
	require.NoError(t, repository.Migrate())
	keyword, _ := repository.Create("initial")
	keyword.KeywordText = "updated"
	updatedKeyword, err := repository.Update(keyword)
	assert.NoError(t, err)
	assert.Equal(t, "updated", updatedKeyword.KeywordText)
}

func TestDelete(t *testing.T) {
	db := NewMockDB()
	repository := NewKeywordRepository(db.DB)
	require.NoError(t, repository.Migrate())
	keyword, _ := repository.Create("deleteMe")
	err := repository.Delete(keyword.ID)
	assert.NoError(t, err)

	_, err = repository.FindByID(keyword.ID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
