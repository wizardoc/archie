package organization_service

import (
	"archie/connection/postgres_conn"
	"archie/constants/organization_rbac"
	"archie/models"
	"gorm.io/gorm"
	"time"
)

func CreateOrganization(org *models.Organization) error {
	user := models.User{ID: org.Owner}
	if err := user.GetUserInfoByID(); err != nil {
		return err
	}

	return postgres_conn.DB.Instance().Transaction(func(tx *gorm.DB) error {
		org.CreateTime = time.Now().String()
		if err := org.Create(); err != nil {
			return err
		}

		m := models.Member{
			OrganizationID: org.ID,
			UserID:         user.ID,
			Role:           organization_rbac.OWNER,
		}

		return m.Create()
	})
}
