package models

import (
	"archie/connection/postgres_conn"
	"archie/utils"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID             string     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"id"`
	Name           string     `gorm:"type:varchar(50)"json:"name"`
	Description    string     `gorm:"type:varchar(200)" json:"description"`
	Cover          string     `gorm:"type:varchar(200)"json:"cover"`
	UserID         string     `gorm:"type:uuid;" json:"userID"` // 分类创建者
	CreateTime     int64      `gorm:"type:bigint"json:"createTime"`
	LastModifyTime int64      `gorm:"type:bigint"json:"lastModifyTime"`
	OrganizationID string     `gorm:"type:uuid;json" json:"organizationID"` // 隶属的组织
	Documents      []Document `gorm:"foreign_key:CategoryId" json:"-"`
}

type ResCategory struct {
	Category
	CreateUser string `json:"createUser"`
}

func (category *Category) New() error {
	category.CreateTime = utils.Now()
	category.LastModifyTime = utils.Now()

	return postgres_conn.WithPostgreConn(func(db *gorm.DB) error {
		return db.Create(category).Error
	})
}

func (category *Category) All(resCategories *[]ResCategory) error {
	var categories []Category

	return postgres_conn.Transaction(func(db *gorm.DB) error {
		if err := db.Where("organization_id = ?", category.OrganizationID).Find(&categories).Error; err != nil {
			return err
		}

		resCategoriesTmp := make([]ResCategory, len(categories))

		for i, category := range categories {
			user := User{
				ID: category.UserID,
			}

			if err := user.GetUserInfoByID(); err != nil {
				return err
			}

			resCategoriesTmp[i].Category = category
			resCategoriesTmp[i].CreateUser = user.Username
		}

		*resCategories = resCategoriesTmp

		return nil
	})
}
