package user_service

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"gorm.io/gorm"
)

func FollowOrganization(userID string, organizationID string) error {
	user := models.User{ID: userID}
	organization := models.Organization{ID: organizationID}

	return postgres_conn.DB.Instance().Transaction(func(tx *gorm.DB) error {
		err := user.AppendAssociation("followOrganizations", &organization)
		if err != nil {
			return err
		}

		return organization.AppendAssociation("followUsers", &user)
	})
}
