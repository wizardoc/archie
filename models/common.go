package models

import (
	"archie/connection/postgres_conn"
	"github.com/jinzhu/gorm"
)

func updateSig(model interface{}, name string, value interface{}) error {
	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		return db.Model(model).Update(name, value).Error
	})
}
