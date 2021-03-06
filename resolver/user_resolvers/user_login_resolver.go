package user_resolvers

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/jwt_utils"
	"context"
)

type UserLoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRes struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func (r *UserResolver) Login(ctx context.Context, params UserLoginParams) (*UserLoginRes, error) {
	// check user is exist
	user, err := models.FindOneByUsername(params.Username)
	if err != nil {
		return nil, robust.USER_DOSE_NOT_EXIST
	}

	// wrong password
	if utils.Hash(params.Password) != user.Password {
		return nil, robust.LOGIN_PASSWORD_NOT_VALID
	}

	if err := user.UpdateLoginTime(); err != nil {
		return nil, err
	}

	claims := jwt_utils.LoginClaims{
		User: user,
	}

	return &UserLoginRes{User: user, Token: claims.SignJWT(24)}, err
}
