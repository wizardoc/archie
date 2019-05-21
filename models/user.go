package models

import (
	"time"
)

type User struct {
	ID           string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	DisplayName  string `gorm:"type:varchar(12)"`
	Username     string `gorm:"type:varchar(20);unique"`
	Password     string `gorm:"type:char(32)"`
	Email        string `gorm:"type:varchar(64)"`
	Avatar       string `gorm:"type:varchar(200)"`
	RegisterTime *time.Time
	LoginTime    *time.Time
}
