package models

import "time"

type UserOrganization struct {
	UserID         string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrganizationID string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	IsOwner        bool   `gorm:"type:bool"`
	JoinTime       *time.Time
}
