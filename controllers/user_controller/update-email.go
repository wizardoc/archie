package user_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateEmail(ctx *gin.Context) {
	var res helper.Res
	email, exist := ctx.Get("email")

	if !exist {
		res.Status(http.StatusBadRequest).Error(robust.MISSING_PARAMS).Send(ctx)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	user := models.User{
		ID: claims.ID,
	}
	if err := user.GetUserInfoByID(); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	// refuse the request if the previous email has not been successful verify
	if !user.IsValidEmail {
		res.Status(http.StatusForbidden).Error(robust.NO_VALID_EMAIL).Send(ctx)
		return
	}

	// refuse the request if the previous email is equal to current email
	if email == user.Email {
		res.Status(http.StatusForbidden).Error(robust.REPEAT_EMAIL).Send(ctx)
		return
	}

	user.Email = email.(string)
	user.IsValidEmail = true
	if err := user.UpdateUserInfo(); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	res.Success(user).Send(ctx)
}
