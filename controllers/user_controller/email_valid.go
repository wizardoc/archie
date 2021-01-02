package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmailValid(ctx *gin.Context) {
	var res helper.Res

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	user := models.User{
		ID:           claims.ID,
		IsValidEmail: true,
	}
	if err := user.UpdateUserInfo(); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	res.Success(user).Send(ctx)
}
