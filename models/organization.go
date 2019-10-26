package models

import (
	"archie/connection"
	"archie/utils"
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

func (organization *Organization) FindOneByOrganizeName() {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	db.Find(organization, "organize_name=?", organization.OrganizeName)
}

func (organization *Organization) New(username string) (ok bool) {
	db, err := connection.GetDB()

	if err != nil {
		return false
	}

	defer db.Close()

	organization.HasValid = true
	organization.CreateTime = utils.Now()
	user := FindOneByUsername(username)
	organization.Owner = user.ID

	db.Create(organization)

	return true
}

func (organization *Organization) GetAllNames() (names []OrganizationName, ok bool) {
	db, err := connection.GetDB()

	if err != nil {
		return nil, ok
	}

	defer db.Close()

	names = []OrganizationName{}
	db.Select("organize_name").Find(organization).Scan(&names)

	return names, true
}

func (organization *Organization) AllByUserId(id string) ([]Organization, error) {
	db, err := connection.GetDB()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var organizations []Organization
	db.Find(&organizations).Where("id=?", id)

	return organizations, nil
}
