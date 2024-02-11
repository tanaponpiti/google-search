package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"server-side/boothstrap"
	"server-side/model"
	"server-side/repository"
	"server-side/response"
	"strconv"
	"strings"
	"sync"
	"time"
)

func GetKeywordPage(pageRequest model.PaginationRequest[model.KeywordFilter]) (*model.PageResponse[model.KeywordFilter, model.Keyword], error) {
	keyword, total, err := repository.KeywordRepositoryInstance.GetFilteredKeywords(pageRequest.Filter, pageRequest.Page, pageRequest.PageSize)
	if err != nil {
		log.Error(err)
		return nil, response.NewErrorResponse("unable to search for keyword", http.StatusInternalServerError)
	}
	pagination := model.CreatePaginationResponse(pageRequest, total)
	pageResponse := model.PageResponse[model.KeywordFilter, model.Keyword]{Pagination: pagination, Data: keyword}
	return &pageResponse, nil
}

func cleanUp(keywords []string) []string {
	keywordMap := make(map[string]bool)
	var cleanedKeywords []string

	for _, keyword := range keywords {
		lowerKeyword := strings.ToLower(keyword)
		if _, exists := keywordMap[lowerKeyword]; !exists {
			keywordMap[lowerKeyword] = true
			cleanedKeywords = append(cleanedKeywords, lowerKeyword)
		}
	}

	return cleanedKeywords
}

func AddKeyword(keywords []string) ([]model.Keyword, error) {
	if len(keywords) > 100 {
		return nil, response.NewErrorResponse("cannot add keyword more than 100 keywords at a time", http.StatusInternalServerError)
	}
	jobList, err := repository.KeywordRepositoryInstance.CreateKeywordJob(cleanUp(keywords))
	if err != nil {
		log.Error(err)
		return nil, response.NewErrorResponse("unable to create keyword", http.StatusInternalServerError)
	}
	return jobList, nil
}

func ScrapeFromGoogleSearch(keywords []string, externalWg *sync.WaitGroup) ([]model.Keyword, error) {
	tobeScrapeList, err := AddKeyword(keywords)
	if err != nil {
		return nil, err
	}
	keywordList := make([]string, len(tobeScrapeList))
	for i, kw := range tobeScrapeList {
		keywordList[i] = kw.KeywordText
	}
	if len(keywordList) > 0 {
		if externalWg != nil {
			externalWg.Add(1) // Indicate that we have a goroutine to wait for.
		}
		go func() {
			defer func() {
				if externalWg != nil {
					externalWg.Done() // Indicate that this goroutine is done.
				}
			}()
			scrapeResultsChan, wg := boothstrap.ScraperInstance.ScrapeFromGoogleSearch(keywords)
			go func() {
				wg.Wait()
				close(scrapeResultsChan)
			}()

			keywordsSearchResultMap := make(map[string]*model.SearchResult)
			for _, kw := range tobeScrapeList {
				if len(kw.SearchResults) > 0 {
					searchResult := kw.SearchResults[0]
					if searchResult.Status == model.Pending || searchResult.Status == model.Failed {
						keywordsSearchResultMap[kw.KeywordText] = &searchResult
					} else {
						log.Warning(fmt.Sprintf("Found non Pending/Failed search result in search result. Skip updating search result with id %s", strconv.Itoa(int(searchResult.ID))))
					}
				}
			}
			for scrapeResult := range scrapeResultsChan {
				if searchResult, found := keywordsSearchResultMap[scrapeResult.Keyword]; found {
					if scrapeResult.Error == nil {
						searchResult.Status = model.Completed
						searchResult.TotalResults = scrapeResult.TotalResults
						searchResult.AdWordsCount = scrapeResult.AdWordsCount
						searchResult.TotalLinks = scrapeResult.TotalLinks
						searchResult.SearchDate = &scrapeResult.SearchDate
						searchResult.PageData = &model.PageData{
							SearchResultID: searchResult.ID,
							HtmlData:       scrapeResult.RawHTML,
						}
					} else {
						searchResult.Status = model.Failed
						searchResult.SearchDate = &scrapeResult.SearchDate
					}
					//Update each search result separately
					//Update in loop is not friendly to database, but it provides realtime update to search request data for user.
					//This can also be change to bulk update later as well if we want to wait for every keyword to complete
					_, err := repository.SearchResultRepositoryInstance.Update(searchResult)
					if err != nil {
						log.Error("unable to update search result", err)
					}
				}
			}
			currentTime := time.Now()
			failedSearchResult := make([]*model.SearchResult, 0)
			for _, searchResult := range keywordsSearchResultMap {
				if searchResult.Status == model.Pending {
					searchResult.Status = model.Failed
					searchResult.SearchDate = &currentTime
					failedSearchResult = append(failedSearchResult, searchResult)
				}
			}
			_, err := repository.SearchResultRepositoryInstance.BulkUpdate(failedSearchResult)
			if err != nil {
				log.Error("unable to update search result", err)
			}
		}()
	}
	return tobeScrapeList, nil
}

func GetSearchResultHTMLCache(searchResultId uint) (*string, error) {
	result, err := repository.PageDataRepositoryInstance.FindBySearchResultID(searchResultId)
	if err != nil {
		return nil, response.NewErrorResponse("unable to find page data of given search result id", http.StatusInternalServerError)
	}
	if result == nil {
		return nil, response.NewErrorResponse("unable to find page data of given search result id", http.StatusNotFound)
	}
	return &result.HtmlData, nil
}
