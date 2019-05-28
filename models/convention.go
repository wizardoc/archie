package models

import "time"

type Convention struct {
	ID             string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ConventionName string `gorm:"type:varchar(30)"`
	LastModifyUser string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Permission     int    `gorm:"type:int"`
	LastModify     *time.Time
	CreateTime     *time.Time
}
