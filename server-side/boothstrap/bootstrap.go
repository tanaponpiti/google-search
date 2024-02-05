package boothstrap

import (
	"github.com/spf13/viper"
	"server-side/database"
)

func Init() (err error) {
	err = LoadConfig()
	if err != nil {
		return err
	}
	InitLogger()
	err = database.InitDatabase(viper.GetString("DB_URI"))
	if err != nil {
		return err
	}
	return err
}
