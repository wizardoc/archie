package db_migrate_utils

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"github.com/jinzhu/gorm"
	"log"
)

// 初始化数据表，在此完成一些初始化数据的工作，migrateCB 提供一些 migrate 的操作
func InitTable(migrateCB func(db *gorm.DB) error) {
	if err := postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		isPermissionTableExist := db.HasTable(models.Permission{})

		err := migrateCB(db)

		models.InitPermissionData(isPermissionTableExist)

		return err
	}); err != nil {
		log.Println(err)
	}
}
