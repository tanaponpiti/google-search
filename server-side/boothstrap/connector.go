package boothstrap

import (
	"github.com/spf13/viper"
	"server-side/connector"
)

func InitConnector() error {
	err := connector.InitCloudRunConnector(viper.GetString("CLOUD_RUN_URL"), viper.GetString("CLOUD_RUN_KEY_PATH"))
	if err != nil {
		return err
	}
	return nil
}
