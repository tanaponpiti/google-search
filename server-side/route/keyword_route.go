package route

import (
	"github.com/gin-gonic/gin"
	"server-side/controller"
	"server-side/service"
)

func RegisterKeywordRoutes(rg *gin.RouterGroup) {
	keywordGroup := rg.Group("/keyword", service.AuthMiddleware())
	keywordGroup.POST("/", controller.AddKeyword)
	keywordGroup.POST("/search", controller.GetKeywordPage)
}
