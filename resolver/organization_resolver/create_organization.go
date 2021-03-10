package organization_resolver

import (
	"archie/models"
	"archie/services/organization_service"
	"context"
)

type CreateOrganizationParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	IsPublic    bool   `json:"isPublic"`
}

type CreateOrganizationResolverParams struct {
	OrganizationInfo CreateOrganizationParams `json:"organizationInfo"`
}

func (r *OrganizationResolver) CreateOrganization(ctx context.Context, params CreateOrganizationResolverParams) (*models.Organization, error) {
	claims, err := r.Auth(ctx)
	if err != nil {
		return nil, err
	}

	orgInfo := params.OrganizationInfo
	org := &models.Organization{
		Name:        orgInfo.Name,
		Description: orgInfo.Description,
		Cover:       orgInfo.Cover,
		IsPublic:    orgInfo.IsPublic,
		Owner:       claims.ID,
	}

	if err := organization_service.CreateOrganization(org); err != nil {
		return nil, err
	}

	return org, err
}
