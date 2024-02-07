package boothstrap

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() error {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("TOKEN_EXPIRE_HOUR", "24")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", "0")
	viper.SetDefault("REDIS_CONNECTION_POOL", "100")
	if os.Getenv("APP_MODE") != "production" {
		if err := godotenv.Load(); err != nil {
			return err
		}
	}
	viper.AutomaticEnv()
	requiredKeys := []string{"DB_URI", "JWT_SECRET", "REDIS_URI", "CLOUD_RUN_URL", "CLOUD_RUN_KEY_PATH"}
	for _, key := range requiredKeys {
		if !viper.IsSet(key) {
			log.Fatal("Required key %s not set in environment", key)
			os.Exit(-1)
		}
	}
	return nil
}
