package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-side/model"
	"server-side/response"
	"server-side/service"
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
