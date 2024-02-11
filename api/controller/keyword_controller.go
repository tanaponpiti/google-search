package controller

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-side/model"
	"server-side/response"
	"server-side/service"
	"strconv"
)

func GetKeywordPage(c *gin.Context) {
	var req = model.NewPaginationRequest[model.KeywordFilter]()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	pageResponse, err := service.GetKeywordPage(req)
	if err != nil {
		return
	}
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}
	c.JSON(http.StatusOK, pageResponse)
}

func AddKeyword(c *gin.Context) {
	var req model.KeywordCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	keywords := req.Keywords
	search, err := service.ScrapeFromGoogleSearch(keywords)
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": search})
}

func AddKeywordFromCSV(c *gin.Context) {
	// The maximum file size is set to 2MB
	const maxFileSize = 2 << 20 // 2 MB
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	if fileHeader.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size should not exceed 2MB"})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading CSV file"})
		return
	}

	// Assuming each record contains a single keyword, and the CSV does not have a header
	var keywords []string
	for _, record := range records {
		keywords = append(keywords, record[0])
	}
	search, err := service.ScrapeFromGoogleSearch(keywords)
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": search})
}

func DownloadKeywordHTMLCache(c *gin.Context) {
	searchResultID := c.Param("searchResultID")
	var id uint
	var err error
	if parsedID, errParse := strconv.ParseUint(searchResultID, 10, 32); errParse == nil {
		id = uint(parsedID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search result ID"})
		return
	}
	pageData, err := service.GetSearchResultHTMLCache(id)
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=\"keyword-cache.html\"")
	c.String(http.StatusOK, *pageData)
}
