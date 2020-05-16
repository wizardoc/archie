package models

import (
	"archie/connection/postgres_conn"
	"archie/utils/db_utils"
	"github.com/jinzhu/gorm"
	"log"
)

const (
	// organization permissions
	ORG_DELETE = iota
	ORG_EDIT

	// category permissions
	CATEGORY_CREATE
	CATEGORY_EDIT

	// document permissions
	DOCUMENT_WRITE
	DOCUMENT_READ
	DOCUMENT_VIEW
	DOCUMENT_DELETE
	DOCUMENT_CREATE
)

type PermissionRecord struct {
	Value       int    `gorm:"type:int"json:"-"`
	Description string `gorm:"type:varchar(200)"json:"-"`
}

type Permission struct {
	ID          string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	Value       int    `gorm:"type:int"json:"-"`
	Description string `gorm:"type:varchar(200)"json:"-"`
}

func (p *Permission) Find(result *Permission) error {
	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		return db.Where(p).First(result).Error
	})
}

func InitPermissionData(isTableExist bool) {
	// 第一次建表的时候插入数据
	if isTableExist {
		return
	}

	initRecords := []PermissionRecord{
		{Value: DOCUMENT_READ, Description: "document readable"},
		{Value: DOCUMENT_VIEW, Description: "document viewable"},
		{Value: DOCUMENT_WRITE, Description: "document writable"},
		{Value: DOCUMENT_CREATE, Description: "document creatable"},
		{Value: DOCUMENT_DELETE, Description: "category deletable"},
		{Value: CATEGORY_EDIT, Description: "category editable"},
		{Value: CATEGORY_CREATE, Description: "category creatable"},
		{Value: ORG_DELETE, Description: "organization deletable "},
		{Value: ORG_EDIT, Description: "organization editable"},
	}

	if err := db_utils.BatchInsert("permissions", []string{"value", "description"}, initRecords); err != nil {
		log.Fatal(err)
	}
}