package boothstrap

import (
	"github.com/spf13/viper"
	"server-side/database"
	"server-side/repository"
)

func InitDatabase() (err error) {
	err = database.InitRedis(viper.GetString("REDIS_URI"), viper.GetString("REDIS_PASSWORD"), viper.GetInt("REDIS_DB"), viper.GetInt("REDIS_CONNECTION_POOL"))
	if err != nil {
		return err
	}
	err = database.InitGORM(viper.GetString("DB_URI"))
	if err != nil {
		return err
	}
	repository.InitTokenRepository()
	err = repository.InitUserRepository()
	if err != nil {
		return err
	}
	err = repository.InitKeywordRepository()
	if err != nil {
		return err
	}
	repository.InitSearchResultRepository()
	repository.InitPageDataRepository()
	return nil
}
