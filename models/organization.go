package models

import (
	"archie/connection"
	"archie/utils"
	"time"
)

type Organization struct {
	ID           string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrganizeName string  `gorm:"type:varchar(20)"`
	Description  string  `gorm:"type:varchar(50)"`
	HasValid     bool    `gorm:"type:bool;default:TRUE"`
	Users        *[]User `gorm:"many2many:user_organizations;"`
	CreateTime   time.Time
}

func (organization *Organization) NewOrganization() (success bool) {
	db, err := connection.GetDB()

	utils.Check(err, func() {
		success = false
	})

	defer db.Close()

	organization.HasValid = true
	organization.CreateTime = time.Now()

	db.Create(organization)

	return
}
