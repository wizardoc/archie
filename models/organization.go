package models

import (
	"archie/connection"
	"fmt"
	"os"
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

type OrganizationName struct {
	OrganizeName string
}

func (organization *Organization) New() (ok bool) {
	db, err := connection.GetDB()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return false
	}

	defer db.Close()

	organization.HasValid = true
	organization.CreateTime = time.Now()

	db.Create(organization)

	return true
}

func (organization *Organization) GetAllNames() (names []OrganizationName, ok bool) {
	db, err := connection.GetDB()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return nil, ok
	}

	defer db.Close()

	names = []OrganizationName{}
	db.Select("organize_name").Find(organization).Scan(&names)

	return names, true
}
