package user_resolvers

import (
	"archie/services/user_service"
	"context"
)

type FollowParams struct {
	ID string `json:"followUserID"`
}

func (r *UserResolver) FollowUser(ctx context.Context, params FollowParams) (string, error) {
	claims, err := r.Auth(ctx)
	if err != nil {
		return "", err
	}

	if err := user_service.FollowUser(claims.ID, params.ID); err != nil {
		return "", err
	}

	return params.ID, nil
}
