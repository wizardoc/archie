package organization_resolver

import (
	"archie/services/organization_service"
	"context"
)

type InviteUserParams struct {
	Token string
}

func (r *OrganizationResolver) InviteUser(ctx context.Context, params InviteUserParams) (string, error) {
	return organization_service.InviteUser(params.Token)
}
