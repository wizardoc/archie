package organization_service

import (
	"archie/connection/postgres_conn"
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

		return org.AppendAssociation("Members", &user)
	})
}
