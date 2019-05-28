package models

type Document struct {
	ID string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
}
