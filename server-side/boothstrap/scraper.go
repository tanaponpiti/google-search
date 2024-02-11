package boothstrap

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
	"net/url"
	"server-side/connector"
	"server-side/model"
	"strings"
	"sync"
	"time"
)

type Scraper struct {
	semaphore chan struct{}
}

var ScraperInstance *Scraper

func InitScraper() {
	ScraperInstance = NewScraper(viper.GetInt("CONCURRENT_SCRAPE_LIMIT"))
}

func NewScraper(concurrencyLimit int) *Scraper {
	return &Scraper{
		semaphore: make(chan struct{}, concurrencyLimit),
	}
}

func generateSearchURL(keyword string) string {
	encodedKeyword := url.QueryEscape(keyword)
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s&hl=EN", encodedKeyword)
	return searchURL
}

func ExtractInformation(htmlContent string) (*model.ExtractedMetadata, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	isHtml := doc.Find("*").Length() > 0
	if !isHtml {
		return nil, errors.New("given content is not HTML")
	}

	// Find and count AdWords advertisers
	advertisersCount := 0
	doc.Find(`[aria-label="Ads"]`).Each(func(i int, s *goquery.Selection) {
		s.Find("span").Each(func(_ int, span *goquery.Selection) {
			if strings.Contains(span.Text(), "Sponsored") {
				advertisersCount++
			}
		})
	})

	// Find and count all links
	linksCount := doc.Find("a").Length()

	// Find the total search results text
	var searchResultsText string
	doc.Find("div#result-stats").Each(func(i int, s *goquery.Selection) {
		searchResultsText = s.Text()
	})

	return &model.ExtractedMetadata{
		AdWordsCount: &advertisersCount,
		TotalLinks:   &linksCount,
		TotalResults: &searchResultsText,
	}, nil
}

func (s *Scraper) ScrapeFromGoogleSearch(keywords []string) (chan model.KeywordScrapeResult, *sync.WaitGroup) {
	var wg sync.WaitGroup
	resultsChan := make(chan model.KeywordScrapeResult, len(keywords))

	for _, keyword := range keywords {
		wg.Add(1)
		go func(keyword string) {
			defer wg.Done()
			s.semaphore <- struct{}{}        // Acquire a slot in the concurrency limit.
			defer func() { <-s.semaphore }() // Release the slot when done.

			searchUrl := generateSearchURL(keyword)
			rawHTML, err := connector.CloudRunConnectorInstance.GetRenderedHTMLFromCloudRun(searchUrl)
			var information *model.ExtractedMetadata
			var scrapeResult model.KeywordScrapeResult
			currentTime := time.Now()
			if err != nil {
				scrapeResult = model.KeywordScrapeResult{
					Keyword:    keyword,
					Error:      err,
					SearchDate: currentTime,
				}
			} else {
				information, err = ExtractInformation(*rawHTML)
				scrapeResult = model.KeywordScrapeResult{
					Keyword:           keyword,
					RawHTML:           *rawHTML,
					SearchDate:        currentTime,
					Error:             err,
					ExtractedMetadata: *information,
				}
			}

			resultsChan <- scrapeResult
		}(keyword)
	}

	return resultsChan, &wg
}
