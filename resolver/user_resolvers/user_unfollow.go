package user_resolvers

import (
	"archie/services/user_service"
	"archie/utils/jwt_utils"
	"context"
)

type UnfollowUserParams struct {
	ID string `json:"id"`
}

func (r *UserResolver) UnfollowUser(ctx context.Context, params UnfollowUserParams) (string, error) {
	var claims jwt_utils.LoginClaims
	if err := r.Auth(ctx, &claims); err != nil {
		return "", err
	}

	if err := user_service.UnfollowUser(claims.ID, params.ID); err != nil {
		return "", err
	}

	return params.ID, nil
}
