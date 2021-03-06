package user_resolvers

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"context"
)

type RegisterParams struct {
	Username    string `json:"username" form:"username" validate:"gt=4,lt=20,required"`
	Password    string `form:"password" validate:"required,gt=4,lt=20"`
	DisplayName string `json:"displayName" form:"displayName" validate:"required,gt=2,lt=10"`
	Email       string `json:"email" form:"email" validate:"email,required"`
}

type CreateUserParams struct {
	UserInfo RegisterParams `json:"userInfo"`
}

func (r *Resolver) CreateUser(ctx context.Context, params CreateUserParams) (*models.User, error) {
	userInfo := params.UserInfo

	if err := helper.ValidParams(userInfo); err != nil {
		return nil, err
	}

	if _, err := models.FindOneByUsername(userInfo.Username); err == nil {
		return nil, robust.USER_ALREADY_EXIST
	}

	user := models.User{}
	utils.CpStruct(&userInfo, &user)

	if err := user.Register(); err != nil {
		return nil, err
	}

	return &user, nil
}
