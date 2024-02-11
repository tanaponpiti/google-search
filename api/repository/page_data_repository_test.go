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
)

func newMockDB() *gorm.DB {
	uniqueID := uuid.New().String()
	connectionString := "file:" + uniqueID + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.PageData{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func TestPageDataRepositoryCreate(t *testing.T) {
	db := newMockDB()
	repo := repository.NewPageDataRepository(db)

	pageData := &model.PageData{
		SearchResultID: 1,
		HtmlData:       "Test Content",
	}

	createdPageData, err := repo.Create(pageData)
	require.NoError(t, err)
	assert.NotNil(t, createdPageData)
	assert.NotZero(t, createdPageData.ID)
	assert.Equal(t, "Test Content", createdPageData.HtmlData)
}

func TestPageDataRepositoryFindByID(t *testing.T) {
	db := newMockDB()
	repo := repository.NewPageDataRepository(db)

	pageData := &model.PageData{
		SearchResultID: 1,
		HtmlData:       "Test Content",
	}
	createdPageData, _ := repo.Create(pageData)

	fetchedPageData, err := repo.FindByID(createdPageData.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedPageData)
	assert.Equal(t, "Test Content", fetchedPageData.HtmlData)
}

func TestPageDataRepositoryUpdate(t *testing.T) {
	db := newMockDB()
	repo := repository.NewPageDataRepository(db)

	pageData := &model.PageData{
		SearchResultID: 1,
		HtmlData:       "Initial Content",
	}
	createdPageData, _ := repo.Create(pageData)

	createdPageData.HtmlData = "Updated Content"
	updatedPageData, err := repo.Update(createdPageData)
	require.NoError(t, err)
	assert.Equal(t, "Updated Content", updatedPageData.HtmlData)
}

func TestPageDataRepositoryDelete(t *testing.T) {
	db := newMockDB()
	repo := repository.NewPageDataRepository(db)

	pageData := &model.PageData{
		SearchResultID: 1,
		HtmlData:       "Content to Delete",
	}
	createdPageData, _ := repo.Create(pageData)

	err := repo.Delete(createdPageData.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(createdPageData.ID)
	assert.Error(t, err)
}

func TestPageDataRepositoryFindBySearchResultID(t *testing.T) {
	t.Run("Get exist page data by search result ID", func(t *testing.T) {
		db := newMockDB()
		repo := repository.NewPageDataRepository(db)

		pageData := &model.PageData{
			SearchResultID: 1,
			HtmlData:       "Content for Specific Search Result",
		}
		repo.Create(pageData)

		fetchedPageData, err := repo.FindBySearchResultID(1)
		require.NoError(t, err)
		assert.NotNil(t, fetchedPageData)
		assert.Equal(t, "Content for Specific Search Result", fetchedPageData.HtmlData)
	})
	t.Run("Get non-exist page data by search result ID", func(t *testing.T) {
		db := newMockDB()
		repo := repository.NewPageDataRepository(db)
		fetchedPageData, err := repo.FindBySearchResultID(1)
		require.NoError(t, err)
		assert.Nil(t, fetchedPageData)
	})
}
