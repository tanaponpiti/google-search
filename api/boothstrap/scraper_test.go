package boothstrap

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"os"
	boothstrap_test "server-side/boothstrap/mock"
	"server-side/connector"
	"testing"
)

func TestNewScraper(t *testing.T) {
	concurrencyLimit := 5
	scraper := NewScraper(concurrencyLimit)
	assert.NotNil(t, scraper)
	assert.Equal(t, cap(scraper.semaphore), concurrencyLimit, "The semaphore should have a capacity equal to the concurrency limit")
}

func TestGenerateSearchURL(t *testing.T) {
	keyword := "test"
	expectedURL := "https://www.google.com/search?q=test&hl=EN"
	actualURL := generateSearchURL(keyword)
	assert.Equal(t, expectedURL, actualURL, "The generated URL does not match the expected URL")
}

func TestExtractInformation(t *testing.T) {
	t.Run("Success extraction", func(t *testing.T) {
		filePath := "./mock/search_cache.html"
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}
		sampleHTML := string(bytes)

		info, err := ExtractInformation(sampleHTML)
		assert.NoError(t, err, "Should not have encountered an error")
		assert.NotNil(t, info, "Extracted information should not be nil")

		assert.Equal(t, *info.AdWordsCount, 5)
		assert.Equal(t, *info.TotalLinks, 105)
		assert.Equal(t, *info.TotalResults, "About 40,500,000 results (0.27 seconds) ")
	})
	//TODO add more html sample for other test case scenario
}

func TestScraper_ScrapeFromGoogleSearch(t *testing.T) {
	mockConnector := new(boothstrap_test.MockHTMLRetrieverConnector)
	connector.HTMLRetrieverConnectorInstance = mockConnector
	keywords := []string{"test1", "test2"}

	sampleHTML := "<html><body>Sample HTML content</body></html>"

	mockConnector.On("GetRenderedHTML", mock.Anything).Return(&sampleHTML, nil).Twice()

	scraper := NewScraper(2)

	// Test
	resultsChan, wg := scraper.ScrapeFromGoogleSearch(keywords)
	wg.Wait()

	// Assertions
	close(resultsChan)
	count := 0
	for result := range resultsChan {
		assert.NotNil(t, result.ExtractedMetadata, "The extracted metadata should not be nil")
		assert.Nil(t, result.Error, "There should be no error in scraping")
		count++
	}

	assert.Equal(t, len(keywords), count, "The number of results should match the number of keywords")

	// Verify that the mock was called as expected
	mockConnector.AssertExpectations(t)
}
