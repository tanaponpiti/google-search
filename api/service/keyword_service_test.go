package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"server-side/boothstrap"
	"server-side/boothstrap/mock"
	"server-side/model"
	"server-side/repository"
	"server-side/repository/mock"
	"sync"
	"testing"
	"time"
)

func TestGetKeywordPage(t *testing.T) {
	mockKeywordRepo := new(repository_test.KeywordRepositoryMock)
	repository.KeywordRepositoryInstance = mockKeywordRepo

	pageRequest := model.PaginationRequest[model.KeywordFilter]{Page: 1, PageSize: 10}
	mockKeywords := []model.Keyword{{KeywordText: "test"}}
	total := int64(1)

	mockKeywordRepo.On("GetFilteredKeywords", pageRequest.Filter, pageRequest.Page, pageRequest.PageSize).Return(mockKeywords, total, nil)

	result, err := GetKeywordPage(pageRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(mockKeywords), len(result.Data))
	mockKeywordRepo.AssertExpectations(t)
}

func TestAddKeyword(t *testing.T) {
	mockKeywordRepo := new(repository_test.KeywordRepositoryMock)
	repository.KeywordRepositoryInstance = mockKeywordRepo

	keywords := []string{"keyword1", "keyword2"}

	cleanedKeywords := cleanUp(keywords)
	expectedKeywords := []model.Keyword{
		{ID: 1, KeywordText: "keyword1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, KeywordText: "keyword2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	mockKeywordRepo.On("CreateKeywordJob", cleanedKeywords).Return(expectedKeywords, nil)

	result, err := AddKeyword(keywords)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedKeywords, result, "The result should match the expected keywords")
	mockKeywordRepo.AssertExpectations(t)
}

func TestScrapeFromGoogleSearch(t *testing.T) {
	// Mock dependencies
	mockScraper := boothstrap_test.NewScraperMock()
	boothstrap.ScraperInstance = mockScraper
	mockKeywordRepo := new(repository_test.KeywordRepositoryMock)
	repository.KeywordRepositoryInstance = mockKeywordRepo
	mockSearchResultRepo := new(repository_test.SearchResultRepositoryMock)
	repository.SearchResultRepositoryInstance = mockSearchResultRepo

	keywords := []string{"test"}
	expectedKeywords := []model.Keyword{
		{KeywordText: "test"},
	}
	mockKeywordRepo.On("CreateKeywordJob", keywords).Return(expectedKeywords, nil)
	expectedResult := []model.Keyword{{KeywordText: "test"}}
	mockScraper.On("ScrapeFromGoogleSearch", keywords).Return(make(chan model.KeywordScrapeResult), new(sync.WaitGroup))
	mockSearchResultRepo.On("Update", mock.AnythingOfType("*model.SearchResult")).Return(nil, nil)
	mockSearchResultRepo.On("BulkUpdate", mock.AnythingOfType("[]*model.SearchResult")).Return(nil, nil)

	var wg sync.WaitGroup
	result, err := ScrapeFromGoogleSearch(keywords, &wg)
	wg.Wait()

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	mockScraper.AssertExpectations(t)
}

func TestGetSearchResultHTMLCache(t *testing.T) {
	mockPageDataRepo := new(repository_test.PageDataRepositoryMock)
	repository.PageDataRepositoryInstance = mockPageDataRepo

	searchResultId := uint(1)
	mockPageData := &model.PageData{HtmlData: "testHTML"}

	mockPageDataRepo.On("FindBySearchResultID", searchResultId).Return(mockPageData, nil)

	result, err := GetSearchResultHTMLCache(searchResultId)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockPageData.HtmlData, *result)

	mockPageDataRepo.AssertExpectations(t)
}
