package boothstrap

import (
	"github.com/spf13/viper"
	"server-side/connector"
)

func InitConnector() error {
	isStandalone := viper.GetBool("HTML_RETRIEVER_STANDALONE")
	if isStandalone {
		err := connector.InitStandaloneConnector(viper.GetString("HTML_RETRIEVER_URL"))
		if err != nil {
			return err
		}
		return nil
	} else {
		err := connector.InitCloudRunConnector(viper.GetString("HTML_RETRIEVER_URL"), viper.GetString("CLOUD_RUN_KEY_PATH"))
		if err != nil {
			return err
		}
		return nil
	}
}
