package user_resolvers

import (
	"archie/services/user_service"
	"context"
)

type UnfollowUserParams struct {
	ID string `json:"id"`
}

func (r *UserResolver) UnfollowUser(ctx context.Context, params UnfollowUserParams) (string, error) {
	claims, err := r.Auth(ctx)
	if err != nil {
		return "", err
	}

	err = user_service.UnfollowUser(claims.ID, params.ID)
	if err != nil {
		return "", err
	}

	return params.ID, nil
}
