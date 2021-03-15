package organization_resolver

import (
	"archie/models"
	"archie/robust"
	"context"
	"fmt"
)

type OrganizationInfoParams struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type OrganizationOverrideMember struct {
	models.Organization
	Members []models.UserWithRole
}

type OrganizationInfo struct {
	OrganizationOverrideMember
}

func (o *OrganizationInfo) Members() ([]models.UserWithRole, error) {
	mem := models.Member{OrganizationID: o.ID}

	var roles []models.UserWithRole
	if err := mem.FindUserWithRoleByOrgID(&roles); err != nil {
		return nil, err
	}

	fmt.Println(roles[0].Role)

	return roles, nil
}

func (r *OrganizationResolver) OrganizationInfo(ctx context.Context, params OrganizationInfoParams) (*OrganizationInfo, error) {
	var org models.Organization
	var err error

	if params.ID == nil && params.Name == nil {
		return nil, robust.ArchieError{Msg: "Argument ID or Name is required"}
	}

	if params.ID != nil {
		org = models.Organization{ID: *params.ID}
		err = org.FindOneByID()
	}

	if params.Name != nil {
		org = models.Organization{Name: *params.Name}
		err = org.FindOneByOrganizeName()
	}

	if err != nil {
		return nil, err
	}

	return &OrganizationInfo{OrganizationOverrideMember{Organization: org}}, nil
}
