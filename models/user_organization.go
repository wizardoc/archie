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
	JoinTime       int64  `gorm:"type:bigint"`
}

func (userOrganization *UserOrganization) New(isOwner bool) {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	userOrganization.JoinTime = time.Now().Unix()
	userOrganization.IsOwner = isOwner

	db.Create(userOrganization)
}

type OrganizationOwnerInfo struct {
	User
	Organization
}

func (userOrganization *UserOrganization) FindUserJoinOrganizations() ([]OrganizationOwnerInfo, error) {
	db, err := connection.GetDB()
	var info []OrganizationOwnerInfo

	if err != nil {
		return info, err
	}

	defer db.Close()

	db.Raw(
		`select * from user_organizations inner join users on users.id=user_organizations.user_id inner join organizations on organizations.id=organization_id where user_id=?`,
		userOrganization.UserID,
	).Scan(&info)

	return info, nil
}
