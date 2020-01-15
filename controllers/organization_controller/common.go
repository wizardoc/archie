package organization_controller

import "archie/models"

func CreateNewOrganization(name string, description string, username string) error {
	organization := models.Organization{
		OrganizeName: name,
		Description:  description,
	}

	return organization.New(username)
}

func InsertUserToOrganization(organizeName string, username string, isOwner bool) error {
	organization := models.Organization{OrganizeName: organizeName}
	err := organization.FindOneByOrganizeName()

	if err != nil {
		return err
	}

	user, err := models.FindOneByUsername(username)

	if err != nil {
		return err
	}

	userOrganization := models.UserOrganization{
		UserID:         user.ID,
		OrganizationID: organization.ID,
	}

	return userOrganization.New(isOwner)
}
