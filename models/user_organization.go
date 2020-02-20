package models

import (
	"archie/connection"
	"archie/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type UserOrganization struct {
	UserID         string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrganizationID string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	IsOwner        bool   `gorm:"type:bool"`
	JoinTime       int64  `gorm:"type:bigint"`
}

func (userOrganization *UserOrganization) New(isOwner bool) error {
	return connection.WithPostgreConn(func(db *gorm.DB) error {
		userOrganization.JoinTime = utils.Now()
		userOrganization.IsOwner = isOwner

		return db.Create(userOrganization).Error
	})
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
	var infos []OrganizationOwnerInfo

	err := connection.WithPostgreConn(func(db *gorm.DB) error {
		var result *gorm.DB

		result = db.
			Raw(
				"select * from user_organizations inner join organizations on organization_id=organizations.id where user_id=?",
				userOrganization.UserID,
			).
			Scan(&infos)

		if len(infos) == 0 {
			return nil
		}

		var organizationIds []string
		for _, info := range infos {
			organizationIds = append(organizationIds, fmt.Sprintf("'%s'", info.Owner))
		}

		var owners []User

		result = db.
			Raw(fmt.Sprintf("select * from users where id in (%s)", strings.Join(organizationIds, ","))).
			Scan(&owners)

		for i, organization := range infos {
			owner, ok := findOwnerByID(organization.Owner, owners)

			if !ok {
				continue
			}

			infos[i].OwnerInfo = owner
		}

		return result.Error
	})

	return infos, err
}
