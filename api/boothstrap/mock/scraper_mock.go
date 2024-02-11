package boothstrap_test

import (
	"github.com/stretchr/testify/mock"
	"server-side/model"
	"sync"
)

type ScraperMock struct {
	mock.Mock
}

func NewScraperMock() *ScraperMock {
	return &ScraperMock{}
}

func (m *ScraperMock) ScrapeFromGoogleSearch(keywords []string) (chan model.KeywordScrapeResult, *sync.WaitGroup) {
	args := m.Called(keywords)
	return args.Get(0).(chan model.KeywordScrapeResult), args.Get(1).(*sync.WaitGroup)
}
