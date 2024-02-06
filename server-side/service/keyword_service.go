package service

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"server-side/model"
	"server-side/repository"
	"server-side/response"
)

func GetKeywordPage(pageRequest model.PaginationRequest[model.KeywordFilter]) (*model.PageResponse[model.KeywordFilter, model.Keyword], error) {
	keyword, err := repository.KeywordRepositoryInstance.GetFilteredKeywords(pageRequest.Filter, pageRequest.Page, pageRequest.PageSize)
	if err != nil {
		log.Error(err)
		return nil, response.NewErrorResponse("unable to search for keyword", http.StatusInternalServerError)
	}
	pageResponse := model.PageResponse[model.KeywordFilter, model.Keyword]{Pagination: pageRequest, Data: keyword}
	return &pageResponse, nil
}

func AddKeyword(keywords []string) ([]model.Keyword, error) {
	jobList, err := repository.KeywordRepositoryInstance.CreateKeywordJob(keywords)
	if err != nil {
		log.Error(err)
		return nil, response.NewErrorResponse("unable to create keyword", http.StatusInternalServerError)
	}
	return jobList, nil
}
