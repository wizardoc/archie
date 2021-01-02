package user_controller

import (
	"archie/models"
	"archie/robust"
	"archie/utils"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPasswordParams struct {
	NewPassword string `json:"newPassword" validate:"required,gt=4,lt=20" form:"newPassword"`
}

func ResetPassword(ctx *gin.Context) {
	var params ResetPasswordParams
	var res helper.Res

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	email, isExist := ctx.Get("email")
	if !isExist {
		res.Status(http.StatusBadRequest).Error(robust.EMAIL_IS_REQUIRED).Send(ctx)
		return
	}

	// find the user by id
	user := models.User{Email: email.(string)}
	if err := user.Find("email", user.Email); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	// the user to search by email does not exist
	if user.ID == "" {
		res.Status(http.StatusNotFound).Error(robust.USER_DOSE_NOT_EXIST).Send(ctx)
		return
	}

	hashNewPassword := utils.Hash(params.NewPassword)

	// new password can't equal to current password
	if user.Password == hashNewPassword {
		res.Status(http.StatusForbidden).Error(robust.REPEAT_PASSWORD).Send(ctx)
		return
	}

	updatedUser := models.User{
		ID:       user.ID,
		Password: hashNewPassword,
	}
	if err := updatedUser.UpdateUserInfo(); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	res.Send(ctx)
}
