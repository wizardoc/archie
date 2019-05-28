package models

import "time"

type Organization struct {
	ID           string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrganizeName string  `gorm:"type:varchar(20)"`
	Description  string  `gorm:"type:varchar(50)"`
	HasValid     bool    `gorm:"type:bool"`
	Users        *[]User `gorm:"many2many:user_organizations;"`
	CreateTime   *time.Time
}
