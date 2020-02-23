package postgres_conn

import (
	"archie/utils"
	"github.com/jinzhu/gorm"
)

func WithPostgreConn(cb func(db *gorm.DB) error) error {
	db, err := GetDB()

	utils.Check(err)
	defer db.Close()

	return cb(db)
}
