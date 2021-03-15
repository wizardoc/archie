package user_resolvers

import (
	"archie/services/user_service"
	"archie/utils/jwt_utils"
	"context"
)

type FollowParams struct {
	ID string `json:"followUserID"`
}

func (r *UserResolver) FollowUser(ctx context.Context, params FollowParams) (string, error) {
	var claims jwt_utils.LoginClaims
	if err := r.Auth(ctx, &claims); err != nil {
		return "", err
	}

	if err := user_service.FollowUser(claims.ID, params.ID); err != nil {
		return "", err
	}

	return params.ID, nil
}
