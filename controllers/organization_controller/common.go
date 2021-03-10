package organization_controller

import "archie/models"

func CreateNewOrganization(name string, description string, username string) error {
	organization := models.Organization{
		Name:        name,
		Description: description,
	}

	return organization.New(username)
}

func InsertUserToOrganization(organizeName string, username string, isOwner bool) error {
	organization := models.Organization{Name: organizeName}
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

	op := models.OrganizationPermission{
		UserID:         user.ID,
		OrganizationID: organization.ID,
	}
	if err := op.NewMulti([]int{
		models.DOCUMENT_VIEW,
		models.DOCUMENT_READ,
		models.ORG_INVITE,
	}); err != nil {
		return err
	}

	return userOrganization.New(isOwner)
}
