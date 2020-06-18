package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAvatar struct {
	Avatar string `form:"avatar" validate:"required"`
}

// 更新用户头像
func UpdateAvatar(ctx *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(ctx)
	res := helper.Res{}

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	userAvatar := UserAvatar{}
	err = helper.BindWithValid(ctx, &userAvatar)

	if err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	user := models.User{
		Avatar: userAvatar.Avatar,
		ID:     parsedClaims.User.ID,
	}

	if err := user.UpdateAvatar(); err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, robust.DB_UPDATE_FAILURE)
		return
	}

	res.Send(ctx, nil)
}
