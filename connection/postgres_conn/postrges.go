package postgres_conn

import (
	"archie/utils/configer"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB = DataBase{}

type DataBase struct {
	db *gorm.DB
}

func (database *DataBase) Instance() *gorm.DB {
	return database.db
}

func (database *DataBase) InitDB() {
	dbConfig := configer.LoadDBConfig()
	db, err := gorm.
		Open(
			postgres.Open(fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=disable",
				dbConfig.Host,
				dbConfig.DBName,
				dbConfig.Port,
				dbConfig.User,
				dbConfig.Password,
			)),
			&gorm.Config{},
		)

	if err != nil {
		log.Fatal("Connect postgres failure")
	}

	database.db = db
}

func (database *DataBase) Transaction(cb func(db *gorm.DB) error) error {
	db := database.db
	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	if err := cb(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
