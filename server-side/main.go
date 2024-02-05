package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"server-side/boothstrap"
)

func main() {
	err := boothstrap.Init()
	if err != nil {
		log.Error("Error:", err)
		return
	}
	router := gin.Default()
	err = router.Run(fmt.Sprintf(":%s", viper.GetString("PORT")))
	if err != nil {
		log.Println("Error:", err)
		return
	}
}
