package user_resolvers

import (
	"archie/models"
	"context"
)

type UserInfoParams struct {
	ID *string `json:"id"`
}

func (r *Resolver) UserInfo(ctx context.Context, params UserInfoParams) (*models.User, error) {
	var parsedID string

	if params.ID == nil {
		claims, err := r.Auth(ctx)
		if err != nil {
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
