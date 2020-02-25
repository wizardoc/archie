package postgres_conn

import (
	"archie/utils/configer"
	"fmt"
	"github.com/jinzhu/gorm"
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

func Transaction(cb func(db *gorm.DB) error) error {
	return WithPostgreConn(func(db *gorm.DB) error {
		tx := db.Begin()

		if tx.Error != nil {
			return tx.Error
		}

		if err := cb(tx); err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	})
}
