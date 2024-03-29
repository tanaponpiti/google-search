package boothstrap

import (
	"fmt"
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
	viper.SetDefault("CONCURRENT_SCRAPE_LIMIT", "3")
	viper.SetDefault("HTML_RETRIEVER_STANDALONE", "true")
	if os.Getenv("GIN_MODE") != "release" {
		if err := godotenv.Load(); err != nil {
			return err
		}
	}
	viper.AutomaticEnv()
	requiredKeys := []string{"DB_URI", "JWT_SECRET", "REDIS_URI", "HTML_RETRIEVER_URL"}
	if !viper.GetBool("HTML_RETRIEVER_STANDALONE") {
		requiredKeys = append(requiredKeys, "CLOUD_RUN_KEY_PATH")
	}
	for _, key := range requiredKeys {
		if !viper.IsSet(key) {
			log.Fatal(fmt.Sprintf("Required key %s not set in environment", key))
			os.Exit(-1)
		}
	}
	return nil
}
