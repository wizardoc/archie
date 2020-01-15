package models

import (
	"archie/connection"
	"archie/utils"
	"github.com/jinzhu/gorm"
)

type Organization struct {
	ID           string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"-"`
	OrganizeName string  `gorm:"type:varchar(20);unique;"json:"organizeName"`
	Description  string  `gorm:"type:varchar(50)"json:"description"`
	HasValid     bool    `gorm:"type:bool;default:TRUE"json:"hasValid"`
	Owner        string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"json:"-"` // related userID
	Users        *[]User `gorm:"many2many:user_organizations;"json:"-"`
	CreateTime   int64   `gorm:"type:bigint"json:"createTime"`
}

type OrganizationName struct {
	OrganizeName string
}

func (organization *Organization) FindOneByOrganizeName() error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Find(organization, "organize_name=?", organization.OrganizeName).Error
	})
}

func (organization *Organization) New(username string) error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		organization.HasValid = true
		organization.CreateTime = utils.Now()
		user, err := FindOneByUsername(username)

		if err != nil {
			return err
		}

		organization.Owner = user.ID

		return db.Create(organization).Error
	})
}

func (organization *Organization) GetAllNames() (names []OrganizationName, err error) {
	names = []OrganizationName{}
	err = connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Select("organize_name").Find(organization).Scan(&names).Error
	})

	return
}

func (organization *Organization) AllByUserId(id string) (organizations []Organization, err error) {
	organizations = []Organization{}
	err = connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Find(&organizations).Where("id=?", id).Error
	})

	return
}

func (organization *Organization) RemoveOrganization() error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		return db.Delete(organization).Error
	})
}
