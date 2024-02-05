package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"server-side/boothstrap"
	"server-side/route"
)

func main() {
	err := boothstrap.Init()
	if err != nil {
		log.Error("Error:", err)
		return
	}
	router := gin.Default()
	apiGroup := router.Group("/api")
	route.RegisterAuthRoutes(apiGroup)
	err = router.Run(fmt.Sprintf(":%s", viper.GetString("PORT")))
	if err != nil {
		log.Println("Error:", err)
		return
	}
}
