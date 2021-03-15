package organization_resolver

import (
	"archie/models"
	"archie/services/organization_service"
	"archie/utils/jwt_utils"
	"context"
	"fmt"
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

func (r *OrganizationResolver) CreateOrganization(ctx context.Context, params CreateOrganizationResolverParams) (string, error) {
	var claims jwt_utils.LoginClaims
	if err := r.Auth(ctx, &claims); err != nil {
		return "", err
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
		return "", err
	}

	fmt.Println(org.ID)

	return org.ID, nil
}
