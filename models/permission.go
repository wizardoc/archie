package models

import (
	"archie/connection/postgres_conn"
	permission_keys "archie/constants/permission"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type Permission struct {
	ID          string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	Value       int    `gorm:"type:int"json:"-"`
	Description string `gorm:"type:varchar(200)"json:"-"`
}

func init() {
	err := postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		//permission := Permission{}
		initRecords := []Permission{
			{Value: permission_keys.DOCUMENT_READ, Description: "document readable"},
			{Value: permission_keys.DOCUMENT_VIEW, Description: "document viewable"},
			{Value: permission_keys.DOCUMENT_WRITE, Description: "document writable"},
			{Value: permission_keys.CATEGORY_READ, Description: "category readable"},
			{Value: permission_keys.CATEGORY_WRITE, Description: "category writable"},
			{Value: permission_keys.CATEGORY_VIEW, Description: "category viewable"},
		}

		return db.Model(Permission{}).Updates(initRecords).Error
	})

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
