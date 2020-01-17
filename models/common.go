package models

import (
	"archie/connection"
	"github.com/jinzhu/gorm"
)

func updateSig(model interface{}, name string, value interface{}) error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Model(model).Update(name, value).Error
	})
}
