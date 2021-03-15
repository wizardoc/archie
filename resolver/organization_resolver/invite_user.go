package organization_resolver

import (
	"archie/services/organization_service"
	"archie/utils/jwt_utils"
	"context"
)

type InviteUserParams struct {
	Token string
}

func (r *OrganizationResolver) InviteUser(ctx context.Context, params InviteUserParams) (string, error) {
	var claims jwt_utils.LoginClaims
	// The invitedUser must be login to accept the invitation
	if err := r.Auth(ctx, &claims); err != nil {
		return "", err
	}

	return organization_service.InviteUser(params.Token, claims)
}
