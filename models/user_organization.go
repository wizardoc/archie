package models

import (
	"archie/connection"
	"archie/utils"
	"time"
)

type UserOrganization struct {
	UserID         string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrganizationID string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	IsOwner        bool   `gorm:"type:bool"`
	JoinTime       int64
}

func (userOrganization *UserOrganization) New(isOwner bool) {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	userOrganization.JoinTime = time.Now().Unix()

	db.Create(userOrganization)
}
