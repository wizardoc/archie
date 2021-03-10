package user_service

import (
	"archie/connection/postgres_conn"
	"archie/models"
	"gorm.io/gorm"
)

func UnfollowOrganization(userID string, organizationID string) error {
	user := models.User{ID: userID}
	organization := models.Organization{ID: organizationID}

	return postgres_conn.DB.Transaction(func(db *gorm.DB) error {
		err := user.DeleteAssociation("FollowOrganizations", &organization)
		if err != nil {
			return err
		}

		return organization.DeleteAssociation("FollowUsers", &user)
	})
}
