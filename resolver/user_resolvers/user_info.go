package user_resolvers

import (
	"archie/models"
	"archie/utils/jwt_utils"
	"context"
)

type UserInfoParams struct {
	ID *string `json:"id"`
}

func (r *UserResolver) UserInfo(ctx context.Context, params UserInfoParams) (*models.User, error) {
	var parsedID string
	var claims jwt_utils.LoginClaims
	if params.ID == nil {
		if err := r.Auth(ctx, &claims); err != nil {
			return nil, err
		}

		parsedID = claims.ID
	} else {
		parsedID = *params.ID
	}

	user := models.User{ID: parsedID}
	if err := user.GetUserInfoByID(); err != nil {
		return nil, err

	}

	return &user, nil
}
