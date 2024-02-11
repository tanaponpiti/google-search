package boothstrap

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	logLevelString := viper.GetString("LOG_LEVEL")
	logLevel, err := log.ParseLevel(strings.ToLower(logLevelString))
	if err != nil {
		logLevel = log.InfoLevel
		log.Warnf("Invalid log level '%s', defaulting to 'info'", logLevelString)
	}
	log.SetLevel(logLevel)
}
