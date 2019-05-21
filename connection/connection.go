package connection

import (
	"archie/utils/configer"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() (*gorm.DB, error) {
	dbConfig := configer.LoadDBConfig()

	return gorm.
		Open(
			"postgres",
			fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=disable",
				dbConfig.Host,
				dbConfig.DBName,
				dbConfig.Port,
				dbConfig.User,
				dbConfig.Password,
			),
		)
}
