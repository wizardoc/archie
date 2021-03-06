package user_resolvers

import (
	"archie/models"
	"context"
)

type UserInfoParams struct {
	UserName string
}

func (r *Resolver) UserInfo(ctx context.Context, params UserInfoParams) (*models.User, error) {
	user, err := models.FindOneByUsername(params.UserName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
