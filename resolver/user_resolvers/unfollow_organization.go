package user_resolvers

import (
	"archie/services/user_service"
	"archie/utils/jwt_utils"
	"context"
)

type UnfollowOrganizationParams struct {
	ID string
}

func (r *UserResolver) UnfollowOrganization(ctx context.Context, params UnfollowOrganizationParams) (string, error) {
	var claims jwt_utils.LoginClaims
	if err := r.Auth(ctx, &claims); err != nil {
		return "", err
	}

	if err := user_service.UnfollowOrganization(claims.ID, params.ID); err != nil {
		return "", err
	}

	return params.ID, nil
}
