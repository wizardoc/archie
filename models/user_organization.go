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

func (userOrganization *UserOrganization) New(isOwner bool) error {
	db, err := connection.GetDB()

	if err != nil {
		return err
	}

	defer db.Close()

	userOrganization.JoinTime = utils.Now()
	userOrganization.IsOwner = isOwner

	db.Create(userOrganization)

	return db.Error
}

type OrganizationOwnerInfo struct {
	OwnerInfo User `json:"ownerInfo"`
	Organization
	JoinTime int64 `json:"joinTime"`
}

func findOwnerByID(id string, owners []User) (User, bool) {
	for _, owner := range owners {
		if owner.ID == id {
			return owner, true
		}
	}

	return User{}, false
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

		fmt.Println(len(organizationIds))

		db.
			Raw(fmt.Sprintf("select * from users where id in (%s)", strings.Join(organizationIds, ","))).
			Scan(&owners)

		for i, organization := range infos {
			owner, ok := findOwnerByID(organization.Owner, owners)

			if !ok {
				continue
			}

			infos[i].OwnerInfo = owner
		}
	}

	return infos, nil
}
