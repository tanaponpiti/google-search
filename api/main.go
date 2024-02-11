package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"server-side/boothstrap"
	"server-side/route"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	err := boothstrap.Init()
	if err != nil {
		log.Error("Error:", err)
		return
	}
	router := gin.Default()
	if os.Getenv("GIN_MODE") != "release" {
		router.Use(CORSMiddleware())
	}
	apiGroup := router.Group("/api")
	route.RegisterAuthRoutes(apiGroup)
	route.RegisterKeywordRoutes(apiGroup)
	err = router.Run(fmt.Sprintf(":%s", viper.GetString("PORT")))
	if err != nil {
		log.Println("Error:", err)
		return
	}
}
