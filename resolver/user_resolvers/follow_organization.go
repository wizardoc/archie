package user_resolvers

import (
	"archie/services/user_service"
	"context"
)

type FollowOrganizationParams struct {
	ID string
}

func (r *UserResolver) FollowOrganization(ctx context.Context, params UnfollowOrganizationParams) (string, error) {
	claims, err := r.Auth(ctx)
	if err != nil {
		return "", err
	}

	if err := user_service.UnfollowOrganization(claims.ID, params.ID); err != nil {
		return "", err
	}

	return params.ID, nil
}
