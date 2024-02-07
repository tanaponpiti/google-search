package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"server-side/connector"
	"server-side/model"
	"strings"
)

func ExtractInformation(htmlContent string) (*model.ExtractedMetadata, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
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

func generateSearchURL(keyword string) string {
	encodedKeyword := url.QueryEscape(keyword)
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s&hl=EN", encodedKeyword)
	return searchURL
}

func ScrapeFromGoogleSearch(keywords []string) []model.KeywordScrapeResult {
	var keywordsScrapeResultList []model.KeywordScrapeResult
	for _, keyword := range keywords {
		searchUrl := generateSearchURL(keyword)
		rawHTML, err := connector.CloudRunConnectorInstance.GetRenderedHTMLFromCloudRun(searchUrl)
		information, err := ExtractInformation(*rawHTML)
		keywordsScrapeResultList = append(keywordsScrapeResultList, model.KeywordScrapeResult{
			Keyword:           keyword,
			RawHTML:           *rawHTML,
			Error:             err,
			ExtractedMetadata: *information,
		})
	}
	return keywordsScrapeResultList
}
