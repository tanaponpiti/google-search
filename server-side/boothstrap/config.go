package boothstrap

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() error {
	viper.SetDefault("PORT", "8080")
	if os.Getenv("APP_MODE") != "production" {
		if err := godotenv.Load(); err != nil {
			return err
		}
	}
	viper.AutomaticEnv()
	return nil
}
