package organization_resolver

import (
	"archie/models"
	"context"
)

type OrganizationInfoParams struct {
	ID string `json:"id"`
}

func (r *OrganizationResolver) OrganizationInfo(ctx context.Context, params OrganizationInfoParams) (*models.Organization, error) {
	organization := models.Organization{ID: params.ID}
	err := organization.FindOneByID()
	if err != nil {
		return nil, err
	}

	return &organization, nil
}
