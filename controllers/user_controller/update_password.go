package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdatePasswordParams struct {
	OriginPassword string `json:"originPassword" form:"originPassword" validate:"required"`
	NewPassword    string `json:"newPassword" form:"newPassword" validate:"required"`
}

func UpdatePassword(ctx *gin.Context) {
	var res helper.Res
	var params UpdatePasswordParams

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Send(ctx, err)
		return
	}

	user := models.User{
		ID: claims.ID,
	}
	if err := user.GetUserInfoByID(); err != nil {
		res.Status(http.StatusForbidden).Error(ctx, robust.USER_DOSE_NOT_EXIST)
		return
	}

	// the new password is equal to origin password
	if utils.Hash(params.NewPassword) == user.Password {
		res.Status(http.StatusForbidden).Error(ctx, robust.REPEAT_PASSWORD)
		return
	}

	// password is invalid
	if utils.Hash(params.OriginPassword) != user.Password {
		res.Status(http.StatusForbidden).Error(ctx, robust.ORIGIN_PASSWORD_FAILURE)
		return
	}

	user.Password = utils.Hash(params.NewPassword)
	if err := user.UpdateUserInfo(); err != nil {
		res.Status(http.StatusForbidden).Error(ctx, err)
		return
	}

	res.Send(ctx, user)
}
