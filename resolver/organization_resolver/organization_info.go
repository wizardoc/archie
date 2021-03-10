package organization_resolver

import (
	"archie/models"
	"archie/robust"
	"context"
)

type OrganizationInfoParams struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

func (r *OrganizationResolver) OrganizationInfo(ctx context.Context, params OrganizationInfoParams) (org *models.Organization, err error) {
	if params.ID == nil && params.Name == nil {
		return nil, robust.ArchieError{Msg: "Argument ID or Name is required"}
	}

	if params.ID != nil {
		org = &models.Organization{ID: *params.ID}
		err = org.FindOneByID()
		return
	}

	if params.Name != nil {
		org = &models.Organization{Name: *params.Name}
		err = org.FindOneByOrganizeName()
		return
	}

	return
}
