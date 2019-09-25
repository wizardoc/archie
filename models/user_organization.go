package models

import (
	"archie/connection"
	"archie/utils"
	"fmt"
	"strings"
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

	userOrganization.JoinTime = utils.Now()
	userOrganization.IsOwner = isOwner

	db.Create(userOrganization)
}

type OrganizationOwnerInfo struct {
	OwnerInfo User `json:"ownerInfo"`
	Organization
	JoinTime int64 `json:"joinTime"`
}

func (userOrganization *UserOrganization) FindUserJoinOrganizations() ([]OrganizationOwnerInfo, error) {
	db, err := connection.GetDB()
	var infos []OrganizationOwnerInfo

	if err != nil {
		return infos, err
	}

	defer db.Close()

	db.
		Raw(
			"select * from user_organizations inner join organizations on organization_id=organizations.id where user_id=?",
			userOrganization.UserID,
		).
		Scan(&infos)

	if len(infos) != 0 {
		var organizationIds []string
		for _, info := range infos {
			organizationIds = append(organizationIds, fmt.Sprintf("'%s'", info.Owner))
		}

		var owners []User

		db.
			Raw(fmt.Sprintf("select * from users where id in (%s)", strings.Join(organizationIds, ","))).
			Scan(&owners)

		for i, ownerInfo := range owners {
			infos[i].OwnerInfo = ownerInfo
		}
	}

	return infos, nil
}
