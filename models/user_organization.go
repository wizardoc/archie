package models

import (
	"archie/connection/postgres_conn"
	"archie/utils"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type UserOrganization struct {
	UserID         string `gorm:"type:uuid;primary_key;"`
	OrganizationID string `gorm:"type:uuid;primary_key;"`
	User           *User
	Organization   *Organization
	IsOwner        bool  `gorm:"type:bool"`
	JoinTime       int32 `gorm:"type:bigint"`
}

type OrganizationOwnerInfo struct {
	Organization
	OwnerInfo   User   `json:"ownerInfo" gorm:"-"`
	Members     []User `json:"members" gorm:"-"`
	JoinTime    int32  `json:"joinTime"`
	Permissions []int  `json:"permissions"`
}

func (userOrganization *UserOrganization) TableName() string {
	return "user_organizations"
}

func (userOrganization *UserOrganization) IsExist() (bool, error) {
	userOrganizations := []UserOrganization{}

	if err := postgres_conn.DB.Instance().Where("user_id = ? AND organization_id = ?", userOrganization.UserID, userOrganization.OrganizationID).Find(&userOrganizations).Error; err != nil {
		return false, err
	}

	return len(userOrganizations) != 0, nil
}

// 寻找指定 Organization ID 的 members
func (userOrganization *UserOrganization) FindMembers(organizationID string, members *[]User) error {
	var userOrganizations []UserOrganization

	err := postgres_conn.DB.Instance().Model(userOrganization).Preload("User").Where("organization_id = ?", organizationID).Find(&userOrganizations).Error

	utils.ArrayMap(userOrganizations, func(item interface{}) interface{} {
		return *item.(UserOrganization).User
	}, members)

	return err
}

func (userOrganization *UserOrganization) New(isOwner bool) error {
	userOrganization.JoinTime = utils.Now()
	userOrganization.IsOwner = isOwner

	return postgres_conn.DB.Instance().Create(userOrganization).Error
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

	err := postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		var result *gorm.DB

		result = db.
			Raw(
				"SELECT * FROM user_organizations INNER JOIN organizations ON organization_id=organizations.id WHERE user_id = ?",
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
			Raw(fmt.Sprintf("SELECT * FROM users WHERE id IN (%s)", strings.Join(organizationIds, ","))).
			Scan(&owners)

		// dist organizations
		for i, organization := range infos {
			owner, ok := findOwnerByID(organization.Owner, owners)

			if !ok {
				continue
			}

			// attach members
			var members []User
			if err := userOrganization.FindMembers(infos[i].ID, &members); err == nil {
				infos[i].Members = members
			} else {
				infos[i].Members = []User{}
			}

			// attach permission info
			var permissionValues []int
			op := OrganizationPermission{UserID: userOrganization.UserID, OrganizationID: organization.ID}
			if err := op.AllAsValue(&permissionValues); err != nil {
				return err
			}

			infos[i].Permissions = permissionValues
			infos[i].OwnerInfo = owner
		}

		return result.Error
	})

	return infos, err
}
