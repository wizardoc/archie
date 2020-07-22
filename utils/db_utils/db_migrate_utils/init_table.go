package db_migrate_utils

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"gorm.io/gorm"
	"log"
)

// 初始化数据表，在此完成一些初始化数据的工作，migrateCB 提供一些 migrate 的操作
func InitTable(migrateCB func(db *gorm.DB) error) {
	isPermissionTableExist := postgres_conn.DB.Instance().Migrator().HasTable(models.Permission{})

	if err := migrateCB(postgres_conn.DB.Instance()); err != nil {
		log.Println(err)
	}

	models.InitPermissionData(isPermissionTableExist)
}
